package parser

import (
	"fmt"

	"github.com/p3t33/interpreter/ast"
	"github.com/p3t33/interpreter/lexer"
	"github.com/p3t33/interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors             []string
	program_statements ast.Program
	single_statement   ast.Statement
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.initialize_current_and_peek_tokens()

	program_statements := ast.Program{}
	program_statements.Statements = []ast.Statement{}

	return p
}

func (p *Parser) initialize_current_and_peek_tokens() {
	p.NextToken()
	p.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {

	for false == p.isEOFToken() {

		p.parseStatement()

		if true == p.isValidStatement() {
			p.appendStatementToAST()
		}

		p.NextToken()
	}

	return &p.program_statements
}

func (p *Parser) appendStatementToAST() {
	p.program_statements.Statements = append(p.program_statements.Statements, p.single_statement)
}

func (p *Parser) isValidStatement() bool {
	if p.single_statement != nil {
		return true
	} else {
		return false
	}
}
func (p *Parser) isEOFToken() bool {
	if token.EOF == p.curToken.Type {
		return true
	} else {
		return false
	}
}

func (p *Parser) parseStatement() {
	switch p.curToken.Type {
	case token.LET:
		p.single_statement = p.parseLetStatement()

	case token.RETURN:
		p.single_statement = p.parseReturnStatement()

	default:
		p.single_statement = nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	return_statement := p.createReturnStatement()

	p.NextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	p.skipExpressionUntilSemicolon()

	return return_statement
}

func (p *Parser) createReturnStatement() *ast.ReturnStatement {
	return &ast.ReturnStatement{Token: p.curToken}
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
		p.peekError(t)
		return false
	}
}
