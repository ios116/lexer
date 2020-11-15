package parser

import (
	"bytes"
	"encoding/xml"
)

func GetFieldByNameFromXML(b []byte, name string) (trTyp string) {
	decoder := xml.NewDecoder(bytes.NewBuffer(b))
	inType := false
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			inType = se.Name.Local == name
		case xml.CharData:
			if inType {
				return string(se)
			}
		}
	}

	return trTyp
}

func GetNewXml(b []byte, name []byte) (result []byte) {
	_, items := Lex(b)
	inType := false
	for res := range items {
		if res.Kind == TokenEOF {
			break
		}
		switch {
		case res.Kind == StartElement:
			if 0 == bytes.Compare(res.Value, name) {
				inType = true
			}
		case res.Kind == CharData:
			if inType {
				return res.Value
			}
		}
	}
	return result

}
