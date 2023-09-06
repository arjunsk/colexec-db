package parser

import (
	"log"
	"strconv"
)

type Lexer struct {
	scanner *Scanner
	stmts   []Statement
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		scanner: newScanner(input),
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	typ, str := l.scanner.nextToken()
	//log.Println(typ, str)

	switch typ {
	case 0:
		return 0
	case INTEGER:
		lval.int64, _ = strconv.ParseInt(str, 10, 64)
	case FLOAT:
		lval.float64, _ = strconv.ParseFloat(str, 64)
	//case STRING, NAME:
	//    lex.str = str
	case EQ, NEQ, LT, LTE, GT, GTE, AND, OR, ADD, DEC:
		lval.int = typ
	}
	lval.str = str

	return typ
}

func (l *Lexer) Error(err string) {
	log.Fatal(err)
}

func (l *Lexer) AppendStmt(stmt Statement) {
	l.stmts = append(l.stmts, stmt)
}
