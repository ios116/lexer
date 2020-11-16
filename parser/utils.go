package parser

import (
	"bytes"
	"encoding/xml"
	"strconv"
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
				//inType= false
				return string(se)
			}
		}
	}

	return trTyp
}

func GetNewXml(b []byte, name []byte) []byte {
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
				//inType= false
				return res.Value
			}
		}
	}
	return nil
}

// RequestMock разбор xml с полями нужными mock сервису.
func RequestMock(b []byte) (*SoapRequest, error) {
	request := &SoapRequest{}
	_, items := Lex(b)
	var isType, isTransactionDate, isReceiptID, isPointsPbp, isAmount, isPayment, isAmountPbp bool
	for res := range items {
		if res.Kind == TokenEOF {
			break
		}
		switch {
		case res.Kind == StartElement:
			switch string(res.Value) {
			// amount_pbp
			case "amount_pbp":
				isAmountPbp = true
			case "points_pbp":
				isPointsPbp = true
			case "amount":
				isAmount = true
			case "receiptId":
				isReceiptID = true
			case "transactionDate":
				isTransactionDate = true
			case "type":
				isType = true
			case "payment":
				isPayment = true
			}

		case res.Kind == CharData:
			switch {
			case isPointsPbp:
				isPointsPbp = false
				n, err := strconv.ParseInt(string(res.Value), 10, 64)
				if err != nil {
					return nil, err
				}
				request.Data.PointsPbp = n
			case isAmount && !isPayment:
				isAmount = false
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				request.Data.Amount = n
			case isAmountPbp:
				isAmountPbp = false
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				request.Data.AmountPbp = n
			case isReceiptID:
				isReceiptID = false
				request.Data.ReceiptID = string(res.Value)
			case isTransactionDate:
				isTransactionDate = false
				request.Data.TransactionDate = string(res.Value)
			case isType && !isPayment:
				isType = false
				request.Data.Type = string(res.Value)
			}
		}
	}

	return request, nil
}