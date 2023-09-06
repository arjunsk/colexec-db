package parser

import (
	"fmt"
)

func Parse(sql string) ([]Statement, error) {
	lexer := NewLexer(sql)
	if yyParse(lexer) != 0 {
		return nil, fmt.Errorf("error while parsing")
	}
	return lexer.stmts, nil
}
