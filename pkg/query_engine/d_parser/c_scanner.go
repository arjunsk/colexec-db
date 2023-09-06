package parser

import (
	"bytes"
	"strings"
)

const (
	UnknownType int = 0
)

type Scanner struct {
	pos   int
	query string
}

func newScanner(input string) *Scanner {
	//log.Println("new scanner", input)
	return &Scanner{
		pos:   0,
		query: input,
	}
}

func (scanner *Scanner) nextToken() (int, string) {
	var ch rune
scan:
	ch, eof := scanner.pop()
	if eof {
		return 0, ""
	}

	switch {
	case isSpace(ch):
		scanner.readSpace()
		goto scan

	case ch == '.' || isDigit(ch):
		scanner.unPop()
		return scanner.readNumber()

	case ch == '\'' || ch == '"':
		scanner.unPop()
		return scanner.readString()

	case ch == '+':
		return ADD, "+"

	case ch == '-':
		return DEC, "-"

	case ch == '*':
		return STAR, "*"

	case ch == '<':
		nextCh, _ := scanner.peek()
		if nextCh == '=' {
			scanner.pop()
			return LTE, "<="
		} else {
			if nextCh == '>' {
				scanner.pop()
				return NEQ, "<>"
			}
		}
		return LT, "<"

	case ch == '>':
		nextCh, _ := scanner.peek()
		if nextCh == '=' {
			scanner.pop()
			return GTE, ">="
		}
		return GT, ">"

	case ch == '=':
		return EQ, "="

	case ch == '(':
		return LEFTC, "("

	case ch == ')':
		return RIGHTC, ")"

	case ch == ';':
		return SEMICOLON, ";"

	case ch == ',':
		return COMMA, ","

	default:
		scanner.unPop()
		return scanner.readToken()
	}

}

func (scanner *Scanner) pop() (rune, bool) {
	if scanner.pos < len(scanner.query) {
		ch := rune(scanner.query[scanner.pos])
		scanner.pos++
		return ch, false
	}
	return rune(' '), true
}

func (scanner *Scanner) unPop() {
	scanner.pos--
}

func (scanner *Scanner) peek() (rune, bool) {
	if scanner.pos < len(scanner.query) {
		return rune(scanner.query[scanner.pos]), false
	}
	return rune(' '), true
}

func (scanner *Scanner) readSpace() {
	for scanner.pos < len(scanner.query) {
		ch, eof := scanner.peek()
		if !eof && isSpace(ch) {
			scanner.pop()
		} else {
			break
		}
	}
}

func (scanner *Scanner) readNumber() (int, string) {
	typ := INTEGER
	var buf bytes.Buffer

	ch, eof := scanner.peek()
	if !eof && ch == '-' {
		scanner.pop()
		buf.WriteRune(ch)
	}
	for scanner.pos < len(scanner.query) {
		ch, eof := scanner.peek()
		if eof || isSpace(ch) {
			break
		}
		if !eof && !isDigit(ch) && ch != '.' {
			break
		}
		if ch == '.' {
			typ = FLOAT
		}
		scanner.pop()
		buf.WriteRune(ch)
	}

	return typ, string(buf.Bytes())
}

func (scanner *Scanner) readString() (int, string) {
	ch0, eof := scanner.peek()
	if eof || (ch0 != '\'' && ch0 != '"') {
		return UnknownType, ""
	}

	var buf bytes.Buffer
	scanner.pop()
	for scanner.pos < len(scanner.query) {
		ch, eof := scanner.pop()
		if ch == ch0 {
			return STRING, string(buf.Bytes())
		} else {
			if eof {
				return UnknownType, string(buf.Bytes())
			} else {
				buf.WriteRune(ch)
			}
		}
	}

	return UnknownType, string(buf.Bytes())
}

func (scanner *Scanner) readToken() (int, string) {
	var buf bytes.Buffer
	for scanner.pos < len(scanner.query) {
		ch, eof := scanner.pop()
		if eof || isSpace(ch) {
			break
		}
		if !eof && !isLetter(ch) && !isDigit(ch) && ch != '_' {
			scanner.unPop()
			break
		}
		buf.WriteRune(ch)
	}

	str := string(buf.Bytes())
	upper := strings.ToUpper(str)
	switch upper {
	case "SELECT":
		return SELECT, upper
	case "FROM":
		return FROM, upper
	case "WHERE":
		return WHERE, upper
	case "AS":
		return AS, upper
	case "GROUP":
		return GROUP, upper
	case "BY":
		return BY, upper
	case "ORDER":
		return ORDER, upper
	case "LIMIT":
		return LIMIT, upper
	case "CREATE":
		return CREATE, upper
	case "UPDATE":
		return UPDATE, upper
	case "INSERT":
		return INSERT, upper
	case "INTO":
		return INTO, upper
	case "DELETE":
		return DELETE, upper
	case "AND":
		return AND, upper
	case "OR":
		return OR, upper
	}

	return NAME, str
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
