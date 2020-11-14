package parser

import (
	"bytes"
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
	for state := lexTag; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}

// lexLeft scans until an opening action delimiter, "<".
//func lexLeft(l *lexer) stateFn {
//	if bytes.HasPrefix(l.input[l.pos:], []byte(leftTeg)) {
//		l.ignore()
//		return lexTag // Next state.
//	}
//	if l.next() == EOF {
//		l.emit(TokenEOF)
//	}
//	return nil // Stop the run loop.
//}

func lexTag(l *lexer) stateFn {
	for {
		r := l.next()
		if r == EOF {
			l.emit(TokenEOF)
			return nil
		}
		switch {
		case bytes.HasPrefix(l.input[l.pos:], key[TokenTyp]):
			l.start=l.pos
			l.pos = l.start + len(key[TokenTyp])
			l.emit(TokenTyp)
			return lexInner
		case bytes.HasPrefix(l.input[l.pos:], key[TokenPartnerId]):
			l.start=l.pos
			l.pos = l.start + len(key[TokenPartnerId])
			l.emit(TokenPartnerId)
			return lexInner
		case bytes.HasPrefix(l.input[l.pos:], key[TokenPointsPbp]):
			l.start=l.pos
			l.pos = l.start + len(key[TokenPointsPbp])
			l.emit(TokenPointsPbp)
			return lexInner
		case bytes.HasPrefix(l.input[l.pos:], key[TokenCard]):
			l.start=l.pos
			l.pos = l.start + len(key[TokenCard])
			l.emit(TokenCard)
			return lexInner
		default:
			l.ignore()
		}
	}
}

func lexInner(l *lexer) stateFn  {
	for {
		switch  {
		case bytes.HasPrefix(l.input[l.pos:],[]byte(leftTeg)):
			l.emit(TokenInner)
			return lexTag
		default:
			r := l.next()
			if r == EOF {
				l.emit(TokenEOF)
				return nil
			}
		}
	}
}
