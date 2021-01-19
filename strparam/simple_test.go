package strparam

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	lex := newLexer([]byte(`foo { bar} fz { bazzzz}`))
	e := yyParse(lex)
	fmt.Println(e)
}

func zzzBenchmarkDemo(b *testing.B) {
	lex := newLexer([]byte(`foo { bar} fz { bazzzz}`))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		yyParse(lex)
	}
}
