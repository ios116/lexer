package parser
// SoapRequest struct.
type SoapRequest struct {
	Data RequestData `xml:"Body>ProcessTransaction>RequestData"`
}

type RequestData struct {
	Type            string  `xml:"type"`
	TransactionDate string  `xml:"transactionDate"`
	ReceiptID       string  `xml:"receiptId"`
	Amount          float64 `xml:"amount"`
	AmountPbp       float64 `xml:"amount_pbp"`
	PointsPbp       int64   `xml:"points_pbp"`
}

type Product struct {
	GroupCode string  `xml:"groupCode"`
	Code      string  `xml:"code"`
	Quantity  float64 `xml:"quantity"`
	Amount    float64 `xml:"amount"`
}