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
				// inType= false
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
				// inType= false
				return res.Value
			}
		}
	}
	return nil
}

// RequestMock разбор xml с полями нужными mock сервису.
func RequestMock(b []byte) (*DiscountRequestXMLEnvelope, error) {
	request := &DiscountRequestXMLEnvelope{}
	_, items := Lex(b)
	payments := make([]Payment, 2)

	paymentCounter := -1
	status := map[string]bool{}
	for res := range items {
		if res.Kind == TokenEOF {
			break
		}
		switch {
		case res.Kind == StartElement:
			switch string(res.Value) {
			case "item":
				if status["payment"] {
					paymentCounter += 1
				}
				status["item"] = true
			default:
				status[string(res.Value)] = true
			}
		case res.Kind == EndElement:
			status[string(res.Value)] = false

		case res.Kind == CharData:
			switch {
			case status["type"] && !status["payment"]:
				request.Data.Type = string(res.Value)
			case status["partnerId"]:
				request.Data.PartnerId = string(res.Value)
			case status["locationCode"]:
				request.Data.LocationCode = string(res.Value)
			case status["terminalCode"]:
				request.Data.TerminalCode = string(res.Value)
			case status["transactionDate"]:
				request.Data.TransactionDate = string(res.Value)
			case status["receiptId"]:
				request.Data.ReceiptID = string(res.Value)
			case status["cardno"]:
				request.Data.Cardno = string(res.Value)
			case status["online"]:
				request.Data.Online = string(res.Value)
			case status["amount"] && !status["payment"] && !status["products"]:
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				request.Data.Amount = n

			case status["payment"] && status["item"] && status["type"]:
				payments[paymentCounter].Type = string(res.Value)

			case status["payment"] && status["item"] && status["amount"]:
				n, err := strconv.Atoi(string(res.Value))
				if err != nil {
					return nil, err
				}
				payments[paymentCounter].Amount = n
			case status["amount_pbp"]:
				request.Data.AmountPbp = string(res.Value)
			case status["points_pbp"]:
				request.Data.PointsPbp = string(res.Value)
			case status["pos_version"]:
				request.Data.PosVersion = string(res.Value)
			case status["ip_cash_desk"]:
				request.Data.IpCashDesk = string(res.Value)
			}
		}
	}
	request.Data.Payment.Item = payments

	return request, nil
}
