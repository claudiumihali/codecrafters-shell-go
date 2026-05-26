// Inspired from https://cs.opensource.google/go/go/+/refs/tags/go1.26.3:src/text/template/parse/lex.go
package lex

import (
	"unicode"
	"unicode/utf8"
)

// item represents a token returned by the lexer.
type item struct {
	typ itemType
	val string
}

type itemType int

const (
	itemError itemType = iota
	itemSpace
	itemEOF
	itemArg
)

const eof = -1

type lexer struct {
	input string
	start int
	pos   int
	atEOF bool
	item  item
}

// stateFn represents the state of the lexer as a function that returns the next
// state.
type stateFn func(*lexer) stateFn

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.atEOF = true
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])

	l.pos += w

	return r
}

func (l *lexer) backup() {
	if l.atEOF || l.pos <= 0 {
		return
	}

	_, w := utf8.DecodeLastRuneInString(l.input[:l.pos])

	l.pos -= w
}

func (l *lexer) peek() rune {
	r := l.next()

	l.backup()

	return r
}

func (l *lexer) thisItem(t itemType) item {
	i := item{
		typ: t,
		val: l.input[l.start:l.pos],
	}
	l.start = l.pos
	return i
}

func (l *lexer) emitItem(i item) stateFn {
	l.item = i
	return nil
}

func (l *lexer) nextItem() item {
	state := lexSpace
	for {
		state = state(l)
		if state == nil {
			return l.item
		}
	}
}

func lexSpace(l *lexer) stateFn {
	var r rune
	for {
		r = l.peek()
		if !unicode.IsSpace(r) {
			break
		}
		l.next()
	}

	if l.pos > l.start {
		return l.emitItem(l.thisItem(itemSpace))
	}

	if r == eof {
		return l.emitItem(l.thisItem(itemEOF))
	}

	return lexArg
}

func lexArg(l *lexer) stateFn {
	var r rune
	for {
		r = l.peek()
		if r == eof || unicode.IsSpace(r) {
			break
		}
		l.next()
	}

	if l.pos > l.start {
		return l.emitItem(l.thisItem(itemArg))
	}

	if r == eof {
		return l.emitItem(l.thisItem(itemEOF))
	}

	return lexSpace
}

func Parse(input string) []string {
	out := make([]string, 0, 1)

	l := &lexer{input: input}
	for i := l.nextItem(); i.typ != itemEOF; i = l.nextItem() {
		if i.typ == itemArg {
			out = append(out, i.val)
		}
	}

	return out
}
