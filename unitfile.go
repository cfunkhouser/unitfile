// Package unitfile contains utilities for encoding and decoding (certain) Go
// structs. The semantics are as similar to encoding/json as possible.
package unitfile

import (
	"reflect"

	"funkhouse.rs/unitfile/internal/parser"
)

// The decoder makes use of ANTLR-generated parser code, which is hidden in
// ./internal/parser to keep the published API and its documentation clean. The
// ANTLR grammars themselves are found in this directory. After making changes
// to them, run go generate for this package and its children, and then
// test everything!
//
// This line verifies that the decoder continues to satisfy the
// UnitParserListener interface across code regenerations, which may change
// the interface if the ANTLR grammar is modified.
//
//go:generate antlr4 -Dlanguage=Go -package parser -o internal/parser UnitLexer.g4 UnitParser.g4
var _ parser.UnitParserListener = &decoder{}

// unitfileStructTag is the struct tag used for encoding and decoding Units.
const unitfileStructTag = "unit"

// container abstracts a Unit section container - currently, assuming a struct -
// for use in transcoding.
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
