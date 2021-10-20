package lexer

import (
	"github.com/p3t33/interpreter/token"
)

const START_OF_INPUT int = 0

type Lexer struct {
	input                string
	input_length         int
	offset               int  // current position in input(points to current char)
	single_raw_charecter byte // current char under examination
	tokens               map[string]token.Token
}

func CreateNewLexer(input string) *Lexer {
	temp := &Lexer{input: input, input_length: len(input), offset: START_OF_INPUT}

	temp.tokens = map[string]token.Token{
		"=": {Type: token.ASSIGN, Literal: "="},
		";": {Type: token.SEMICOLON, Literal: ";"},
		"(": {Type: token.LPAREN, Literal: "("},
		")": {Type: token.RPAREN, Literal: ")"},
		",": {Type: token.COMMA, Literal: ","},
		"+": {Type: token.PLUS, Literal: "+"},
		"{": {Type: token.LBRACE, Literal: "{"},
		"}": {Type: token.RBRACE, Literal: "}"},
		"":  {Type: token.EOF, Literal: ""}}

	return temp
}

func (l *Lexer) NextToken() token.Token {

	l.readSingleByteFromInput()

	var tok token.Token

	if 0 == l.single_raw_charecter {
		tok.Literal = ""
		tok.Type = token.EOF
	} else {
		tok = l.tokens[string(l.single_raw_charecter)]
	}

	return tok
}

func (l *Lexer) readSingleByteFromInput() {

	if l.offset < l.input_length {
		l.single_raw_charecter = l.input[l.offset]
		l.offset += 1

	} else {
		l.single_raw_charecter = byte(0)
	}
}
