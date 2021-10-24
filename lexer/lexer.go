package lexer

import (
	"github.com/p3t33/interpreter/token"
)

type Lexer struct {
	input                string
	input_length         int
	position             int // current position in input(points to current char)
	read_position        int
	single_raw_charecter byte // current char under examination
	tokens               map[string]token.Token
	single_token         token.Token
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

	l.resetTokenValue()
	l.skipWhitespace()

	if l.isSpecialCharecterExistInTokenMap() {
		if l.isEqualsign() {
			if l.isNextByteAnEqualsign() {
				l.createEqualToOperatorToken()
				l.readSingleByteFromInput()
			}
		}
		if l.isExclamationMark() {
			if l.isNextByteAnEqualsign() {
				l.createNotEqualToOperatorToken()
				l.readSingleByteFromInput()
			}
		}

		l.readSingleByteFromInput()
		return l.single_token
	}

	if true == l.isEmptyValue() {
		l.createEndOfFileToken()
		l.readSingleByteFromInput()
		return l.single_token
	}

	if true == l.isLetter() {
		l.createIdentToken()
		return l.single_token
	} else if l.isDigit() {
		l.createNumberToken()
		return l.single_token

	} else {
		l.createillegalToken()
	}

	l.readSingleByteFromInput()
	return l.single_token

}

func (l *Lexer) createEqualToOperatorToken() {
	l.single_token = l.tokens["=="]
}

func (l *Lexer) createNotEqualToOperatorToken() {
	l.single_token = l.tokens["!="]
}

func (l *Lexer) isNextByteAnExclamationMark() bool {
	if "!" == string(l.peekNextByteFromInput()) {
		return true
	} else {
		return false
	}
}

func (l *Lexer) isNextByteAnEqualsign() bool {
	if "=" == string(l.peekNextByteFromInput()) {
		return true
	} else {
		return false
	}
}

func (l *Lexer) isEqualsign() bool {
	if "=" == l.single_token.Literal {
		return true

	} else {
		return false
	}
}

func (l *Lexer) isExclamationMark() bool {
	if "!" == l.single_token.Literal {
		return true

	} else {
		return false
	}
}

func (l *Lexer) createillegalToken() {
	l.single_token.Literal = string(l.single_raw_charecter)
	l.single_token.Type = token.ILLEGAL
}

func (l *Lexer) createIdentToken() {
	l.single_token.Literal = l.readIdentifier()
	l.single_token.Type = token.LookupIdent(l.single_token.Literal)
}

func (l *Lexer) createEndOfFileToken() {
	l.single_token.Literal = ""
	l.single_token.Type = token.EOF
}

func (l *Lexer) isSpecialCharecterExistInTokenMap() bool {
	var ok bool
	if l.single_token, ok = l.tokens[string(l.single_raw_charecter)]; ok {
		return true
	} else {
		return false
	}

}

func (l *Lexer) resetTokenValue() {
	l.single_token = token.Token{}
}

func (l *Lexer) createNumberToken() {
	l.single_token.Type = token.INT
	l.single_token.Literal = l.readNumber()
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
	if '0' <= l.single_raw_charecter && l.single_raw_charecter <= '9' {
		return true
	} else {
		return false
	}
}
