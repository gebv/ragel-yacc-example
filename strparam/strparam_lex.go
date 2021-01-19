//line strparam.rl:1
package strparam

import (
	"fmt"
)

//line strparam_lex.go:11
const strparam_start int = 1
const strparam_first_final int = 1
const strparam_error int = 0

const strparam_en_main int = 1

//line strparam.rl:15

type lexer struct {
	// It must be an array containting the data to process.
	data []byte

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
	p int

	// This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
	ts int

	// Also a pointer to character data.
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data: data,
		pe:   len(data),
	}

//line strparam_lex.go:56
	{
		lex.cs = strparam_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line strparam.rl:51
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line strparam_lex.go:72
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		switch lex.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		}
		goto st_out
	tr0:
//line strparam.rl:75
		lex.te = (lex.p) + 1

		goto st1
	tr3:
//line strparam.rl:59
		lex.te = (lex.p) + 1
		{
			// fmt.Printf("ragel: openb = %q\n", string(lex.data[lex.ts:lex.te]))
			tok = OPENB
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr4:
//line strparam.rl:70
		lex.te = (lex.p) + 1
		{
			// fmt.Printf("ragel: closeb = %q\n", string(lex.data[lex.ts:lex.te]))
			tok = CLOSEB
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr5:
//line strparam.rl:64
		lex.te = (lex.p)
		(lex.p)--
		{
			out.str = string(lex.data[lex.ts:lex.te])
			// fmt.Printf("ragel: word = %q\n", string(lex.data[lex.ts:lex.te]))
			tok = WORD
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	st1:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
		lex.ts = (lex.p)

//line strparam_lex.go:131
		switch lex.data[(lex.p)] {
		case 32:
			goto tr0
		case 123:
			goto tr3
		case 125:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 13 {
				goto tr0
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st2
			}
		case lex.data[(lex.p)] >= 65:
			goto st2
		}
		goto tr5
	st_out:
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 2:
				goto tr5
			}
		}

	_out:
		{
		}
	}

//line strparam.rl:79

	return tok
}

func (lex *lexer) Error(e string) {
	fmt.Println("error:", e)
}
