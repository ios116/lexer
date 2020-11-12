package parser

type TokenType int

type Token struct {
	Kind  TokenType
	Value []byte
}

const (
	TokenEOF      TokenType = iota
	TokenInner              // inner text
	TokenTagStart           // some tag
	TokenTagEnd             // some tag
)

const EOF = -1

const (
	leftTeg  = "<"
	rightTeg = ">"
	endTeg   = "/>"
)
