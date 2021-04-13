package ast

import "github.com/suisuss/monkey-interpreter/token"

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

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // Node.TokenLiteral()
	} else {
		return ""
	}
}

// Let statement node
type LetStatement struct {
	Token token.Token // the token.LET token
	Name *identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }


// Identifier node
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i * Identifier) expressionNode() {}
func (i * Identifier) TokenLiteral() string { return ls.Token.Literal }



