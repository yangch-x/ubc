package res_models

type PackingListRes struct {
	CustomerPo  string  `json:"customerPo"`
	StyleCode   string  `json:"styleCode"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
	UbcPi       string  `json:"ubcPi"`
	SalePrice   float64 `json:"salePrice"`
	ListID      int     `json:"listId"`
	ShipID      int     `json:"shipId"`
	ProjID      int     `json:"projId"`
	PackName    string  `json:"packName"`
	CartonCnt   int     `json:"cartonCnt"`
	ItemCnt     int     `json:"itemCnt"`
	MeasVol     float64 `json:"measVol"`
	GrossWeight float64 `json:"grossWeight"`
	NetWeight   float64 `json:"netWeight"`
	CartonSize  string  `json:"cartonSize"`
	PackCnt     int     `json:"packCnt"`
}

type NewShipmenPackingRes struct {
	Id            string
	CustomerPo    string
	StyleCode     string
	Color         string
	SalePrice     float64
	TotalQuantity float64
	CartonCnt     float64
}

type PackingsRes struct {
	Id            int
	ListId        int
	ProjId        int
	CustomerPo    string
	StyleCode     string
	StyleName     string
	Color         string
	SalePrice     float64
	TotalQuantity float64
	CartonCnt     float64
	Fabrication   string
}
