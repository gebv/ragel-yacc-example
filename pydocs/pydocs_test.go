package pydocs

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	lex := newLexer([]byte(`
foo bar
foo
    bar
"foo"
"foo bar"
"foo \"\"bar" baz
baz "foo \"bar"
baz

'qwd' # inline comments

#foobar

"""
yes
no

wd
"wqd"

"""

pok="s"

iu

`))
	e := yyParse(lex)
	fmt.Println(e)

	// === RUN   TestDemo
	// line=CR
	// line=[free=("foo") + free=("bar")]CR
	// line=[free=("foo")]CR
	// line=[+ free=("bar")]CR
	// line=[str=("\"foo\"")]CR
	// line=[str=("\"foo bar\"")]CR
	// line=[str=("\"foo \\\"\\\"bar\"") + free=("baz")]CR
	// line=[free=("baz") + str=("\"foo \\\"bar\"")]CR
	// line=[free=("baz")]CR
	// line=CR
	// line=[str=("'qwd'") + #=("# inline comments\n")]CR
	// line=[#=("#foobar\n")]CR
	// line=[>=("\"\"\"\nyes\nno\n\nwd\n\"wqd\"\n\n\"\"\"\n")]CR
	// line=[free=("pok=") str=("\"s\"")]CR
	// line=CR
	// line=[free=("iu")]CR
	// line=CR
	// 0
	// --- PASS: TestDemo (0.00s)
}
