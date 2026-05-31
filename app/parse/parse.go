package parse

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type tokenizer struct {
	in    io.RuneReader
	token []rune
	err   error
}

type tokenizerState func(*tokenizer) tokenizerState

func (t *tokenizer) nextToken() string {
	t.token = t.token[:0]
	t.err = nil
	state := start
	for {
		state = state(t)
		if state == nil {
			return string(t.token)
		}
	}
}

func Split(in string) ([]string, error) {
	out := make([]string, 0, 1)

	t := tokenizer{in: strings.NewReader(in)}
	for tok := t.nextToken(); t.err != io.EOF; tok = t.nextToken() {
		if t.err != nil {
			return nil, fmt.Errorf("error tokenizing input: %w", t.err)
		}

		out = append(out, tok)
	}

	return out, nil
}

func start(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		t.err = err
		return nil
	}

	switch {
	case unicode.IsSpace(r):
		return start
	case r == '\'':
		return inSingleQuotes
	case r == '"':
		return inDoubleQuotes
	case r == '\\':
		return inBackslash
	default:
		t.token = append(t.token, r)
		return inArg
	}
}

func inArg(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		t.err = err
		return nil
	}

	switch {
	case unicode.IsSpace(r):
		return nil
	case r == '\'':
		return inSingleQuotes
	case r == '"':
		return inDoubleQuotes
	case r == '\\':
		return inBackslash
	default:
		t.token = append(t.token, r)
		return inArg
	}
}

var unterminatedSingleQuoteErr = errors.New("unterminated single quote")

func inSingleQuotes(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			t.err = unterminatedSingleQuoteErr
			return nil
		}

		t.err = err
		return nil
	}

	switch r {
	case '\'':
		return inArg
	default:
		t.token = append(t.token, r)
		return inSingleQuotes
	}
}

var unterminatedDoubleQuoteErr = errors.New("unterminated double quote")

func inDoubleQuotes(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			t.err = unterminatedDoubleQuoteErr
			return nil
		}

		t.err = err
		return nil
	}

	switch r {
	case '"':
		return inArg
	case '\\':
		t.token = append(t.token, r)
		return inBackslashInDoubleQuotes
	default:
		t.token = append(t.token, r)
		return inDoubleQuotes
	}
}

var missingEscapedCharErr = errors.New("missing escaped character after \\")

func inBackslash(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			t.err = missingEscapedCharErr
			return nil
		}

		t.err = err
		return nil
	}

	t.token = append(t.token, r)
	return inArg
}

func inBackslashInDoubleQuotes(t *tokenizer) tokenizerState {
	r, _, err := t.in.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			t.err = unterminatedDoubleQuoteErr
			return nil
		}

		t.err = err
		return nil
	}

	switch r {
	case '"', '\\':
		t.token = t.token[:len(t.token)-1]
		t.token = append(t.token, r)
		return inDoubleQuotes
	default:
		t.token = append(t.token, r)
		return inDoubleQuotes
	}
}
