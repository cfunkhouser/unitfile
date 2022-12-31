// Package unitfile contains utilities for encoding and decoding (certain) Go
// structs. The semantics are as similar to encoding/json as possible.

//go:generate antlr4 -Dlanguage=Go -package parser -o internal/parser UnitLexer.g4 UnitParser.g4
package unitfile

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"funkhouse.rs/unitfile/internal/parser"
)

const unitfileStructTag = "unit"

var (
	// ErrInvalidDestination is returned when the destination is invalid.
	ErrInvalidDestination = errors.New("invalid unmarshal target")

	// ErrMalformedUnitFile describes an error decodig the Unit file input.
	ErrMalformedUnitFile = errors.New("malformed unit file")

	errUnknown = errors.New("unknown error")
	contValRe  = regexp.MustCompile("\\\\\r?\n")

	// This line verifies that the decoder continues to satisfy the
	// UnitParserListener interface across code regenerations, which may change
	// the interface if the ANTLR grammar is modified.
	_ parser.UnitParserListener = &decoder{}
)

// cleanValue sanitizes a field value by replacing line continuations with
// newlines, and trimming whitespace from the front and back.
func cleanValue(val string) string {
	return contValRe.ReplaceAllString(strings.TrimSpace(val), "\n")
}

// must wraps any function which returns a value and an error, and panics if the
// error is not nil.
func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// systemdBool parses boolean values according to the rules in the SystemD
// syntax definition. See
// https://www.freedesktop.org/software/systemd/man/systemd.syntax.html for
// details.
func systemdBool(val string) (bool, error) {
	val = strings.ToLower(val)
	switch val {
	case "1", "yes", "true", "on":
		return true, nil
	case "0", "no", "false", "off":
		return false, nil
	}
	return false, fmt.Errorf("%q cannot be converted to unitfile boolean value", val)
}

// Option for unmarshalling.
type Option func(*decoder)

// Strictly causes decoding to error when a field contained in the Unit contents
// would be ignored because the destination does not contain a matching field.
func Strictly(d *decoder) {
	d.strict = true
}

// Unmarshal parses SystemD Unit file-encoded data and stores the result in to,
// which must be a pointer to a struct. When an error is returned, the value of
// to is undefined.
func Unmarshal(data []byte, to any, opts ...Option) error {
	// The ANTLR walker panics when it encounteres malformed Unit file data. In
	// order to make this library safe to use in contexts where panicking would be
	// inconvenient, recover here and return the reason for the panic as an error.
	var err error
	func(data []byte, to any, opts []Option) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			if rerr, ok := r.(error); ok {
				err = rerr
				return
			}
			err = fmt.Errorf("%w: %v", errUnknown, r)
		}()
		err = panickyUnmarshal(data, to, opts...)
	}(data, to, opts)
	return err
}

// panickyUnmarshal is like Unmarshal, but will panic when parsing errors are
// encountered.
func panickyUnmarshal(data []byte, to any, opts ...Option) error {
	p := parser.NewUnitParser(
		antlr.NewCommonTokenStream(
			parser.NewUnitLexer(antlr.NewInputStream(string(data))),
			antlr.TokenDefaultChannel))

	dest := reflect.ValueOf(to)
	// TODO(christian): Make the invalid destination errors more clear.
	if dest.Kind() != reflect.Pointer || dest.IsNil() {
		return fmt.Errorf("%w: %v", ErrInvalidDestination, reflect.TypeOf(to))
	}
	dest = dest.Elem()
	if dest.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %v", ErrInvalidDestination, reflect.TypeOf(to))
	}
	d := &decoder{
		dest: dest,
	}
	for _, opt := range opts {
		opt(d)
	}

	antlr.ParseTreeWalkerDefault.Walk(d, p.Unit())
	return nil
}

type container struct {
	val   reflect.Value
	byTag map[string]string
}

// field resolves a container field by tag name and then struct field name, in
// that order. A zero-value Value is returned if neither is found.
func (c *container) field(tagOrName string) reflect.Value {
	if n := c.byTag[tagOrName]; n != "" {
		return c.val.FieldByName(n)
	}
	return c.val.FieldByName(tagOrName)
}

func prepareContainer(val reflect.Value) *container {
	tags := make(map[string]string)
	vt := val.Type()
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if tag, ok := field.Tag.Lookup(unitfileStructTag); ok {
			tags[tag] = field.Name
		}
	}
	return &container{
		val:   val,
		byTag: tags,
	}
}

// decoder handles parsing and decoding Unit files. It implements the ANTLR-
// generated UnitParserListener interface.
type decoder struct {
	*parser.BaseUnitParserListener

	// strict, when set, causes the walker to panic when it encounters a section
	// or field which cannot be mapped to the destination.
	strict bool

	// dest is the reflect value of the data structure to which we are decoding.
	dest reflect.Value

	// The following values are used to track the state of decoding.

	unitName    string
	currSection *container
	currField   reflect.Value
	currVals    []string
}

func ignoredField() reflect.Value {
	return reflect.New(reflect.TypeOf("string")).Elem()
}

func (d *decoder) ExitFieldName(c *parser.FieldNameContext) {
	id := c.IDENTIFIER().GetText()
	field := d.currSection.field(id)
	if !field.IsValid() {
		if d.strict {
			panic(fmt.Errorf("%w: destination struct contains no field %q", ErrInvalidDestination, id))
		}
		field = ignoredField()
	}
	if !field.CanSet() {
		panic(fmt.Errorf("%w: destination field %q cannot be set", ErrInvalidDestination, id))
	}
	if field.Kind() == reflect.Pointer {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		field = reflect.Indirect(field)
	}
	d.currField = field
}

func (d *decoder) handleFieldValue(val string) {
	val = cleanValue(val)
	if val == "" {
		return
	}

	switch k := d.currField.Kind(); k {
	case reflect.Bool:
		d.currField.SetBool(must(systemdBool(val)))
	case reflect.Float32, reflect.Float64:
		d.currField.SetFloat(must(strconv.ParseFloat(val, 64)))
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		d.currField.SetInt(must(strconv.ParseInt(val, 10, 64)))
	case reflect.String:
		d.currField.SetString(val)
	case reflect.Slice:
		d.currField.Set(reflect.Append(d.currField, reflect.ValueOf(val)))
	// TODO(christian): Date and Duration support
	default:
		panic(fmt.Errorf("%w: field type %v is not supported", ErrInvalidDestination, k))
	}
}

func (d *decoder) ExitValues(c *parser.ValuesContext) {
	for _, val := range d.currVals {
		d.handleFieldValue(val)
	}
	d.currVals = make([]string, 0)
}

func (d *decoder) ExitPlainValue(c *parser.PlainValueContext) {
	if len(d.currVals) == 0 {
		d.currVals = append(d.currVals, c.VALUE().GetText())
		return
	}
	d.currVals[0] += c.VALUE().GetText()
}

func (d *decoder) ExitQuotedValue(c *parser.QuotedValueContext) {
	d.currVals = append(d.currVals, c.VALUE().GetText())
}

func ignoredContainer() reflect.Value {
	return reflect.New(reflect.TypeOf(struct{}{})).Elem()
}

func (d *decoder) ExitSectionName(c *parser.SectionNameContext) {
	id := c.IDENTIFIER().GetText()
	// The first section name identifies the name of the Unit.
	if d.unitName == "" {
		d.unitName = id
		d.currSection = prepareContainer(d.dest)
		return
	}
	// Next section is a subsection of the Unit. Look for a member of dest
	// matching the name.
	// TODO(christian): Tag names.
	// TODO(christian): Allow ignoring fields which don't exist in struct.
	section := d.dest.FieldByName(id)
	if !section.IsValid() {
		if d.strict {
			panic(fmt.Errorf("%w: destination struct contains no field %q", ErrInvalidDestination, id))
		}
		section = ignoredContainer()
	}
	if section.Kind() == reflect.Pointer {
		section.Set(reflect.New(section.Type().Elem()))
		section = reflect.Indirect(section)
	}
	if k := section.Kind(); k != reflect.Struct {
		panic(fmt.Errorf("%w: type %v is not a supported container type", ErrInvalidDestination, k))
	}
	d.currSection = prepareContainer(section)
}
