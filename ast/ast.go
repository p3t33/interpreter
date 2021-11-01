package ast

// variable bindings are statements let x = 5; - let statements
// let <identifier> = <expression>;

//few words about the difference between statements and expressions
// Expressions produce values, statements don’t.
// let x = 5 doesn’t produce a value whereas 5 does (the value it produces is 5 )
//A return 5; statement doesn’t produce a value, but add(5, 5) does.+

/*
An expression is something that can be reduced to a value,
for example "1+3" is an expression, but "foo = 1+3" is not.
*/

import (
	"github.com/p3t33/interpreter/token"
)

// let <identifier> = <expression>;
// let      x       =        7

// Token - token.LET
// Name - x - Identifier struct which holds the expression data
// Value - 7 - expression that produces the value.
type LetStatement struct {
	Token token.Token // - the token.LET token - the LetStatement node also needs to keep track of the token the AST node is associated with
	Name  *Identifier // - to hold the identifier of the binding
	Value Expression  // - for the expression that produces the value
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// <identifier>
// Token - token.IDEN
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This is the root of the AST
type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
