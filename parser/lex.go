package parser

import "sync"

type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
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

var pool = sync.Pool{
	New: func() interface{} {
	    l:=new(lexer)
	    l.items =make(chan Token)
		return l
	},
}

func (l *lexer) reset()  {
	l.pos=0
	l.input=nil
	l.start = 0
	l.pos = 0
	l.width = 0
	l.line = 0
}

func Lex(input []byte) (*lexer, chan Token) {
    l:=pool.Get().(*lexer)
    l.input = input
	go l.run() // Concurrently run state machine.
	return l, l.items
}

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) run() {
	for state := Processor; state != nil; {
		state = state(l)
	}
	l.reset()
	pool.Put(l)
	//close(l.items) // No more tokens will be delivered.
}

func Processor(l *lexer) stateFn {
	switch {
	case l.input[l.pos] == openTag && l.input[l.pos+1] == slash:
		l.pos +=2
		l.ignore()
		return lexTagEnd

	case l.input[l.pos]== openTag:
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
		case l.input[l.pos] == closeTeg:
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
		case l.input[l.pos] == closeTeg:
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
		case l.input[l.pos] == openTag && l.input[l.pos+1] == slash:
			l.emit(CharData)
			l.pos +=2
			l.ignore()
			return lexTagEnd
		// если следующий тег а не
		case l.input[l.pos] == openTag:
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
