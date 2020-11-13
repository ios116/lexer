package parser

type TokenType int

type Token struct {
	Kind  TokenType
	Value []byte
}

const (
	TokenEOF     TokenType = iota
	StartElement           // some tag
	CharData               // inner text
	EndElement             // some tag
)

const EOF = -1

const (
	openTag  = "<"
	closeTeg = ">"
	endTeg   = "</"
)
