package parser

import (
	"bytes"
	"fmt"
)

type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name  string     // used only for error reports.
	input []byte     // the string being scanned.
	start int        // start position of this item.
	pos   int        // current position in the input.
	width int        // width of last rune read from input.
	items chan Token // channel of scanned items.
	line  int        // 1+number of newlines seen
}

// emit passes an item back to the client.
func (l *lexer) emit(t TokenType) {
	l.items <- Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func Lex(name string, input []byte) (*lexer, chan Token) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan Token),
	}
	go l.run() // Concurrently run state machine.
	return l, l.items
}

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) run() {
	fmt.Println("total=", len(l.input))
	for state := Processor; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}

func Processor(l *lexer) stateFn {
	switch {
	case bytes.HasPrefix(l.input[l.pos:], []byte(endTeg)):
		l.next()
		l.next()
		l.ignore()
		return lexTagEnd

	case bytes.HasPrefix(l.input[l.pos:], []byte(openTag)):
		l.next()
		l.ignore()
		return lexTagStart



	default:
		r := l.next()
		if r == EOF {
			l.ignore()
			l.emit(TokenEOF)
			return nil
		}
		//fmt.Println("state=", l.pos, string(l.input[l.pos]))
		return Processor
	}
}

func lexTagEnd(l *lexer) stateFn {
	for {
		switch {
		case bytes.HasPrefix(l.input[l.pos:], []byte(closeTeg)):
			l.emit(EndElement)
			return Processor
		default:
			r := l.next()
			if r == EOF {
				l.ignore()
				l.emit(TokenEOF)
				return nil
			}
		}
	}
}

func lexTagStart(l *lexer) stateFn {
	for {
		switch {
		case bytes.HasPrefix(l.input[l.pos:], []byte(closeTeg)):
			l.emit(StartElement)
			l.next()
			l.ignore()
			return lexInner
		default:
			r := l.next()
			if r == EOF {
				l.ignore()
				l.emit(TokenEOF)
				return nil
			}
		}
	}
}

func lexInner(l *lexer) stateFn {
	for {
		switch {
		case bytes.HasPrefix(l.input[l.pos:], []byte(endTeg)):
			l.emit(CharData)
			l.next()
			l.next()
			l.ignore()
			return lexTagEnd
		// если следующий тег
		case bytes.HasPrefix(l.input[l.pos:], []byte(openTag)):
			l.next()
			l.ignore()
			return lexTagStart

		default:
			r := l.next()
			if r == EOF {
				l.ignore()
				l.emit(TokenEOF)
				return nil
			}
		}
	}
}
