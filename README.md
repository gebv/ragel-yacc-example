# ragel-yacc-example

Examples
* [strparam](strparam) named params parser. [See more details strparam/simple_test.go](strparam/simple_test.go)
* [sql](sql) example of sql parser. [See more details sql/sql_test.go](sql/sql_test.go)

References

* https://github.com/calio/ragel-cheat-sheet Ragel Cheat Sheet
* https://en.wikipedia.org/wiki/LL_parser
* https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/ simple calculator
* https://github.com/postgres/postgres/blob/master/src/backend/parser/gram.y postgres yacc grammer parser
* https://blog.gopheracademy.com/advent-2014/parsers-lexers/ https://github.com/benbjohnson/sql-parser native parser on the go
* Simple sql parser example (academic) https://github.com/vignif/lex-yacc-SQL-parser
* Example on go ragel + yacc for writing your own simple pseudo language https://mhamrah.medium.com/lexing-with-ragel-and-parsing-with-yacc-using-go-81e50475f88f
* example of how ragel works to parse the expression `<param> = <value>` http://thingsaaronmade.com/blog/a-simple-intro-to-writing-a-lexer-with-ragel.html
* https://github.com/mhamrah/thermostat example ragel+yacc for handle commands
```
target temperature 5
heat on
target temperature 10
heat off
```
* https://github.com/johan-bjareholt/lua-compiler/blob/master/src/grammar.yy lua grammar
* https://github.com/vcflib/bio-vcf/blob/master/ragel/gen_vcfheaderline_parser.rl example ragel lex for DSL language
* cockroachdb yacc grammer parser https://github.com/cockroachdb/cockroach/blob/master/pkg/sql/parser/sql.y
