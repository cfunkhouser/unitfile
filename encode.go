package unitfile

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// TODO(christian): Default and required field values for strict mode.

// MarshalOption modifies the behavior of Marshal.
type MarshalOption func(*encoder)

var (
	// ErrInvalidSource is returned when the destination is invalid.
	ErrInvalidSource = errors.New("invalid marshal source")

	escapeContRe = regexp.MustCompile("\r?\n")
)

const topSection = ""

func descendable(t reflect.Type) bool {
	if t == nil {
		return false
	}
	kind := t.Kind()
	if kind == reflect.Pointer {
		kind = t.Elem().Kind()
	}
	// TODO(christian): Support walking maps?
	return kind == reflect.Struct
}

// Marshal returns the SystemD Unit file encoding of from, which must be a
// pointer to a struct.
func Marshal(from any, opts ...MarshalOption) ([]byte, error) {
	top := reflect.ValueOf(from)
	rt := reflect.TypeOf(from)
	if top.Kind() != reflect.Pointer || top.IsNil() || !descendable(rt) {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSource, rt)
	}

	e := encoder{
		unitName: "Unit", // TODO(christian): Allow setting this.
		src:      top,
		values: map[string]map[string][]string{
			topSection: make(map[string][]string),
		},
	}
	if err := visitAll(&e, rt, top); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := e.write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type encoder struct {
	unitName string
	src      reflect.Value
	level    int
	curr     string
	values   map[string]map[string][]string // section name -> field name -> values
}

// sortedKeys returns a sorted slice of map keys excluding the empty string.
func sortedKeys[T any](m map[string]T) (keys []string) {
	// Spare some cycles for determinism.
	for s := range m {
		if s == "" {
			continue
		}
		keys = append(keys, s)
	}
	sort.Strings(keys)
	return
}

func writeFields(out io.Writer, fields map[string][]string) error {
	for _, fn := range sortedKeys(fields) {
		lval := fields[fn]
		if len(lval) == 0 {
			continue
		}
		sort.Strings(lval)
		// TODO(christian): Format values appropriately.
		var err error
		for _, vs := range lval {
			_, err = fmt.Fprintf(out, "%v=%+v\n", fn, vs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *encoder) write(out io.Writer) error {
	// make sure the top section is handled first
	sns := append([]string{topSection}, sortedKeys(e.values)...)
	for _, sn := range sns {
		lfields := e.values[sn]
		if sn == topSection {
			sn = e.unitName
		}
		if _, err := fmt.Fprintf(out, "[%v]\n", sn); err != nil {
			return err
		}
		if err := writeFields(out, lfields); err != nil {
			return err
		}
		// TODO(christian): put a blank line between sections.
	}
	return nil
}

func (e *encoder) Enter(section reflect.StructField, val reflect.Value) error {
	if section.Anonymous {
		return nil
	}

	e.level++
	if e.level > 1 {
		return fmt.Errorf("%w: too deeply nested", ErrInvalidSource)
	}

	sn := section.Name
	if tn, ok := section.Tag.Lookup(unitfileStructTag); ok && tn != "" {
		sn = tn
	}
	e.values[sn] = make(map[string][]string)
	e.curr = sn
	return nil
}

func (e *encoder) Leave(field reflect.StructField, val reflect.Value) {
	if field.Anonymous {
		return
	}
	e.level--
}

func prerender(fn string, val reflect.Value, acc map[string][]string) error {
	if !val.IsValid() {
		return nil
	}
	k := val.Kind()
	switch k {
	case reflect.Bool:
		// Booleans are rendered no matter what, for now.
		// TODO(christian): Allow this conditionally, perhaps via tag.
		acc[fn] = []string{fmt.Sprintf("%v", val.Bool())}
	case reflect.Float32, reflect.Float64:
		if val.IsZero() {
			return nil
		}
		acc[fn] = []string{strings.TrimRight(fmt.Sprintf("%f", val.Float()), "0")}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if val.IsZero() {
			return nil
		}
		acc[fn] = []string{fmt.Sprintf("%d", val.Int())}
	case reflect.String:
		if val.IsZero() {
			return nil
		}
		// insert line continuations for multiline strings
		vs := escapeContRe.ReplaceAllString(strings.TrimSpace(val.String()), "\\\n")
		acc[fn] = append(acc[fn], vs)
	case reflect.Slice:
		l := val.Len()
		if l == 0 {
			return nil
		}
		for i := 0; i < l; i++ {
			v := reflect.Indirect(val.Index(i))
			if err := prerender(fn, v, acc); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("%w: field type %v is not supported", ErrInvalidSource, k)
	}
	return nil
}

func (e *encoder) Visit(field reflect.StructField, val reflect.Value) error {
	if descendable(field.Type) {
		// Handle this field when we enter it, instead of here.
		return nil
	}
	if field.Type.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	fn := field.Name
	if tn, ok := field.Tag.Lookup(unitfileStructTag); ok && tn != "" {
		fn = tn
	}
	return prerender(fn, val, e.values[e.curr])
}

func maybeDereference(typ reflect.Type, val reflect.Value) (reflect.Type, reflect.Value, bool) {
	if k := typ.Kind(); k == reflect.Ptr {
		return typ.Elem(), val.Elem(), true
	}
	return typ, val, false
}

type walker interface {
	Enter(field reflect.StructField, val reflect.Value) error
	Leave(field reflect.StructField, val reflect.Value)
	Visit(reflect.StructField, reflect.Value) error
}

func visitWalkable(w walker, field reflect.StructField, fval reflect.Value) error {
	if err := w.Enter(field, fval); err != nil {
		return err
	}
	defer w.Leave(field, fval)
	return visitAll(w, field.Type, fval)
}

func visitAll(w walker, typ reflect.Type, val reflect.Value) error {
	rt, rv, deref := maybeDereference(typ, val)
	if deref && val.IsNil() {
		return nil
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fval := rv.Field(i)

		// All members are visited first. If the member is walkable, it is then
		// entered.
		if err := w.Visit(field, fval); err != nil {
			return err
		}
		if !descendable(field.Type) {
			continue
		}
		if err := visitWalkable(w, field, fval); err != nil {
			return err
		}
	}
	return nil
}
