package parse

import (
	"unicode"
	"unicode/utf8"
)

type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF
	tokenArg
)

type token struct {
	typ tokenType
	val string
}

type tokenizer struct {
	in    string
	start int
	pos   int
	token token
}

const eof = -1

func (t *tokenizer) next() rune {
	if t.pos >= len(t.in) {
		return eof
	}

	r, w := utf8.DecodeRuneInString(t.in[t.pos:])
	t.pos += w
	return r
}

type tokenizerState func(*tokenizer) tokenizerState

func (t *tokenizer) nextToken() token {
	state := start
	for {
		state = state(t)
		if state == nil {
			return t.token
		}
	}
}

func Split(in string) []string {
	out := make([]string, 0, 1)

	t := tokenizer{in: in}
	for tok := t.nextToken(); tok.typ != tokenEOF; tok = t.nextToken() {
		out = append(out, tok.val)
	}

	return out
}

func (t *tokenizer) emit(typ tokenType) tokenizerState {
	t.token.typ = typ
	t.start = t.pos
	return nil
}

func start(t *tokenizer) tokenizerState {
	r := t.next()
	switch {
	case r == eof:
		return t.emit(tokenEOF)
	case unicode.IsSpace(r):
		t.start = t.pos
		return start
	default:
		t.token.val = t.in[t.start:t.pos]
		return inArg
	}
}

func inArg(t *tokenizer) tokenizerState {
	r := t.next()
	switch {
	case r == eof || unicode.IsSpace(r):
		return t.emit(tokenArg)
	default:
		t.token.val = t.in[t.start:t.pos]
		return inArg
	}
}
