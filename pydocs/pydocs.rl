package pydocs

import (
        "fmt"
)

%%{
    machine pydocs;

    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    action capture {
        // fmt.Printf("capture: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)
    }
    action releaseCR {
        // fmt.Printf("release: pe=%d ts=%d te=%d p=%d\n", lex.pe, lex.ts, lex.te, lex.p)
        // fmt.Printf("release: %q\n", string(lex.data[lex.ts:lex.p]))
        lex.p-- // shift left one character
    }
    action lineCounter { out.line++}

    whitespace = [ \t]+;
    newline = [\n\r] @lineCounter;

    # Each line of a block comment starts with a # and a single space (unless it is indented text inside the comment).

    tripleQuotes = "'''" | '"""';
    singleQuotes = "'" | '"';

    escSingleQuotes = [\\]{1} singleQuotes;
    notEscSingleQuotes = [^\\]{1} . singleQuotes;

    # textword = [^ \t\n\r]+ -- tripleQuotes -- startLineComment -- singleQuotes;
    anyWithoutSpace = [^ \t\n\r]+;

    anyNotQuotes = (anyWithoutSpace -- [\\"'])+;
    strLiteral = singleQuotes (anyNotQuotes|whitespace|escSingleQuotes)* singleQuotes;

    inlineComment = '#' (anyWithoutSpace|whitespace)* newline >capture %releaseCR;

    multilineCommentContent = (anyWithoutSpace|whitespace|newline)+ -- tripleQuotes;
    multilineCommentBoundary = whitespace* tripleQuotes newline;
    multilineComment = multilineCommentBoundary
        multilineCommentContent
        multilineCommentBoundary >capture %releaseCR;
}%%

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
        pe: len(data),
    }
    %% write init;
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    %%{
        main := |*
            strLiteral => {
               out.token.val = string(lex.data[lex.ts:lex.te])
               out.token.line = out.line
               out.token.kind = StringLiteral

               tok = TOKEN
               fbreak;
            };
            multilineComment => {
                out.token.val = string(lex.data[lex.ts:lex.te])
                out.token.line = out.line
                out.token.kind = MultilineComment

                tok = TOKEN
                fbreak;
            };
            inlineComment => {
                out.token.val = string(lex.data[lex.ts:lex.te])
                out.token.line = out.line
                out.token.kind = InlineComment

                tok = TOKEN
                fbreak;
            };
            newline => {
                tok = CR
                fbreak;
            };
            whitespace => {
                tok = WHITESPACE
                fbreak;
            };
            anyNotQuotes => {
                out.token.val = string(lex.data[lex.ts:lex.te])
                out.token.line = out.line
                out.token.kind = FreeText

                tok = TOKEN
                fbreak;
            };
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
}
