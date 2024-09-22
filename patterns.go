package csskit

type ClassPattern struct {
	Name     string
	Matchers []TokenMatcher
	UnitReq  bool
	Generate func(tokens []Token) ([]CSSProperty, error)
}

func literalMatcher(l string) TokenMatcher {
	return TokenMatcher{
		TokT:   TokenKeyword,
		ValT:   ValueFixed,
		Values: []string{l},
	}
}

func numberMatcher() TokenMatcher {
	return TokenMatcher{
		TokT:   TokenNumber,
		ValT:   ValueArbitrary,
		Values: []string{},
	}
}

func unitMatcher(units []string) TokenMatcher {
	return TokenMatcher{
		TokT:   TokenUnit,
		ValT:   ValueOneOf,
		Values: units,
	}
}

func hyphenMatcher() TokenMatcher {
	return TokenMatcher{
		TokT:   TokenHyphen,
		ValT:   ValueFixed,
		Values: []string{"-"},
	}
}

var classPatternCount int

func init() {
	classPatternCount = len(classPatterns)
}

var classPatterns = []ClassPattern{
	{
		Name: "width",
		Matchers: []TokenMatcher{
			literalMatcher("w"),
			hyphenMatcher(),
			numberMatcher(),
			unitMatcher([]string{"px", "%", "vw", "vh"}),
		},
		UnitReq: false,
		Generate: func(tokens []Token) ([]CSSProperty, error) {
			num := tokens[2].Value

			hasUnit := tokens[len(tokens)-1].Type == TokenUnit
			unit := ""
			if hasUnit {
				unit = tokens[3].Value
			} else {
				unit = "rem"
			}

			props := []CSSProperty{
				{Property: "width", Value: num + unit},
			}

			return props, nil
		},
	},
}
