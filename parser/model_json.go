package parser

type DiscountRequest struct {
	LocationCode    string    `json:"locationCode"`
	ReceiptID       string    `json:"receiptId"`
	RUID            string    `json:"ruid"`
	TransactionDate string    `json:"transactionDate"`
	Cardno          string    `json:"cardno"`
	PartnerId       string    `json:"partnerId"`
	Products        []Product `json:"products"`
	Type            string    `json:"type"`
	Amount          float64   `json:"amount"`
	Payment         []Payment
	AmountPbp       string `xml:"amount_pbp"`
	PointsPbp       string `xml:"points_pbp"`
	TransactionHash string
	PosVersion      string `xml:"Pos_version"`
	ScanType        string `xml:"scan_type"`
	IpCashDesk      string `xml:"ip_cash_desk"`
}
type Product struct {
	Quantity float64 `json:"quantity" xml:"quantity"`
	// Указывает является ли товар промотоваром - если GroupCode == PROMO, то он исключается из обработки при ExcludePromo = true если значение REGULAR, то то это не промо товар
	GroupCode string `json:"groupCode" xml:"groupCode"`
	// id продукта
	Code   string  `json:"code" xml:"code"`
	Amount float64 `json:"amount" xml:"amount"`
}
type Payment struct {
	Type   string `xml:"type"`
	Amount int    `xml:"amount"`
}
