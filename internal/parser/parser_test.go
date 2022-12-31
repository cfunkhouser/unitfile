// Package parser_test tests the Unit file parser.
package parser_test

import (
	"embed"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/google/go-cmp/cmp"

	"funkhouse.rs/unitfile/internal/parser"
)

// The tests and associated files named "documentation example" are taken
// directly from the SystemD documentation at
// https://www.freedesktop.org/software/systemd/man/systemd.syntax.html.
//
//go:embed testdata
var testdataFS embed.FS

func readTestFile(tb testing.TB, path string) string {
	tb.Helper()
	data, err := testdataFS.ReadFile(path)
	if err != nil {
		tb.Fatalf("failed to read file %q: %v", path, err)
	}
	return string(data)
}

type testErrListener struct {
	*antlr.DefaultErrorListener
	tb testing.TB
}

func (l *testErrListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.tb.Helper()
	l.tb.Fatalf("syntax error: %d:%d: %v", line, column, msg)
}

type testParseEvent struct {
	event, value string
}

type testWalker struct {
	*parser.BaseUnitParserListener
	got []testParseEvent
}

func (w *testWalker) ExitSectionName(c *parser.SectionNameContext) {
	w.got = append(w.got, testParseEvent{"BEGIN SECTION", c.IDENTIFIER().GetText()})
}

func (w *testWalker) ExitSection(c *parser.SectionContext) {
	w.got = append(w.got, testParseEvent{"END SECTION", ""})
}

func (w *testWalker) ExitFieldName(c *parser.FieldNameContext) {
	w.got = append(w.got, testParseEvent{"BEGIN FIELD", c.IDENTIFIER().GetText()})
}

func (w *testWalker) ExitPlainValue(c *parser.PlainValueContext) {
	w.got = append(w.got, testParseEvent{"PLAIN VALUE FRAGMENT", c.VALUE().GetText()})
}

func (w *testWalker) ExitQuotedValue(c *parser.QuotedValueContext) {
	w.got = append(w.got, testParseEvent{"QUOTED VALUE FRAGMENT", c.VALUE().GetText()})
}

func (w *testWalker) ExitValues(c *parser.ValuesContext) {
	w.got = append(w.got, testParseEvent{"END FIELD", ""})
}

func TestParser(t *testing.T) {
	opts := []cmp.Option{cmp.AllowUnexported(testParseEvent{})}
	for tn, tc := range map[string]struct {
		data string
		want []testParseEvent
	}{
		"documentation example 1a": {
			data: readTestFile(t, "testdata/example1a.unit"),
			want: []testParseEvent{
				{"BEGIN SECTION", "Section A"},
				{"BEGIN FIELD", "KeyOne"},
				{"PLAIN VALUE FRAGMENT", "value"},
				{"PLAIN VALUE FRAGMENT", " 1"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "KeyTwo"},
				{"PLAIN VALUE FRAGMENT", "value"},
				{"PLAIN VALUE FRAGMENT", " 2"},
				{"END FIELD", ""},
				{"END SECTION", ""},
			},
		},
		"documentation example 1b": {
			data: readTestFile(t, "testdata/example1b.unit"),
			want: []testParseEvent{
				{"BEGIN SECTION", "Section B"},
				{"BEGIN FIELD", "Setting"},
				{"QUOTED VALUE FRAGMENT", "something"},
				{"QUOTED VALUE FRAGMENT", "some thing"},
				{"QUOTED VALUE FRAGMENT", "…"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "KeyTwo"},
				{"PLAIN VALUE FRAGMENT", "value"},
				{"PLAIN VALUE FRAGMENT", " 2 "},
				{"PLAIN VALUE FRAGMENT", "\\\n"},
				{"PLAIN VALUE FRAGMENT", "\tvalue 2 continued"},
				{"END FIELD", ""},
				{"END SECTION", ""},
			},
		},
		"documentation example 1c": {
			data: readTestFile(t, "testdata/example1c.unit"),
			want: []testParseEvent{
				{"BEGIN SECTION", "Section C"},
				{"BEGIN FIELD", "KeyThree"},
				{"PLAIN VALUE FRAGMENT", "value"},
				{"PLAIN VALUE FRAGMENT", " 3"},
				{"PLAIN VALUE FRAGMENT", "\\\n"},
				{"PLAIN VALUE FRAGMENT", "\t\t\t\tvalue 3 continued"},
				{"END FIELD", ""},
				{"END SECTION", ""},
			},
		},
		"multi-section": {
			data: readTestFile(t, "testdata/multi-section.unit"),
			want: []testParseEvent{
				{"BEGIN SECTION", "Unit"},
				{"BEGIN FIELD", "One"},
				{"PLAIN VALUE FRAGMENT", "1"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "Two"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "Three"},
				{"PLAIN VALUE FRAGMENT", "3"},
				{"END FIELD", ""},
				{"END SECTION", ""},
				{"BEGIN SECTION", "SectionA"},
				{"BEGIN FIELD", "One"},
				{"QUOTED VALUE FRAGMENT", "foo"},
				{"QUOTED VALUE FRAGMENT", "bar"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "Two"},
				{"QUOTED VALUE FRAGMENT", "foo"},
				{"QUOTED VALUE FRAGMENT", "bar"},
				{"QUOTED VALUE FRAGMENT", "baz"},
				{"END FIELD", ""},
				{"BEGIN FIELD", "Three"},
				{"PLAIN VALUE FRAGMENT", "This"},
				{"PLAIN VALUE FRAGMENT", " is a line"},
				{"PLAIN VALUE FRAGMENT", "\\\n"},
				{"PLAIN VALUE FRAGMENT", "\t"},
				{"PLAIN VALUE FRAGMENT", "but it continues."},
				{"END FIELD", ""},
				{"END SECTION", ""},
			},
		},
	} {
		t.Run(tn, func(t *testing.T) {
			l := parser.NewUnitLexer(antlr.NewInputStream(tc.data))
			l.AddErrorListener(&testErrListener{tb: t})
			w := &testWalker{}
			p := parser.NewUnitParser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
			antlr.ParseTreeWalkerDefault.Walk(w, p.Unit())
			if diff := cmp.Diff(w.got, tc.want, opts...); diff != "" {
				t.Errorf("mismatch: (-got,+want):\n%v", diff)
			}
		})
	}
}
