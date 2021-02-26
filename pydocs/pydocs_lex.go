//line pydocs.rl:1
package pydocs

import (
	"fmt"
)

//line pydocs_lex.go:11
const pydocs_start int = 28
const pydocs_first_final int = 28
const pydocs_error int = 0

const pydocs_en_main int = 28

//line pydocs.rl:49

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

//line pydocs_lex.go:56
	{
		lex.cs = pydocs_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line pydocs.rl:85
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line pydocs_lex.go:72
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		switch lex.cs {
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 1:
			goto st_case_1
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
		case 31:
			goto st_case_31
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 0:
			goto st_case_0
		case 25:
			goto st_case_25
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 26:
			goto st_case_26
		case 34:
			goto st_case_34
		case 27:
			goto st_case_27
		case 35:
			goto st_case_35
		}
		goto st_out
	tr0:
//line pydocs.rl:121
		(lex.p) = (lex.te) - 1
		{
			tok = WHITESPACE
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr3:
//line NONE:1
		switch lex.act {
		case 1:
			{
				(lex.p) = (lex.te) - 1

				out.token.val = string(lex.data[lex.ts:lex.te])
				out.token.line = out.line
				out.token.kind = StringLiteral

				tok = TOKEN
				{
					(lex.p)++
					lex.cs = 28
					goto _out
				}
			}
		case 5:
			{
				(lex.p) = (lex.te) - 1

				tok = WHITESPACE
				{
					(lex.p)++
					lex.cs = 28
					goto _out
				}
			}
		}

		goto st28
	tr28:
//line pydocs.rl:93
		lex.te = (lex.p) + 1
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = StringLiteral

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr30:
//line pydocs.rl:125
		(lex.p) = (lex.te) - 1
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = FreeText

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr36:
//line pydocs.rl:23
		out.line++
//line pydocs.rl:117
		lex.te = (lex.p) + 1
		{
			tok = CR
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr40:
//line pydocs.rl:125
		lex.te = (lex.p)
		(lex.p)--
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = FreeText

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr41:
//line pydocs.rl:121
		lex.te = (lex.p)
		(lex.p)--
		{
			tok = WHITESPACE
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr44:
//line pydocs.rl:18

		// fmt.Printf("release: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)
		// fmt.Printf("release: %q\n", string(lex.data[lex.ts:lex.p]))
		lex.p-- // shift left one character

//line pydocs.rl:101
		lex.te = (lex.p)
		(lex.p)--
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = MultilineComment

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr45:
//line pydocs.rl:93
		lex.te = (lex.p)
		(lex.p)--
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = StringLiteral

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	tr46:
//line pydocs.rl:18

		// fmt.Printf("release: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)
		// fmt.Printf("release: %q\n", string(lex.data[lex.ts:lex.p]))
		lex.p-- // shift left one character

//line pydocs.rl:109
		lex.te = (lex.p)
		(lex.p)--
		{
			out.token.val = string(lex.data[lex.ts:lex.te])
			out.token.line = out.line
			out.token.kind = InlineComment

			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 28
				goto _out
			}
		}
		goto st28
	st28:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
//line NONE:1
		lex.ts = (lex.p)

//line pydocs_lex.go:300
		switch lex.data[(lex.p)] {
		case 9:
			goto tr35
		case 10:
			goto tr36
		case 13:
			goto tr36
		case 32:
			goto tr35
		case 34:
			goto st23
		case 35:
			goto tr38
		case 39:
			goto st27
		case 92:
			goto st0
		}
		goto st29
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr40
		case 32:
			goto tr40
		case 34:
			goto tr40
		case 39:
			goto tr40
		case 92:
			goto tr40
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr40
		}
		goto st29
	tr35:
//line NONE:1
		lex.te = (lex.p) + 1

//line pydocs.rl:121
		lex.act = 5
		goto st30
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
//line pydocs_lex.go:353
		switch lex.data[(lex.p)] {
		case 9:
			goto tr35
		case 32:
			goto tr35
		case 34:
			goto st1
		case 39:
			goto st21
		}
		goto tr41
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if lex.data[(lex.p)] == 34 {
			goto st2
		}
		goto tr0
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if lex.data[(lex.p)] == 34 {
			goto st3
		}
		goto tr0
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		}
		goto tr3
	tr4:
//line pydocs.rl:23
		out.line++
		goto st4
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line pydocs_lex.go:404
		switch lex.data[(lex.p)] {
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 34:
			goto st15
		case 39:
			goto st18
		}
		goto st5
	tr6:
//line pydocs.rl:23
		out.line++
		goto st5
	tr9:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line pydocs_lex.go:431
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr10
		case 39:
			goto tr11
		}
		goto st5
	tr10:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line pydocs_lex.go:458
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr12
		case 39:
			goto tr11
		}
		goto st5
	tr12:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line pydocs_lex.go:485
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr13
		case 39:
			goto tr11
		}
		goto st5
	tr13:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st8
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line pydocs_lex.go:512
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto st9
		}
		goto tr3
	tr14:
//line pydocs.rl:23
		out.line++
		goto st31
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
//line pydocs_lex.go:531
		goto tr44
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto st10
		}
		goto tr3
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		}
		goto tr3
	tr11:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st11
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line pydocs_lex.go:570
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr10
		case 39:
			goto tr17
		}
		goto st5
	tr17:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st12
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line pydocs_lex.go:597
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr10
		case 39:
			goto tr18
		}
		goto st5
	tr18:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st13
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line pydocs_lex.go:624
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 39:
			goto st14
		}
		goto tr3
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 39:
			goto st10
		}
		goto tr3
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr20
		case 39:
			goto tr11
		}
		goto st5
	tr20:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st16
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
//line pydocs_lex.go:679
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr21
		case 39:
			goto tr11
		}
		goto st5
	tr21:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st17
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
//line pydocs_lex.go:706
		if lex.data[(lex.p)] == 34 {
			goto st9
		}
		goto tr3
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr10
		case 39:
			goto tr22
		}
		goto st5
	tr22:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st19
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
//line pydocs_lex.go:742
		switch lex.data[(lex.p)] {
		case 9:
			goto tr9
		case 10:
			goto tr6
		case 13:
			goto tr6
		case 32:
			goto tr9
		case 34:
			goto tr10
		case 39:
			goto tr23
		}
		goto st5
	tr23:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

		goto st20
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
//line pydocs_lex.go:769
		if lex.data[(lex.p)] == 39 {
			goto st14
		}
		goto tr3
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		if lex.data[(lex.p)] == 39 {
			goto st22
		}
		goto tr0
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		if lex.data[(lex.p)] == 39 {
			goto st3
		}
		goto tr0
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto tr27
		case 39:
			goto tr28
		case 92:
			goto st25
		}
		goto st24
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto tr28
		case 39:
			goto tr28
		case 92:
			goto st25
		}
		goto st24
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 34:
			goto st24
		case 39:
			goto st24
		}
		goto st0
	tr27:
//line NONE:1
		lex.te = (lex.p) + 1

//line pydocs.rl:93
		lex.act = 1
		goto st32
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
//line pydocs_lex.go:856
		if lex.data[(lex.p)] == 34 {
			goto st3
		}
		goto tr45
	tr38:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st33
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
//line pydocs_lex.go:871
		switch lex.data[(lex.p)] {
		case 9:
			goto st26
		case 10:
			goto tr32
		case 13:
			goto tr32
		case 32:
			goto st26
		case 34:
			goto st26
		case 39:
			goto st26
		case 92:
			goto st26
		}
		goto tr38
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr32
		case 13:
			goto tr32
		}
		goto st26
	tr32:
//line pydocs.rl:15

		// fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)

//line pydocs.rl:23
		out.line++
		goto st34
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
//line pydocs_lex.go:914
		goto tr46
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto tr28
		case 39:
			goto tr33
		case 92:
			goto st25
		}
		goto st24
	tr33:
//line NONE:1
		lex.te = (lex.p) + 1

//line pydocs.rl:93
		lex.act = 1
		goto st35
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
//line pydocs_lex.go:946
		if lex.data[(lex.p)] == 39 {
			goto st3
		}
		goto tr45
	st_out:
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
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
	_test_eof31:
		lex.cs = 31
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
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 29:
				goto tr40
			case 30:
				goto tr41
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr3
			case 4:
				goto tr3
			case 5:
				goto tr3
			case 6:
				goto tr3
			case 7:
				goto tr3
			case 8:
				goto tr3
			case 31:
				goto tr44
			case 9:
				goto tr3
			case 10:
				goto tr3
			case 11:
				goto tr3
			case 12:
				goto tr3
			case 13:
				goto tr3
			case 14:
				goto tr3
			case 15:
				goto tr3
			case 16:
				goto tr3
			case 17:
				goto tr3
			case 18:
				goto tr3
			case 19:
				goto tr3
			case 20:
				goto tr3
			case 21:
				goto tr0
			case 22:
				goto tr0
			case 32:
				goto tr45
			case 33:
				goto tr40
			case 26:
				goto tr30
			case 34:
				goto tr46
			case 35:
				goto tr45
			}
		}

	_out:
		{
		}
	}

//line pydocs.rl:136

	return tok
}

func (lex *lexer) Error(e string) {
	fmt.Println("error:", e)
}
