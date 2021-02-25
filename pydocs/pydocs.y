%{
package pydocs

import (
    "fmt"
)

type KindToken uint

func (t KindToken) String() string {
    humanName, exists := map[KindToken]string{
        FreeText: "free",
        InlineComment: "#",
        MultilineComment: ">",
        StringLiteral: "str",
    }[t]
    if !exists  {
        return "unknownKindToken"
    }
    return humanName
}

const (
    unknownKindToken KindToken = iota
    FreeText
    StringLiteral
    InlineComment
    MultilineComment
)


type Token struct {
    val string
    line uint
    kind KindToken
}

func (t Token) String() string {
    if t.val == "+" {
        return t.val
    }
    if t.val == "CR" {
        return t.val
    }
    return fmt.Sprintf("%s=(%q)", t.kind, t.val)
}

%}

%union {
   empty struct{}
   token Token
   tokens []Token
   str string

   line uint
}

%token <empty> CR SINGLEQUOTES WHITESPACE
%token <token> TOKEN

%type <token> word
%type <tokens> words
%type <str> line lines

%start doc
%%

doc: /*empty*/
    | lines {
// TODO: append to global store
    };

lines:
    line {
        fmt.Printf("line=%v\n", $1)
    }|
    lines line {
        fmt.Printf("line=%v\n", $2)
    }
;

line:
    CR {
        $$ = fmt.Sprint("CR")
    }|
    words CR {
        $$ = fmt.Sprintf("%vCR", $1)
    }
;


words:
  word {
      $$ = []Token{$1}
  } |
  words word {
      $$ = append($1, $2)
  }
;


word: TOKEN |
    WHITESPACE {
        $$ = Token{val:"+"}
    };
