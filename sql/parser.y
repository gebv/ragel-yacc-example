%{
package sql

import (
    "fmt"
)
%}

%union {
  str string
  arrstr []string
}

%token COMMA SELECT FROM AS
%token <str> STRING

%type <str> command, select_stmt, base_select, table_reference, field
%type <arrstr> list_fields

%start any_command
%%

any_command:
  command semicolon_opt
  {
    fmt.Println("command", $1)
  };

semicolon_opt:
/*empty*/ {}
| ';' {};

command:
  select_stmt
  {
    $$ = $1
  };

select_stmt:
  base_select
  {
    $$ = $1
  };

base_select:
  SELECT list_fields FROM table_reference
  {
    $$ = fmt.Sprintf("fields: %v, table: %v", $2, $4)
  };

list_fields: field
  {
    $$ = []string{$1}
  }
| list_fields COMMA field
  {
    $$ = append($$, $3)
  };

field:
  STRING
  {
    $$ = fmt.Sprintf("field name = %q", $1)
  }
| STRING AS STRING
  {
    $$ = fmt.Sprintf("field name = %q (as %q)", $1, $3)
  }

table_reference:
  STRING
  {
    $$ = fmt.Sprintf("table name = %q", $1)
  };
