//line lexer.rl:1
package sql

import (
	"fmt"
)

//line lexer.go:11
const sql_parser_start int = 1
const sql_parser_first_final int = 1
const sql_parser_error int = 0

const sql_parser_en_main int = 1

//line lexer.rl:15

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

//line lexer.go:56
	{
		lex.cs = sql_parser_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line lexer.rl:51
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line lexer.go:72
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
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		}
		goto st_out
	tr0:
//line lexer.rl:68
		lex.te = (lex.p) + 1

		goto st1
	tr2:
//line lexer.rl:62
		lex.te = (lex.p) + 1
		{
			tok = COMMA
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr7:
//line lexer.rl:63
		lex.te = (lex.p)
		(lex.p)--
		{
			out.str = string(lex.data[lex.ts:lex.te])
			tok = STRING
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr9:
//line NONE:1
		switch lex.act {
		case 1:
			{
				(lex.p) = (lex.te) - 1
				tok = SELECT
				{
					(lex.p)++
					lex.cs = 1
					goto _out
				}
			}
		case 2:
			{
				(lex.p) = (lex.te) - 1
				tok = FROM
				{
					(lex.p)++
					lex.cs = 1
					goto _out
				}
			}
		case 3:
			{
				(lex.p) = (lex.te) - 1
				tok = AS
				{
					(lex.p)++
					lex.cs = 1
					goto _out
				}
			}
		case 5:
			{
				(lex.p) = (lex.te) - 1

				out.str = string(lex.data[lex.ts:lex.te])
				tok = STRING
				{
					(lex.p)++
					lex.cs = 1
					goto _out
				}
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

//line lexer.go:157
		switch lex.data[(lex.p)] {
		case 32:
			goto tr0
		case 44:
			goto tr2
		case 65:
			goto st2
		case 70:
			goto st4
		case 83:
			goto st7
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 13 {
				goto tr0
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		default:
			goto tr4
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
		switch lex.data[(lex.p)] {
		case 83:
			goto tr8
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	tr4:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:63
		lex.act = 5
		goto st3
	tr8:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:61
		lex.act = 3
		goto st3
	tr12:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:60
		lex.act = 2
		goto st3
	tr17:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:59
		lex.act = 1
		goto st3
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
//line lexer.go:242
		if lex.data[(lex.p)] == 95 {
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr9
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 82:
			goto st5
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch lex.data[(lex.p)] {
		case 79:
			goto st6
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr12
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 69:
			goto st8
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 76:
			goto st9
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 69:
			goto st10
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 67:
			goto st11
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr17
		case 95:
			goto tr4
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr4
			}
		case lex.data[(lex.p)] >= 65:
			goto tr4
		}
		goto tr7
	st_out:
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 2:
				goto tr7
			case 3:
				goto tr9
			case 4:
				goto tr7
			case 5:
				goto tr7
			case 6:
				goto tr7
			case 7:
				goto tr7
			case 8:
				goto tr7
			case 9:
				goto tr7
			case 10:
				goto tr7
			case 11:
				goto tr7
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:72

	return tok
}

func (lex *lexer) Error(e string) {
	fmt.Println("error:", e)
}
