package parser
//easyjson:skip
type DiscountRequestXMLEnvelope struct {
	Data DiscountRequestXML `xml:"Body>ProcessTransaction>RequestData"`
}

//easyjson:skip
type DiscountRequestXML struct {
	Text            string  `xml:",chardata"`
	Type            string  `xml:"type"`
	PartnerId       string  `xml:"partnerId"`
	RUID            string  `xml:"ruid"`
	LocationCode    string  `xml:"locationCode"`
	TerminalCode    string  `xml:"terminalCode"`
	TransactionDate string  `xml:"transactionDate"`
	ReceiptID       string  `xml:"receiptId"`
	Online          string  `xml:"online"`
	Cardno          string  `xml:"cardno"`
	Amount          float64 `xml:"amount"`
	AmountPbp       string  `xml:"amount_pbp"`
	PointsPbp       string  `xml:"points_pbp"`
	Products        struct {
		Text string                `xml:",chardata"`
		Item []DiscountRequestItem `xml:"item"`
	} `xml:"products"`
	PosVersion string `xml:"Pos_version"`
	ScanType   string `xml:"scan_type"`
	IpCashDesk string `xml:"ip_cash_desk"`
}

type DiscountRequestItem struct {
	Quantity float64 `json:"quantity" xml:"quantity"`
	// Указывает является ли товар промотоваром - если GroupCode == PROMO, то он исключается из обработки при ExcludePromo = true если значение REGULAR, то то это не промо товар
	GroupCode string `json:"groupCode" xml:"groupCode"`
	// id продукта
	Code        string  `json:"code" xml:"code"`
	Amount      float64 `json:"amount" xml:"amount"`
	ProductInfo *Product
}

type Product struct {
	Restricted  bool     `json:"restricted"` // является ли ограниченным (например, табаком)
	Mrprice     int      `json:"mrprice"`    // мрц в копейках
	ProductID   string   `json:"product_id"`
	Description string   `json:"description"`
	ParentCode  string   `json:"category_id"`
	Segments    []string `json:"product_segment_id"`
	Categories  []string `json:"categories"`
}