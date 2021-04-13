package parser

import (
	"github.com/suisuss/monkey-interpreter/token"
	"github.com/suisuss/monkey-interpreter/lexer"
	"github.com/suisuss/monkey-interpreter/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token  // current token
	peekToken token.Token // next token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.peekToken()
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.peekToken()
}

func (p *Parser) ParserProgram() * ast.Program {
	return nil
}
