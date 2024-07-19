package res_models

import "time"

// SearchShipment 结构体定义
type SearchShipment struct {
	ShipID       int       `json:"shipId"`
	RmbInv       string    `json:"rmbInv"`
	MasterPO     string    `json:"masterPo"`
	CustomerCode string    `json:"customerCode"`
	UbcPi        string    `json:"ubcPi"`
	MarkURL      string    `json:"markUrl"`
	OrigCountry  string    `json:"origCountry"`
	ShipMethod   string    `json:"shipMethod"`
	ShipTerm     string    `json:"shipTerm"`
	InvoiceTtl   float64   `json:"invoiceTtl"`
	ShipFrom     string    `json:"shipFrom"`
	MasterBLNum  string    `json:"masterBlNum"`
	HouseBLNum   string    `json:"houseBlNum"`
	Exporter     string    `json:"exporter"`
	ShipName     string    `json:"shipName"`
	PackDt       time.Time `json:"packDt"`
	ShipDt       time.Time `json:"shipDt"`
	ArriveDt     time.Time `json:"arriveDt"`
	Notes        string    `json:"notes"`
	GrossWeight  float64   `json:"grossWeight"`
	ItemCnt      int       `json:"itemCnt"`
	CartonCnt    int       `json:"cartonCnt"`
	CartonSize   float64   `json:"cartonSize"`
}
