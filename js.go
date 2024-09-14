package csskit

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var ErrUnterminatedString = errors.New("delimiter cannot be empty")

type peekIterator struct {
	reader *bufio.Reader
	peekC  rune
}

func newPeekIterator(rd *bufio.Reader) *peekIterator {
	pit := &peekIterator{reader: rd}
	pit.next()
	return pit
}

func (pit *peekIterator) next() (rune, rune, error) {
	if pit.peekC == 0 {
		return 0, 0, io.EOF
	}
	currentC := pit.peekC

	c, _, err := pit.reader.ReadRune()
	if errors.Is(err, io.EOF) {
		c = 0
	} else if err != nil {
		return 0, 0, err
	}
	pit.peekC = c

	return currentC, pit.peekC, nil
}

func (pit *peekIterator) skipLine() error {
	for {
		_, isPrefix, err := pit.reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if !isPrefix {
			break
		}
	}
	c, _, err := pit.reader.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			pit.peekC = 0
		} else {
			return err
		}
	} else {
		pit.peekC = c
	}
	return nil
}

func (pit *peekIterator) skipUntil(delim ...rune) error {
	if len(delim) == 0 {
		return ErrUnterminatedString
	}

	delimLen := len(delim)
	delimPtr := 0

	for {
		c, _, err := pit.next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if c == delim[delimPtr] {
			delimPtr++
			if delimPtr == delimLen {
				return nil
			}
		} else {
			delimPtr = 0
		}
	}
}

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
					return nil, errors.New("unterminated string")
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

