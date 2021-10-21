package lexer

import (
	"github.com/p3t33/interpreter/token"
)

const START_OF_INPUT int = 0

type Lexer struct {
	input                string
	input_length         int
	position             int // current position in input(points to current char)
	read_position        int
	single_raw_charecter byte // current char under examination
	tokens               map[string]token.Token
}

func CreateNewLexer(input string) *Lexer {
	temp := &Lexer{input: input, input_length: len(input)}

	temp.tokens = map[string]token.Token{
		"=":  {Type: token.ASSIGN, Literal: "="},
		";":  {Type: token.SEMICOLON, Literal: ";"},
		"(":  {Type: token.LPAREN, Literal: "("},
		")":  {Type: token.RPAREN, Literal: ")"},
		",":  {Type: token.COMMA, Literal: ","},
		"+":  {Type: token.PLUS, Literal: "+"},
		"{":  {Type: token.LBRACE, Literal: "{"},
		"}":  {Type: token.RBRACE, Literal: "}"},
		"-":  {Type: token.MINUS, Literal: "-"},
		"!":  {Type: token.BANG, Literal: "!"},
		"*":  {Type: token.ASTERISK, Literal: "*"},
		"/":  {Type: token.SLASH, Literal: "/"},
		"<":  {Type: token.LT, Literal: "<"},
		">":  {Type: token.GT, Literal: ">"},
		"==": {Type: token.EQ, Literal: "=="},
		"!=": {Type: token.NOT_EQ, Literal: "!="}}

	temp.readSingleByteFromInput()

	return temp
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if tok, ok := l.tokens[string(l.single_raw_charecter)]; ok {
		if "=" == tok.Literal {
			if "=" == string(l.peekNextByteFromInput()) {
				tok = l.tokens["=="]
				l.readSingleByteFromInput()
			}
		}

		if "!" == tok.Literal {
			if "=" == string(l.peekNextByteFromInput()) {
				tok = l.tokens["!="]
				l.readSingleByteFromInput()
			}
		}

		l.readSingleByteFromInput()
		return tok
	}

	if true == l.isEmptyValue() {
		tok.Literal = ""
		tok.Type = token.EOF
		l.readSingleByteFromInput()
		return tok
	}

	if true == l.isLetter() {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	} else if l.isDigit() {
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return tok

	} else {
		tok.Literal = string(l.single_raw_charecter)
		tok.Type = token.ILLEGAL
	}

	l.readSingleByteFromInput()
	return tok

}

func (l *Lexer) readSingleByteFromInput() {

	if l.read_position >= l.input_length {
		l.single_raw_charecter = 0
	} else {
		l.single_raw_charecter = l.input[l.read_position]
	}

	l.position = l.read_position
	l.read_position += 1
}

func (l *Lexer) peekNextByteFromInput() byte {
	if l.read_position >= l.input_length {
		return 0
	} else {
		return l.input[l.read_position]
	}
}

func (l *Lexer) isLetter() bool {
	if 'a' <= l.single_raw_charecter && 'z' >= l.single_raw_charecter || 'A' <= l.single_raw_charecter && 'Z' >= l.single_raw_charecter || '_' == l.single_raw_charecter {
		return true
	} else {
		return false
	}
}

func (l *Lexer) isEmptyValue() bool {
	if 0 == l.single_raw_charecter {
		return true
	} else {
		return false
	}
}

func (l *Lexer) readIdentifier() string {
	start_index := l.position
	for true == l.isLetter() {
		l.readSingleByteFromInput()
	}

	return (l.input[start_index:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.single_raw_charecter == ' ' || l.single_raw_charecter == '\t' || l.single_raw_charecter == '\n' || l.single_raw_charecter == '\r' {
		l.readSingleByteFromInput()
	}
}

func (l *Lexer) readNumber() string {
	start_position := l.position
	for l.isDigit() {
		l.readSingleByteFromInput()
	}

	return l.input[start_position:l.position]
}
func (l *Lexer) isDigit() bool {
	return '0' <= l.single_raw_charecter && l.single_raw_charecter <= '9'
}
