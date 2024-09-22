package csskit

import (
	"fmt"
	"slices"
)

type ValueType int

const (
	_ ValueType = iota
	ValueFixed
	ValueOneOf
	ValueArbitrary
)

type TokenMatcher struct {
	TokT   TokenType
	ValT   ValueType
	Values []string
}

type RawCSSClass struct {
	Tokens []Token
	Props  []CSSProperty
}

type CSSProperty struct {
	Property string
	Value    string
}

type Parser struct {
	lexer *Lexer
}

func NewParser(lex *Lexer) *Parser {
	return &Parser{lexer: lex}
}

func (p *Parser) Parse() ([]RawCSSClass, error) {
	var classes []RawCSSClass

	var tokens []Token
	var prevTok Token
	collecting := true

	for {
		tok := p.lexer.NextToken()

		if tok.Type == TokenEOF {
			if len(tokens) > 0 {
				if collecting && isValidLastToken(prevTok) {
					class, err := p.parseClass(tokens)
					if err == nil {
						classes = append(classes, class)
					}
				}
			}
			break
		}

		if tok.Type == TokenSpace {
			if len(tokens) > 0 {
				if collecting && isValidLastToken(prevTok) {
					class, err := p.parseClass(tokens)
					if err == nil {
						classes = append(classes, class)
					}
				}
				tokens = nil
				prevTok = Token{}
			}
			collecting = true
			continue
		}

		switch tok.Type {
		case TokenKeyword:
			if prevTok.Type != TokenHyphen && len(tokens) > 0 {
				collecting = false
				continue
			}
		case TokenNumber:
			if prevTok.Type != TokenHyphen {
				collecting = false
				continue
			}
		case TokenUnit:
			if prevTok.Type != TokenNumber {
				collecting = false
				continue
			}
		case TokenHyphen:
			if prevTok.Type != TokenKeyword {
				collecting = false
				continue
			}
		case TokenGarbage:
			collecting = false
			continue
		}

		if collecting {
			tokens = append(tokens, tok)
			prevTok = tok
		}
	}

	return classes, nil
}

func isValidLastToken(tok Token) bool {
	switch tok.Type {
	case TokenKeyword, TokenNumber, TokenUnit:
		return true
	default:
		return false
	}
}

func (p *Parser) parseClass(tokens []Token) (RawCSSClass, error) {
	for i := 0; i < classPatternCount; i++ {
		pattern := &classPatterns[i]
		if !matchPattern(pattern, tokens) {
			continue
		}
		props, err := pattern.Generate(tokens)
		if err != nil {
			return RawCSSClass{}, fmt.Errorf("generation error: %v", err)
		}
		return RawCSSClass{Tokens: tokens, Props: props}, nil
	}
	return RawCSSClass{}, fmt.Errorf("no matching pattern for tokens: %v", tokens)
}

func matchPattern(pattern *ClassPattern, tokens []Token) bool {
	matcherCount := len(pattern.Matchers)
	tokenCount := len(tokens)
	lastMatcherType := pattern.Matchers[matcherCount-1].TokT
	lastTokenType := tokens[tokenCount-1].Type

	if pattern.UnitReq || lastMatcherType != TokenUnit {
		if matcherCount != tokenCount {
			return false
		}
	} else {
		targetCount := -1
		if lastTokenType == TokenUnit {
			targetCount = matcherCount
		} else {
			targetCount = matcherCount - 1
		}
		if targetCount != tokenCount {
			return false
		}
	}

	for matcherIdx, matcher := range pattern.Matchers {
		if matcherIdx >= tokenCount {
			return matcher.TokT == TokenUnit && matcherIdx == matcherCount-1 && !pattern.UnitReq
		}

		tok := tokens[matcherIdx]

		if tok.Type != matcher.TokT {
			return false
		}

		switch matcher.ValT {
		case ValueFixed:
			if tok.Value != matcher.Values[0] {
				return false
			}
		case ValueOneOf:
			if !slices.Contains(matcher.Values, tok.Value) {
				return false
			}
		case ValueArbitrary:
		}
	}

	return true
}
