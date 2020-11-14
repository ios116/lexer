package parser

import (
	"bytes"
	"unicode/utf8"
)

// next returns the next rune in the input.
func (l *lexer) next() rune {
	r, s := utf8.DecodeRune(l.input[l.pos:])
	l.width = s
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}
	return r
}

// возвращает руту но откатывает счетчик назад
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

//  откатывает назад на руну
func (l *lexer) backup() {
	l.pos -= l.width
	// Correct newline count.
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(b []byte) bool {
	if bytes.ContainsRune(b,l.next()) {
		return true
	}
	l.backup()
	return false
}
