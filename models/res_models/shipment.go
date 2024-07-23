package res_models

// SearchShipment 结构体定义
type SearchShipment struct {
	ShipID       int     `json:"shipId"`
	RmbInv       string  `json:"rmbInv"`
	MasterPO     string  `json:"masterPo"`
	CustomerCode string  `json:"customerCode"`
	UbcPi        string  `json:"ubcPi"`
	MarkURL      string  `json:"markUrl"`
	OrigCountry  string  `json:"origCountry"`
	ShipMethod   string  `json:"shipMethod"`
	ShipTerm     string  `json:"shipTerm"`
	InvoiceTtl   float64 `json:"invoiceTtl"`
	ShipFrom     string  `json:"shipFrom"`
	MasterBLNum  string  `json:"masterBlNum"`
	HouseBLNum   string  `json:"houseBlNum"`
	Exporter     string  `json:"exporter"`
	ShipName     string  `json:"shipName"`
	PackDt       string  `json:"packDt"`
	ShipDt       string  `json:"shipDt"`
	ArriveDt     string  `json:"arriveDt"`
	Notes        string  `json:"notes"`
	GrossWeight  float64 `json:"grossWeight"`
	ItemCnt      int     `json:"itemCnt"`
	CartonCnt    int     `json:"cartonCnt"`
	CartonSize   float64 `json:"cartonSize"`
}

type ShipmentAndInvoice struct {
	ShipID          int     `json:"shipId"`
	RmbInv          string  `json:"rmbInv"`
	MasterPO        string  `json:"masterPo"`
	CustomerCode    string  `json:"customerCode"`
	UbcPi           string  `json:"ubcPi"`
	OrigCountry     string  `json:"origCountry"`
	ShipMethod      string  `json:"shipMethod"`
	ShipTerm        string  `json:"shipTerm"`
	InvoiceTtl      float64 `json:"invoiceTtl"`
	ShipFrom        string  `json:"shipFrom"`
	MasterBLNum     string  `json:"masterBlNum"`
	HouseBLNum      string  `json:"houseBlNum"`
	Exporter        string  `json:"exporter"`
	ShipName        string  `json:"shipName"`
	PackDt          string  `json:"packDt"`
	ShipDt          string  `json:"shipDt"`
	ArriveDt        string  `json:"arriveDt"`
	GrossWeight     float64 `json:"grossWeight"`
	ItemCnt         int     `json:"itemCnt"`
	CartonCnt       int     `json:"cartonCnt"`
	CartonSize      float64 `json:"cartonSize"`
	InvoiceID       int     `gorm:"column:invoice_id;primaryKey;autoIncrement"`
	InvoiceCode     string  `gorm:"column:invoice_code;size:100;unique;not null"`
	InvoiceAmt      int     `gorm:"column:invoice_amt"`
	ReceivedAmt     int     `gorm:"column:received_amt"`
	InvoiceDt       string  `gorm:"column:invoice_dt"`
	InvoiceDue      string  `gorm:"column:invoice_due"`
	InvoiceCurrency string  `gorm:"column:invoice_currency;size:100;default:USD;not null"`
	Notes           string  `gorm:"column:notes;type:text"`
	ShipTo          string  `json:"ship_to"`
	BillingContact  string  `json:"billing_contact"`
	Term            string  `json:"term"`
}

type ShipAndInvoice struct {
	ShipID         int     `json:"shipId"`
	InvoiceId      int     `json:"invoiceId"`
	MasterPO       string  `json:"masterPo"`
	CustomerCode   string  `json:"customerCode"`
	UbcPi          string  `json:"ubcPi"`
	OrigCountry    string  `json:"countryOfOrigin"`
	ShipMethod     string  `json:"shipMethod"`
	ShipTerm       string  `json:"shipTerm"`
	InvoiceTtl     float64 `json:"invoiceTtl"`
	ShipFrom       string  `json:"shipFrom"`
	HouseBlNum     string  `json:"billOfLanding"`
	Manufacture    string  `json:"manufacture"`
	ShipName       string  `json:"vesselFlight"`
	ShipDt         string  `json:"shipDt"`
	InvoiceCode    string  `json:"invoiceCode"`
	InvoiceAmt     int     `json:"invoiceAmt"`
	ReceivedAmt    int     `json:"receivedAmt"`
	InvoiceDt      string  `json:"invoiceDt"`
	InvoiceDue     string  `json:"invoiceDue"`
	ShipTo         string  `json:"ship_to"`
	BillingContact string  `json:"billing_contact"`
	Term           string  `json:"term"`
}
