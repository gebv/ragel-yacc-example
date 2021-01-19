%{
package strparam

import (
    "fmt"
)
%}

%union {
   empty struct{}
   str string
}

%token <empty> OPENB CLOSEB
%token <str> WORD
%%

commands: /* empty */
        | commands command;

command:
        param
        |
        word;
word:
    WORD
    {
        fmt.Printf("word => %q\n", $1);
    };
param:
    OPENB WORD CLOSEB
    {
        fmt.Printf("param = %q\n", $2);
    };
