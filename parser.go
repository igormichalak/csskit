package csskit

import (
	"fmt"
	"slices"
	"strings"
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

type CSSClass struct {
	Name  string
	Props []CSSProperty
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

func (p *Parser) Parse() ([]CSSClass, error) {
	var classes []CSSClass

	var tokens []Token
	var prevTok Token
	collecting := true

	for {
		tok := p.lexer.NextToken()

		if tok.Type == TokenEOF {
			if len(tokens) > 0 {
				if collecting && isValidLastToken(prevTok) {
					class, err := p.parseClass(tokens)
					if err != nil {
						return nil, err
					}
					classes = append(classes, class)
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
				tokens = tokens[:0]
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

func (p *Parser) parseClass(tokens []Token) (CSSClass, error) {
	for i := 0; i < classPatternsCount; i++ {
		pattern := &classPatterns[i]
		if !matchPattern(pattern, tokens) {
			continue
		}
		name, err := joinTokens(tokens)
		if err != nil {
			return CSSClass{}, err
		}
		props, _ := pattern.Generate(tokens)
		return CSSClass{Name: name, Props: props}, nil
	}
	return CSSClass{}, fmt.Errorf("no matching pattern for tokens: %v", tokens)
}

func joinTokens(tokens []Token) (string, error) {
	var sb strings.Builder
	for _, tok := range tokens {
		if _, err := sb.WriteString(tok.Value); err != nil {
			return "", err
		}
	}
	return sb.String(), nil
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
			targetCount = matcherCount-1
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

