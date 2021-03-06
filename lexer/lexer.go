package lexer

import (
	"github.com/suisuss/monkey-interpreter/token"
	// "fmt"
)

type Lexer struct {
	input string
	position int // Serves as a index for the current position in input (points to current char)
	readPosition int // Serves as a index for the current reading position in input (points to next char)
	ch byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input} // Create Lexer
	l.readChar() // Read in first character
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // l.ch equals 0 which is the ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 // l.ch equals 0 which is the ASCII code for "NUL"
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // Skip over any white space

	switch l.ch {
		case '=':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
			} else {
			tok = newToken(token.ASSIGN, l.ch) // newToken(TYPE, LITERAL) // LITERALL = Literal value
			}
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.BANG, l.ch) // newToken(TYPE, LITERAL) // LITERALL = Literal value
			}
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)			
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		// Not a recognised character. So must be a idenitifier/keyword
		default:
			if isLetter(l.ch) {
				// If l.ch is a letter we need to read the rest of the identifier/keyword until it encounters a non-letter-character
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdentifier(tok.Literal)
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	l.readChar() // read the next character
	return tok
}