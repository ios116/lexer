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
func RequestMock(b []byte) (*DiscountRequest, error) {
	request := &DiscountRequest{}
	_, items := Lex(b)
	payments := make([]Payment, 2)
	products := make([]Product, 0, 0)
	paymentCounter := -1
	productCounter := -1
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
				} else if status["products"] {
					productCounter += 1
					products = append(products, Product{})
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
				request.Type = string(res.Value)
			case status["partnerId"]:
				request.PartnerId = string(res.Value)
			case status["locationCode"]:
				request.LocationCode = string(res.Value)
			case status["transactionDate"]:
				request.TransactionDate = string(res.Value)
			case status["receiptId"]:
				request.ReceiptID = string(res.Value)
			case status["cardno"]:
				request.Cardno = string(res.Value)
				//
			case status["amount"] && !status["payment"] && !status["products"]:
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				request.Amount = n
			case status["payment"] && status["item"] && status["type"]:
				payments[paymentCounter].Type = string(res.Value)
			case status["payment"] && status["item"] && status["amount"]:
				n, err := strconv.Atoi(string(res.Value))
				if err != nil {
					return nil, err
				}
				payments[paymentCounter].Amount = n

			// парсинг продуктов
			case status["products"] && status["item"] && status["code"]:
				products[productCounter].Code = string(res.Value)
			case status["products"] && status["item"] && status["groupCode"]:
				products[productCounter].GroupCode = string(res.Value)
			case status["products"] && status["item"] && status["quantity"]:
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				products[productCounter].Quantity = n
			case status["products"] && status["item"] && status["amount"]:
				n, err := strconv.ParseFloat(string(res.Value), 64)
				if err != nil {
					return nil, err
				}
				products[productCounter].Amount = n

			case status["amount_pbp"]:
				request.AmountPbp = string(res.Value)
			case status["points_pbp"]:
				request.PointsPbp = string(res.Value)
			case status["pos_version"]:
				request.PosVersion = string(res.Value)
			case status["ip_cash_desk"]:
				request.IpCashDesk = string(res.Value)
			}
		}
	}
	request.Payment = payments
	request.Products = products

	return request, nil
}
