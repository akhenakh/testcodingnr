// package wordsplit is splitting input into 3 words triplet
package wordsplit

import (
	"bufio"
	"context"
	"errors"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Triplet [3]string

type TripletIterator struct {
	scanner *bufio.Scanner
	words   []string
}

var Done = errors.New("no more items")

func ParseAndSplit(ctx context.Context, r io.Reader) *TripletIterator {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanWords)
	return &TripletIterator{scanner: scanner}
}

func (t *TripletIterator) Next() (Triplet, error) {
	if !t.scanner.Scan() {
		return Triplet{}, Done
	}
	t.words = append(t.words, strings.ToLower(t.scanner.Text()))

	if len(t.words) == 1 {
		for i := 0; i < 2; i++ {
			if !t.scanner.Scan() {
				return Triplet{}, Done
			}
			t.words = append(t.words, strings.ToLower(t.scanner.Text()))
		}

		return Triplet{t.words[0], t.words[1], t.words[2]}, nil
	}

	_, t.words = t.words[0], t.words[1:]

	return Triplet{t.words[0], t.words[1], t.words[2]}, nil
}

// isSpaceOrPunct reports whether the character is a Unicode white space character
// or punctuation as defined by unicode.IsSpace or unicode.IsPunct.
func isSpaceOrPunct(r rune) bool {
	return IsPunct(r) || unicode.IsSpace(r)
}

// IsPunct reports whether the rune is a Unicode punctuation character excluding \'.
func IsPunct(r rune) bool {
	if r == '\'' {
		return false
	}
	return unicode.IsPunct(r)
}

// ScanWords is a split function for a Scanner that returns each
// space-separated word of text, with surrounding spaces & puntucation deleted. It will
// never return an empty string.
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpaceOrPunct(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSpaceOrPunct(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
