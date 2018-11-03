package lexer

import "github.com/yuseiatlas/lexer/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	lexer.ch = lexer.peekChar()
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		// sets ch to ASCII code of NUL (i.e. 0) to signify EOF
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}

}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case ':':
		tok = newToken(token.COLON, lexer.ch)
	case '.':
		tok = newToken(token.DOT, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '\'':
		tok = newToken(token.SINGLE_QUOTE, lexer.ch)
	case '"':
		tok = newToken(token.DOUBLE_QUOTE, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case '[':
		tok = newToken(token.LBRACKET, lexer.ch)
	case ']':
		tok = newToken(token.RBRACKET, lexer.ch)
	case '=':
		if lexer.peekChar() == '=' {
			current := lexer.ch
			lexer.readChar()
			tok = token.Token{Type: token.EQUAL, Literal: string(current) + string(lexer.ch)}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			current := lexer.ch
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQUAL, Literal: string(current) + string(lexer.ch)}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.GetIdentifierType(tok.Literal)

			// early return to avoid increasing read position
			return tok
		} else if isDigit(lexer.ch) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.NUMBER

			// early return to avoid increasing read position
			return tok
		} else {
			return newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	if lexer.ch == '.' {
		lexer.readChar()

		if isDigit(lexer.ch) {
			for isDigit(lexer.ch) {
				lexer.readChar()
			}
		} else {
			lexer.position--
			lexer.readPosition--
		}
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
