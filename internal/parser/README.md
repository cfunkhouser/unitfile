# ANTLR Generated SystemD Unit file lexer and parser

Do not try to use this package directly (obviously). It is mostly generated from
the ANTLR grammars defined in the root of the `funkhouse.rs/unitfile` source
tree, and is hidden in an internal package to avoid complicating the public API
of `funkhouse.rs/unitfile`.

The tests in this package validate the contract of the ANTLR parser will contiue
to work for the parent package. If you make changes to `UnitLexer.g4` or
`UnitParser.g4`, be sure to regenerate the entire `funkhouse.rs/unitfile`
package recursively, and then run all of the tests. **Nontrivial grammar changes
are almost guaranteed to break the tests!**

From the root of `funkhouse.rs/unitfile`:

```console
$ go generate -x ./...
antlr4 -Dlanguage=Go -package parser -o internal/parser UnitLexer.g4 UnitParser.g4

$ go test ./...
ok  	funkhouse.rs/unitfile	0.015s
ok  	funkhouse.rs/unitfile/internal/parser	0.005s
```