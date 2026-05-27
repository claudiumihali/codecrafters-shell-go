// Inspired from https://cs.opensource.google/go/go/+/refs/tags/go1.26.3:src/text/template/parse/lex.go
package lex

import (
	"unicode"
	"unicode/utf8"
)

func Split(in string) []string {
	out := make([]string, 0, 1)

	l := lexer{in: in}
	for t := l.nextToken(); t.typ != tokenEOF; t = l.nextToken() {
		out = append(out, t.val)
	}

	return out
}

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

type lexer struct {
	in    string
	start int
	pos   int
	token token
}

type state func(*lexer) state

const eof = -1

func (l *lexer) next() rune {
	if l.pos >= len(l.in) {
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.in[l.pos:])

	l.pos += w

	return r
}

func (l *lexer) backup() {
	if l.pos <= 0 {
		return
	}

	_, w := utf8.DecodeLastRuneInString(l.in[:l.pos])

	l.pos -= w
}

func (l *lexer) emit(typ tokenType) state {
	l.token.typ = typ
	l.token.val = l.in[l.start:l.pos]
	l.start = l.pos
	return nil
}

func (l *lexer) nextToken() token {
	state := lexArg
	for {
		state = state(l)
		if state == nil {
			return l.token
		}
	}
}

func lexArg(l *lexer) state {
	r := l.next()
	switch {
	case unicode.IsSpace(r):
		l.backup()

		if l.pos > l.start {
			return l.emit(tokenArg)
		}

		l.next()
		l.start = l.pos

		return lexArg
	case r == eof:
		if l.pos > l.start {
			return l.emit(tokenArg)
		}

		return l.emit(tokenEOF)
	default:
		return lexArg
	}
}
