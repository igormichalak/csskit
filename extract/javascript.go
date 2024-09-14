package extract

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var ErrUnterminatedString = errors.New("unterminated string")

func FromJS(rd io.Reader) ([]string, error) {
	pit := newPeekIterator(bufio.NewReader(rd))

	const (
		stateCode = iota
		stateString
	)

	state := stateCode
	var quoteC rune
	var sb strings.Builder
	var acc []string

	for {
		c, peekC, err := pit.next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if state == stateString {
					return nil, ErrUnterminatedString
				}
				return acc, nil
			} else {
				return nil, err
			}
		}

		switch state {
		case stateCode:
			if c == '"' || c == '\'' {
				state = stateString
				quoteC = c
			} else if c == '/' {
				if peekC == '/' {
					if err := pit.skipLine(); err != nil {
						return nil, err
					}
				} else if peekC == '*' {
					if err := pit.skipUntil('*', '/'); err != nil {
						return nil, err
					}
				}
			}
		case stateString:
			if c == '\\' {
				c2, _, err := pit.next()
				if err != nil {
					if errors.Is(err, io.EOF) {
						return nil, ErrUnterminatedString
					}
					return nil, err
				}
				_, err = sb.WriteRune('\\')
				if err != nil {
					return nil, err
				}
				_, err = sb.WriteRune(c2)
				if err != nil {
					return nil, err
				}
			} else if c == quoteC {
				acc = append(acc, sb.String())
				sb.Reset()
				state = stateCode
			} else {
				_, err := sb.WriteRune(c)
				if err != nil {
					return nil, err
				}
			}
		}
	}
}
