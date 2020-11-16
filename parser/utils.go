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
func RequestMock(b []byte) (*DiscountRequestXMLEnvelope, error) {
	request := &DiscountRequestXMLEnvelope{}
	_, items := Lex(b)
	status := map[string]bool{}
	for res := range items {
		if res.Kind == TokenEOF {
			break
		}
		switch {
		case res.Kind == StartElement:
			switch string(res.Value) {
			case "type":
				status["type"] = true
			case "partnerId":
				status["partnerId"] = true
			case "locationCode":
				status["locationCode"] = true
			case "terminalCode":
				status["terminalCode"] = true
			case "transactionDate":
				status["transactionDate"] = true
			case "receiptId":
				status["receiptId"] = true
			case "cardno":
				status["cardno"] = true
			case "online":
				status["online"] = true
			case "amount":
				status["amount"] = true
			case "amount_pbp":
				status["amount_pbp"] = true
			case "points_pbp":
				status["points_pbp"] = true
			case "payment":
				status["payment"] = true
			case "pos_version":
				status["pos_version"] = true
			case "ip_cash_desk":
				status["ip_cash_desk"] = true
			case "products":
				status["products"] = true
			}

		case res.Kind == CharData:
			switch {
			case status["type"] && !status["payment"]:
				status["type"] = false
				request.Data.Type = string(res.Value)
			case status["partnerId"]:
				status["partnerId"] = false
				request.Data.PartnerId = string(res.Value)
			case status["locationCode"]:
				status["locationCode"] = false
				request.Data.LocationCode = string(res.Value)
			case status["terminalCode"]:
				status["terminalCode"] = false
				request.Data.TerminalCode = string(res.Value)
			case status["transactionDate"]:
				status["transactionDate"] = false
				request.Data.TransactionDate = string(res.Value)
			case status["receiptId"]:
				status["receiptId"] = false
				request.Data.ReceiptID = string(res.Value)
			case status["cardno"]:
				status["cardno"] = false
				request.Data.Cardno = string(res.Value)
			case status["online"]:
				status["online"] = false
				request.Data.Online = string(res.Value)
			case status["amount_pbp"]:
				status["amount_pbp"] = false
				request.Data.AmountPbp = string(res.Value)
			case status["points_pbp"]:
				status["points_pbp"] = false
				request.Data.PointsPbp = string(res.Value)
			case status["pos_version"]:
				status["pos_version"] = false
				request.Data.PosVersion = string(res.Value)
			case status["ip_cash_desk"]:
				status["ip_cash_desk"] = false
				request.Data.IpCashDesk = string(res.Value)
			case status["amount"] && !status["payment"]:
				status["amount"] = false
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				request.Data.Amount = n

			}
		}
	}

	return request, nil
}
