package parser

import (
	"github.com/p3t33/interpreter/ast"
	"github.com/p3t33/interpreter/lexer"
	"github.com/p3t33/interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for false == p.isEOFToken() {
		stnt := p.parseStatement()

		if stnt != nil {
			program.Statements = append(program.Statements, stnt)
		}

		p.NextToken()
	}

	return program
}

func (p *Parser) isEOFToken() bool {
	if token.EOF == p.curToken.Type {
		return true
	} else {
		return false
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()

	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	let_statement := p.createLetStatement()

	if false == p.isPeekedTokenAnIDENT() {
		return nil
	}

	p.NextToken()

	let_statement.Name = p.createLetStatementIdentifier()

	if false == p.isPeekedTokenAnASSIGN() {
		return nil
	}

	p.NextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	p.skipExpressionUntilSemicolon()

	return let_statement
}
func (p *Parser) createLetStatement() *ast.LetStatement {
	return &ast.LetStatement{Token: p.curToken}
}

func (p *Parser) createLetStatementIdentifier() *ast.Identifier {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) skipExpressionUntilSemicolon() {
	for false == p.isCurrentEXpressionSemicolon() {
		p.NextToken()
	}
}

func (p *Parser) isCurrentEXpressionSemicolon() bool {
	if true == p.isCurrentToken(token.SEMICOLON) {
		return true
	} else {
		return false
	}
}

func (p *Parser) isCurrentToken(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) isPeekedTokenAnIDENT() bool {
	if true == p.isPeekTokenExpected(token.IDENT) {
		return true
	} else {
		return false
	}
}

func (p *Parser) isPeekedTokenAnASSIGN() bool {
	if true == p.isPeekTokenExpected(token.ASSIGN) {
		return true
	} else {
		return false
	}
}

func (p *Parser) isPeekTokenExpected(t token.TokenType) bool {
	if p.isPeekToken(t) {
		return true
	} else {
		return false
	}
}
