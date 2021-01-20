package sql

import (
	"fmt"
	"testing"
)

func TestParseSQL(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	lex := newLexer([]byte(`SELECT fa AS fff, fb FROM t`))
	e := yyParse(lex)
	fmt.Println(e)

	// === RUN   TestParseSQL
	// command fields: [field name = "fa" (as "fff") field name = "fb"], table: table name = "t"
	// 0
	// --- PASS: TestParseSQL (0.00s)
}
