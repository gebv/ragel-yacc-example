package pydocs

import (
	"fmt"
	"testing"
)

var case1 = `

foo # inline comments

# inline comments

"""
abc

foo="bar" # inline comments

bac
"""
bar baz

	"""
	abc

	bac

	"""

qwd
	# inline comments

foo="wdwd"
foo="wd\"wd\""
foo = "bar"r ""
""""

baz
`

func TestDemo(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	lex := newLexer([]byte(case1))
	e := yyParse(lex)
	fmt.Println(e)

	// === RUN   TestDemo
	// line=CR
	// line=CR
	// line=[free=("foo") + #=("# inline comments")]CR
	// line=CR
	// line=[#=("# inline comments")]CR
	// line=CR
	// line=[>=("\"\"\"\nabc\n\nfoo=\"bar\" # inline comments\n\nbac\n\"\"\"")]CR
	// line=[free=("bar") + free=("baz")]CR
	// line=CR
	// line=[>=("\t\"\"\"\n\tabc\n\n\tbac\n\n\t\"\"\"")]CR
	// line=CR
	// line=[free=("qwd")]CR
	// line=[+ #=("# inline comments")]CR
	// line=CR
	// line=[free=("foo=") str=("\"wdwd\"")]CR
	// line=[free=("foo=") str=("\"wd\\\"wd\\\"\"")]CR
	// line=[free=("foo") + free=("=") + str=("\"bar\"") free=("r") + str=("\"\"")]CR
	// line=[str=("\"\"") str=("\"\"")]CR
	// line=CR
	// line=[free=("baz")]CR
	// 0
	// --- PASS: TestDemo (0.00s)
}
