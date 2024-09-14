package csskit

import (
	"bufio"
	"io"
	"strings"
	"errors"
)

var classAttrStart = []rune("class=\"")

func FromHTML(rd io.Reader) ([]string, error) {
	pit := newPeekIterator(bufio.NewReader(rd))

	const (
		stateClassAttr = iota
		stateOther
	)

	state := stateOther
	var sb strings.Builder
	var acc []string

	for {
		for state == stateClassAttr {
			c, _, err := pit.next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil, errors.New("unterminated class attr")
				}
				return nil, err
			}
			if c == '"' {
				acc = append(acc, sb.String())
				sb.Reset()
				state = stateOther
			} else {
				sb.WriteRune(c)
			}
		}

		err := pit.skipUntil(classAttrStart...)
		if err != nil {
			return nil, err
		}
		if pit.peekC == 0 {
			return acc, nil
		} else if pit.peekC != '"' {
			state = stateClassAttr
		}
	}
}

