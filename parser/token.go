package parser

type TokenType int

type Token struct {
	Kind  TokenType
	Value []byte
}

const (
	TokenEOF TokenType = iota
	TokenInner
	TokenTyp
	TokenPartnerId
	TokenPointsPbp
	TokenAmountPbp
	TokenPayment
	TokenGroupCode
)

const EOF = -1

var key = map[TokenType][]byte{
	TokenTyp:       []byte("<type>"),
	TokenPartnerId: []byte("<partnerId>"),
	TokenPointsPbp: []byte("<points_pbp>"),
	TokenAmountPbp: []byte("<amount_pbp>"),
	TokenPayment:   []byte("<payment>"),
	TokenGroupCode: []byte("<groupCode>"),
}

const (
	leftTeg  = "<"
	rightTeg = ">"
	endTeg   = "/>"
)
