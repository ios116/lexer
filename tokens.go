package lexer

import "fmt"

type itemType int

type item struct {
	typ itemType
	val string
}
const EOF  = 0
const (
	itemError itemType = iota

	itemDot
	itemEOF
	itemElse       // else keyword
	itemEnd        // end keyword
	itemField      // identifier, starting with '.'
	itemIdentifier // identifier
	itemIf         // if keyword
	itemLeftMeta   // left meta-string
	itemNumber     // number
	itemPipe       // pipe symbol
	itemRange      // range keyword
	itemRawString  // raw quoted string (includes quotes)
	itemRightMeta  // right meta-string
	itemString     // quoted string (includes quotes)
	itemText       // plain text
)

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}