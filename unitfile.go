// Package unitfile contains utilities for encoding and decoding (certain) Go
// structs. The semantics are as similar to encoding/json as possible.
package unitfile

import "funkhouse.rs/unitfile/internal/parser"

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
