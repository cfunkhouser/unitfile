package parser_test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/google/go-cmp/cmp"

	"funkhouse.rs/unitfile/internal/parser"
)

type testToken struct {
	ttype, value string
}

func TestLexer(t *testing.T) {
	opts := []cmp.Option{cmp.AllowUnexported(testToken{})}
	for tn, tc := range map[string]struct {
		data string
		want []testToken
	}{
		"documentation example 1a": {
			data: readTestFile(t, "testdata/example1a.unit"),
			want: []testToken{
				{"LBRACKET", "["},
				{"IDENTIFIER", "Section A"},
				{"RBRACKET", "]"},
				{"EOL", "\n"},
				{"IDENTIFIER", "KeyOne"},
				{"ASSIGN", "="},
				{"VALUE", "value"},
				{"VALUE", " 1"},
				{"EOL", "\n"},
				{"IDENTIFIER", "KeyTwo"},
				{"ASSIGN", "="},
				{"VALUE", "value"},
				{"VALUE", " 2"},
				{"EOF", "<EOF>"},
			},
		},
		"documentation example 1b": {
			data: readTestFile(t, "testdata/example1b.unit"),
			want: []testToken{
				{"LBRACKET", "["},
				{"IDENTIFIER", "Section B"},
				{"RBRACKET", "]"},
				{"EOL", "\n"},
				{"IDENTIFIER", "Setting"},
				{"ASSIGN", "="},
				{"DQUOTE", `"`},
				{"VALUE", "something"},
				{"DQUOTE", `"`},
				{"DQUOTE", `"`},
				{"VALUE", "some thing"},
				{"DQUOTE", `"`},
				{"DQUOTE", `"`},
				{"VALUE", "…"},
				{"DQUOTE", `"`},
				{"EOL", "\n"},
				{"IDENTIFIER", "KeyTwo"},
				{"ASSIGN", "="},
				{"VALUE", "value"},
				{"VALUE", " 2 "},
				{"VALUE", "\\\n"},
				{"VALUE", "	value 2 continued"},
				{"EOF", "<EOF>"},
			},
		},
		"documentation example 1c": {
			data: readTestFile(t, "testdata/example1c.unit"),
			want: []testToken{
				{"LBRACKET", "["},
				{"IDENTIFIER", "Section C"},
				{"RBRACKET", "]"},
				{"EOL", "\n"},
				{"IDENTIFIER", "KeyThree"},
				{"ASSIGN", "="},
				{"VALUE", "value"},
				{"VALUE", " 3"},
				{"VALUE", "\\\n"},
				{"COMMENT", "# this line is ignored\n"},
				{"COMMENT", "; this line is ignored too\n"},
				{"VALUE", "	\t\t\tvalue 3 continued"},
				{"EOF", "<EOF>"},
			},
		},
		"multi-section": {
			data: readTestFile(t, "testdata/multi-section.unit"),
			want: []testToken{
				{"COMMENT", "# This is an example unit file.\n"},
				{"COMMENT", ";\t\n"},
				{"LBRACKET", "["},
				{"IDENTIFIER", "Unit"},
				{"RBRACKET", "]"},
				{"EOL", "\n"},
				{"IDENTIFIER", "One"},
				{"ASSIGN", "="},
				{"VALUE", "1"},
				{"EOL", "\n"},
				{"IDENTIFIER", "Two"},
				{"ASSIGN", "="},
				{"EOL", "\n"},
				{"IDENTIFIER", "Three"},
				{"ASSIGN", "="},
				{"CONTINUATION", "\\\n"},
				{"VALUE", "3"},
				{"EOL", "\n"},
				{"EOL", "\n"},
				{"LBRACKET", "["},
				{"IDENTIFIER", "SectionA"},
				{"RBRACKET", "]"},
				{"EOL", "\n"},
				{"COMMENT", "# Comment\n"},
				{"COMMENT", "; Comment\n"},
				{"IDENTIFIER", "One"},
				{"ASSIGN", "="},
				{"DQUOTE", `"`},
				{"VALUE", "foo"},
				{"DQUOTE", `"`},
				{"DQUOTE", `"`},
				{"VALUE", "bar"},
				{"DQUOTE", `"`},
				{"EOL", "\n"},
				{"IDENTIFIER", "Two"},
				{"ASSIGN", "="},
				{"DQUOTE", `"`},
				{"VALUE", "foo"},
				{"DQUOTE", `"`},
				{"CONTINUATION", "\\\n"},
				{"DQUOTE", `"`},
				{"VALUE", "bar"},
				{"DQUOTE", `"`},
				{"CONTINUATION", "\\\n"},
				{"DQUOTE", `"`},
				{"VALUE", "baz"},
				{"DQUOTE", `"`},
				{"EOL", "\n"},
				{"IDENTIFIER", "Three"},
				{"ASSIGN", "="},
				{"VALUE", "This"},
				{"VALUE", " is a line"},
				{"VALUE", "\\\n"},
				{"COMMENT", "# with some\n"},
				{"VALUE", "\t"},
				{"COMMENT", ";comments inside\n"},
				{"VALUE", "but it continues."},
				{"EOL", "\n"},
				{"EOF", "<EOF>"},
			},
		},
	} {
		t.Run(tn, func(t *testing.T) {
			l := parser.NewUnitLexer(antlr.NewInputStream(tc.data))
			l.AddErrorListener(&testErrListener{tb: t})
			n := l.SymbolicNames
			var got []testToken
			for {
				t := l.NextToken()
				if t.GetTokenType() == antlr.TokenEOF {
					got = append(got, testToken{"EOF", t.GetText()})
					break
				}
				got = append(got, testToken{n[t.GetTokenType()], t.GetText()})
			}
			if diff := cmp.Diff(got, tc.want, opts...); diff != "" {
				t.Errorf("mismatch: (-got,+want):\n%v", diff)
			}
		})
	}
}
