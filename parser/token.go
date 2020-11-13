package parser

type TokenType int

type Token struct {
	Kind  TokenType
	Value []byte
}

const (
	TokenEOF     TokenType = iota
	CharData               // inner text
	StartElement           // some tag
	EndElement             // some tag
)

const EOF = -1

const (
	openTag  = "<"
	closeTeg = ">"
	endTeg   = "</"
)
