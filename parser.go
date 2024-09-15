package csskit

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

type ClassPattern struct {
	Name     string
	Matchers []TokenMatcher
	UnitReq  bool
	Generate func(tokens []Token) ([]CSSProperty, error)
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

func (p *Parser) parseClass(_ []Token) (CSSClass, error) {
	return CSSClass{}, nil
}

