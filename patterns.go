package csskit

import (
	"fmt"
	"strconv"
)

const (
	ColorSlate   = "slate"
	ColorGray    = "gray"
	ColorZinc    = "zinc"
	ColorNeutral = "neutral"
	ColorStone   = "stone"
	ColorRed     = "red"
	ColorOrange  = "orange"
	ColorAmber   = "amber"
	ColorYellow  = "yellow"
	ColorLime    = "lime"
	ColorGreen   = "green"
	ColorEmerald = "emerald"
	ColorTeal    = "teal"
	ColorCyan    = "cyan"
	ColorSky     = "sky"
	ColorBlue    = "blue"
	ColorIndigo  = "indigo"
	ColorViolet  = "violet"
	ColorPurple  = "purple"
	ColorFuchsia = "fuchsia"
	ColorPink    = "pink"
	ColorRose    = "rose"
)

var allColors = []string{
	ColorSlate, ColorGray, ColorZinc, ColorNeutral,
	ColorStone, ColorRed, ColorOrange, ColorAmber,
	ColorYellow, ColorLime, ColorGreen, ColorEmerald,
	ColorTeal, ColorCyan, ColorSky, ColorBlue,
	ColorIndigo, ColorViolet, ColorPurple, ColorFuchsia,
	ColorPink, ColorRose,
}

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

func colorMatcher() TokenMatcher {
	return TokenMatcher{
		TokT:   TokenKeyword,
		ValT:   ValueOneOf,
		Values: allColors,
	}
}

func getSizeValue(tokens []Token) (string, error) {
	tokenCount := len(tokens)
	lastToken := tokens[tokenCount-1]

	switch lastToken.Type {
	case TokenUnit:
		unit := lastToken.Value
		num := tokens[tokenCount-2].Value
		return num + unit, nil
	case TokenNumber:
		fl64, err := strconv.ParseFloat(lastToken.Value, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%.4frem", fl64/4.0), nil
	default:
		panic(fmt.Errorf("number token expected: %v", tokens))
	}
}

var classPatternCount int

func init() {
	classPatternCount = len(classPatterns)
}

var classPatterns = []ClassPattern{
	{
		Name: "Width",
		Matchers: []TokenMatcher{
			literalMatcher("w"),
			hyphenMatcher(),
			numberMatcher(),
			unitMatcher([]string{"px", "%", "vw", "vh"}),
		},
		UnitReq: false,
		Generate: func(tokens []Token) ([]CSSProperty, error) {
			val, err := getSizeValue(tokens)
			if err != nil {
				return nil, err
			}
			return []CSSProperty{
				{Property: "width", Value: val},
			}, nil
		},
	},
}

