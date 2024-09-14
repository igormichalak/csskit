package extract

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

type peekIterator struct {
	reader *bufio.Reader
	peekC  rune
}

func newPeekIterator(rd *bufio.Reader) *peekIterator {
	pit := &peekIterator{reader: rd, peekC: unicode.ReplacementChar}
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
		return errors.New("delimiter cannot be empty")
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
