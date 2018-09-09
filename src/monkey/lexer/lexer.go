package lexer

import "../../monkey/token"

/*
字句解析
 */
type Lexer struct {
	// 解析対象
	input        string
	// 現在位置
	position     int
	// 現在読み込み位置
	readPosition int
	// 読み込んだ文字
	ch           byte
}

/*
字句解析生成
 */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*

 */
func (l *Lexer) readChar() {
	// 現在の位置が文字数以上の場合、
	if l.readPosition >= len(l.input) {
		// 0を返す
		l.ch = 0
	} else {
		// 文字を返す
		l.ch = l.input[l.readPosition]
	}

	// 位置を更新
	l.position = l.readPosition
	// 読み込み位置を更新
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// 識別子の読み込み
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 文字列
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// トークンの生成
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 改行、空白、タブをスキップする
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 数値の読み込み
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 数値判定
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// のぞき見
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0

	} else {
		return l.input[l.readPosition]
	}
}
