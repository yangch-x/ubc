// Code generated by goctl. DO NOT EDIT.
package types

type SignInRequest struct {
	Name     string `json:"name,optional"`     // 邮箱
	Password string `json:"password,optional"` // 密码
}

type SignInResponse struct {
	Token        string `json:"token"`        // 访问令牌
	AccessExpire int64  `json:"accessExpire"` // 访问令牌过期时间
	RefreshAfter int64  `json:"refreshAfter"` // 刷新令牌时间
	Role         string `json:"role"`         // 角色
}

type Shipment struct {
	BillOfLanding   string `json:"billOfLanding,optional"`   // 提单
	CustomerCode    string `json:"customerCode,optional"`    // 客户代码
	ShipFrom        string `json:"shipFrom,optional"`        // 发货地
	ShipMethod      string `json:"shipMethod,optional"`      // 运输方式
	CountryOfOrigin string `json:"countryOfOrigin,optional"` // 原产国
	Manufacture     string `json:"manufacture,optional"`     // 制造商
	VesselFlight    string `json:"vesselFlight,optional"`    // 船/航班
	ETDDt           int64  `json:"etdDt,optional"`           // 预计离港日期
	UBCPI           string `json:"ubcPi,optional"`           // UBC PI
	MarksAndNumbers string `json:"marksAndNumbers,optional"` // 标记和号码
}

type Packing struct {
	PO     string  `json:"po,optional"`     // 采购订单号
	Style  string  `json:"style,optional"`  // 款式
	Color  string  `json:"color,optional"`  // 颜色
	ProjID int     `json:"projId,optional"` // projection id
	UBCPI  string  `json:"ubcPi,optional"`  // UBC PI
	Price  float64 `json:"price,optional"`  // 价格
	Carton int     `json:"carton,optional"` // 纸箱数量
	QtyPC  int     `json:"qtyPC,optional"`  // 数量（件）
	GWKG   int     `json:"gwKg,optional"`   // 毛重（千克）
	NWKG   int     `json:"nwKg,optional"`   // 净重（千克）
	Meas   int     `json:"meas,optional"`   // 尺寸
}

type Invoice struct {
	InvoiceCode    string  `json:"invoiceCode,optional"`    // 发票代码
	AdditionalCost float64 `json:"additionalCost,optional"` // 附加费用
	DepositAmt     float64 `json:"depositAmt,optional"`     // 定金金额
	InvoiceDt      int64   `json:"invoiceDt,optional"`      // 发票日期
	InvoiceDue     int64   `json:"invoiceDue,optional"`     // 发票到期日
}

type ShipmentAndInvoice struct {
	ShipmentId                int     `json:"shipmentId,optional"`
	InvoiceId                 int     `json:"invoiceId,optional"`
	BillOfLanding             string  `json:"billOfLanding,optional"`             // 提单
	ShipFrom                  string  `json:"shipFrom,optional"`                  // 发货地
	Manufacture               string  `json:"manufacture,optional"`               // 制造商
	CountryOfOrigin           string  `json:"countryOfOrigin,optional"`           // 原产国
	VesselFlight              string  `json:"vesselFlight,optional"`              // 船/航班
	UBCPI                     string  `json:"ubcPi,optional"`                     // UBC PI
	ETDDt                     string  `json:"etdDt,optional"`                     // 预计离港日期
	CustomerCode              string  `json:"customerCode,optional"`              // 客户代码
	ShipMethod                string  `json:"shipMethod,optional"`                // 运输方式
	InvoiceCode               string  `json:"invoiceCode,optional"`               // 发票代码
	AdditionalCost            float64 `json:"additionalCost,optional"`            // 附加费用
	AdditionalCostDescription string  `json:"additionalCostDescription,optional"` // 附加费用说明
	DepositAmt                float64 `json:"depositAmt,optional"`                // 定金金额
	InvoiceDt                 string  `json:"invoiceDt,optional"`                 // 发票日期
	InvoiceDue                string  `json:"invoiceDue,optional"`                // 发票到期日
}

type ShipmentAndInvoiceRes struct {
	ShipmentId int `json:"shipmentId"`
	InvoiceId  int `json:"invoiceId"`
}

type UpdateShippingInfo struct {
	BillOfLading    string `json:"billOfLading,optional"`    // 提单号
	CustomerCode    string `json:"customerCode,optional"`    // 客户代码
	MasterPO        string `json:"masterPO,optional"`        // 主采购订单号
	RMBInvoice      string `json:"rmbInvoice,optional"`      // 人民币发票
	ShipDate        string `json:"shipDate,optional"`        // 装运日期
	MarksAndNumbers string `json:"marksAndNumbers,optional"` // 标记和编号
	ShipFrom        string `json:"shipFrom,optional"`        // 发货地
	ShipMethod      string `json:"shipMethod,optional"`      // 运输方式
	CountryOfOrigin string `json:"countryOfOrigin,optional"` // 原产国
	Term            string `json:"term,optional"`            // 术语
	UBCPI           string `json:"ubcPI,optional"`           // UBC PI
	MasterBINum     string `json:"masterBINum,optional"`     // 主BI号
	Exporter        string `json:"exporter,optional"`        // 出口商
	ShipName        string `json:"shipName,optional"`        // 船名
	ShipDt          int64  `json:"shipDt,optional"`          // 装运日期
	ETDDt           int64  `json:"etdDt,optional"`           // 预计出发日期
	InvoiceTtl      string `json:"invoiceTtl,optional"`      // 发票标题
	Notes           string `json:"notes,optional"`           // 备注
}

type SaveShipment struct {
	ShipmentInfo Shipment `json:" shipmentAndInvoice,optional"`
}

type SaveOrUpdateShipment struct {
	ID           int     `json:"shipId"`
	RMBInv       string  `json:"rmbInv"`
	InvoiceTTL   float64 `json:"invoiceTtl"`
	CustomerCode string  `json:"customerCode"`
	MasterPO     string  `json:"masterPo"`
	ShipFrom     string  `json:"shipFrom"`
	UBCPI        string  `json:"ubcPi"`
	OrigCountry  string  `json:"origCountry"`
	ShipMethod   string  `json:"shipMethod"`
	ShipTerm     string  `json:"shipTerm"`
	MasterBLNum  string  `json:"masterBlNum"`
	HouseBLNum   string  `json:"houseBlNum"`
	Exporter     string  `json:"exporter"`
	ShipName     string  `json:"shipName"`
	ShipDT       string  `json:"shipDt"`
	ArriveDT     string  `json:"arriveDt"`
	Notes        string  `json:"notes"`
}

type SaveShipmentRes struct {
	ShinmentId string `json:"shinmentId"`
}

type UploadRes struct {
	Res interface{} `json:"res"`
}

type Query struct {
	PageNo       int    `json:"pageNo,default=0"`
	PageSize     int    `json:"pageSize,default=10"`
	SearchParams string `json:"searchParams,optional"`
	Order        string `json:"order,optional"`
}

type ListRes struct {
	Res      interface{} `json:"res"`
	Total    int         `json:"total"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
}

type PackingListReq struct {
	ShipmentId int `json:"shapmentId"`
	Total      int `json:"total"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
}

type UploadFile struct {
	UsedFor string `form:"usedFor"`
}

type UploadFileRes struct {
	Po        string `json:"customerPo"`
	Color     string `json:"color"`
	Style     string `json:"styleCode"`
	SalePrice string `json:"salePrice"`
	CartonCnt string `json:"cartonCnt"`
	MeasVol   string `json:"measVol"`
}

type SearchAllCustomerAndPrijection struct {
	Customers   interface{} `json:"customers"`
	Projections interface{} `json:"projections"`
}

type RemoveShipment struct {
	Id string `form:"shipId"`
}

type RemoveProjection struct {
	Id string `form:"projID"`
}

type RemoveCustomer struct {
	Id string `form:"customerID"`
}

type SaveProection struct {
	ProjID          int     `json:"projID,optional"`
	ArriveDt        string  `json:"arriveDt,optional"`
	UbcPi           string  `json:"ubcPi,optional"`
	FobLdp          string  `json:"fobLdp,optional"`
	CustomerCode    string  `json:"customerCode,optional"`
	Country         string  `json:"country,optional"`
	CustomerPo      string  `json:"customerPo,optional"`
	MasterPo        string  `json:"masterPo,optional"`
	StyleCode       string  `json:"styleCode,optional"`
	StyleName       string  `json:"styleName,optional"`
	Fabrication     string  `json:"fabrication,optional"`
	Color           string  `json:"color,optional"`
	Size            string  `json:"size,optional"`
	PoQty           int     `json:"poQty,optional"`
	ShipQty         int     `json:"shipQty,optional"`
	SalePrice       float64 `json:"salePrice,optional"`
	SaleCustPrice   float64 `json:"saleCustPrice,optional"`
	TtlBuy          float64 `json:"ttlBuy,optional"`
	TtlSell         float64 `json:"ttlSell,optional"`
	SaleCurrency    string  `json:"saleCurrency,optional"`
	InvoiceCode     string  `json:"invoiceCode,optional"`
	Receiving       string  `json:"receiving,optional"`
	Notes           string  `json:"notes,optional"`
	CostPrice       float64 `json:"costPrice,optional"`
	CostCurrency    string  `json:"costCurrency,optional"`
	RmbInv          string  `json:"rmbInv,optional"`
	Exporter        string  `json:"exporter,optional"`
	UbcPayable      float64 `json:"ubcPayable,optional"`
	PayPeriod       string  `json:"payPeriod,optional"`
	SalesPerson     string  `json:"salesPerson,optional"`
	SalesCommission float64 `json:"salesCommission,optional"`
	CommPaid        float64 `json:"commPaid,optional"`
}

type SaveOrUpdateCustomer struct {
	CustomerID      int    `json:"customerID,optional"`
	CustomerCode    string `json:"customerCode,optional"`
	CustomerEmail   string `json:"customerEmail,optional"`
	CustomerName    string `json:"customerName,optional"`
	BillingContact  string `json:"billingContact,optional"`
	NotifyContact   string `json:"notifyContact,optional"`
	PaymentTerm     string `json:"paymentTerm,optional"`
	ShipTo          string `json:"shipTo,optional"`
	SalesPerson     string `json:"salesPerson,optional"`
	UbcMerchandiser string `json:"ubcMerchandiser,optional"`
	Country         string `json:"country,optional"`
	DischargeLoc    string `json:"dischargeLoc,optional"`
	Status          string `json:"status,optional"`
	DueDateGap      int    `json:"dueDateGap,optional"`
}

type CreateInvoiceReq struct {
	Invoice  CreateInvoice   `json:"invoice"`
	Shipment CreateShipment  `json:"shipment"`
	Packings []CreatePacking `json:"packings"`
}

type CreateInvoiceRes struct {
	Res interface{} `json:"res"`
}

type CreateInvoice struct {
	SubTotal     float64 `json:"subTotal"`
	TotalCartons int     `json:"totalCartons"`
	TotalPCs     int     `json:"totalPCs"`
}

type CreateShipment struct {
	BillOfLanding             string  `json:"billOfLanding"`
	ShipFrom                  string  `json:"shipFrom"`
	Manufacture               string  `json:"manufacture"`
	CountryOfOrigin           string  `json:"countryOfOrigin"`
	VesselFlight              string  `json:"vesselFlight"`
	UbcPi                     string  `json:"ubcPi"`
	EtdDt                     string  `json:"etdDt"`
	CustomerCode              string  `json:"customerCode"`
	ShipMethod                string  `json:"shipMethod"`
	InvoiceCode               string  `json:"invoiceCode"`
	AdditionalCost            float64 `json:"additionalCost,optional"`
	AdditionalCostDescription string  `json:"additionalCostDescription,optional"`
	DepositAmt                float64 `json:"depositAmt,optional"`
	InvoiceDt                 string  `json:"invoiceDt,optional"`
	InvoiceDue                string  `json:"invoiceDue"`
	BillingContact            string  `json:"billingContact,optional"`
	ShipTo                    string  `json:"shipTo"`
	CustomerPos               string  `json:"customerPos"`
	Term                      string  `json:"term"`
}

type CreatePacking struct {
	ListId        int     `json:"listId,optional"`
	ProjId        int     `json:"projId"`
	CustomerPo    string  `json:"customerPo,optional"`
	StyleCode     string  `json:"styleCode,optional"`
	Color         string  `json:"color,optional"`
	SalePrice     float64 `json:"salePrice"`
	CartonCnt     int     `json:"cartonCnt"`
	TotalQuantity int     `json:"totalQuantity"`
	StyleName     string  `json:"styleName,optional"`
	Fabrication   string  `json:"fabrication,optional"`
	Size          string  `json:"size,optional"`
}

type CreatePackings struct {
	ShipId         int             `json:"shipId"`
	CreatePackings []CreatePacking `json:"createPackings"`
}

type QueryPacking struct {
	ShipId int `form:"shipId"`
}

type QueryPackingResp struct {
	Packings interface{} `json:"packings"`
	Shipment interface{} `json:"shipment"`
}
