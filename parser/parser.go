package parser

import (
	"../lexer"
	"../token"
)

// Parser strict
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New function
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken()
	p.nextToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
