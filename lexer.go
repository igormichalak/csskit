package csskit

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	_ TokenType = iota
	TokenKeyword
	TokenNumber
	TokenUnit
	TokenHyphen
	TokenSpace
	TokenGarbage
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input    []rune
	inputLen int
	pos      int
	currChar rune
	peekChar rune
	prevTok  Token
}

func GetTokenTypeName(tt TokenType) string {
	switch tt {
	case TokenKeyword:
		return "keyword"
	case TokenNumber:
		return "number"
	case TokenUnit:
		return "unit"
	case TokenHyphen:
		return "hyphen"
	case TokenSpace:
		return "space"
	case TokenGarbage:
		return "garbage"
	case TokenEOF:
		return "EOF"
	default:
		panic(fmt.Errorf("unrecognized token type: %d", tt))
	}
}

func NewLexer(input string) *Lexer {
	inputRunes := []rune(input)
	l := &Lexer{input: inputRunes, inputLen: len(inputRunes)}
	if l.inputLen > 0 {
		l.currChar = l.input[0]
	}
	if l.inputLen > 1 {
		l.peekChar = l.input[1]
	}
	return l
}

func (l *Lexer) readChar() {
	l.currChar = l.peekChar
	l.pos++
	if l.pos+1 >= l.inputLen {
		l.peekChar = 0
	} else {
		l.peekChar = l.input[l.pos+1]
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token
	for {
		switch {
		case isLowerLetter(l.currChar):
			if l.prevTok.Type == TokenNumber {
				tok.Value = l.readUnit()
				tok.Type = TokenUnit
			} else {
				tok.Value = l.readKeyword()
				tok.Type = TokenKeyword
			}
		case isDigit(l.currChar):
			tok.Value = l.readNumber()
			tok.Type = TokenNumber
		case l.currChar == '%' && l.prevTok.Type == TokenNumber:
			tok = Token{Type: TokenUnit, Value: "%"}
			l.readChar()
		case l.currChar == '-':
			tok = Token{Type: TokenHyphen, Value: "-"}
			l.readChar()
		case unicode.IsSpace(l.currChar):
			tok = Token{Type: TokenSpace, Value: " "}
			l.readChar()
			for unicode.IsSpace(l.currChar) {
				l.readChar()
			}
		case l.currChar == 0:
			tok = Token{Type: TokenEOF, Value: ""}
		case l.prevTok.Type == TokenGarbage:
			l.readChar()
			continue
		default:
			tok = Token{Type: TokenGarbage, Value: ""}
			l.readChar()
		}

		l.prevTok = tok
		return tok
	}
}

func (l *Lexer) readKeyword() string {
	start := l.pos
	l.readChar()
	for isLowerLetter(l.currChar) {
		l.readChar()
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) readNumber() string {
	start := l.pos
	l.readChar()
	for isDigit(l.currChar) {
		l.readChar()
	}
	if l.currChar == '.' && isDigit(l.peekChar) {
		l.readChar()
		for isDigit(l.currChar) {
			l.readChar()
		}
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) readUnit() string {
	start := l.pos
	l.readChar()
	for isLowerLetter(l.currChar) {
		l.readChar()
	}
	return string(l.input[start:l.pos])
}

func isLowerLetter(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}
