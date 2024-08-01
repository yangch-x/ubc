package models

import (
	"UBC/models/res_models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Shipment struct {
	ShipID       int     `json:"ship_id" gorm:"column:ship_id;primaryKey;autoIncrement"`
	RmbInv       string  `json:"rmb_inv" gorm:"column:rmb_inv"`
	MasterPo     string  `json:"master_po" gorm:"column:master_po"`
	CustomerCode string  `json:"customer_code" gorm:"column:customer_code"`
	UbcPi        string  `json:"ubc_pi" gorm:"column:ubc_pi"`
	Markurl      string  `json:"markurl" gorm:"column:markurl"`
	OrigCountry  string  `json:"orig_country" gorm:"column:orig_country"`
	ShipMethod   string  `json:"ship_method" gorm:"column:ship_method"`
	ShipTerm     string  `json:"ship_term" gorm:"column:ship_term"`
	InvoiceTtl   float64 `json:"invoice_ttl" gorm:"column:invoice_ttl"`
	ShipFrom     string  `json:"ship_from" gorm:"column:ship_from"`
	MasterBlNum  string  `json:"master_bl_num" gorm:"column:master_bl_num;notNull"`
	HouseBlNum   string  `json:"house_bl_num" gorm:"column:house_bl_num"`
	Exporter     string  `json:"exporter" gorm:"column:exporter"`
	ShipName     string  `json:"ship_name" gorm:"column:ship_name"`
	PackDt       string  `json:"pack_dt" gorm:"column:pack_dt"`
	ShipDt       string  `json:"ship_dt" gorm:"column:ship_dt"`
	ArriveDt     string  `json:"arrive_dt" gorm:"column:arrive_dt"`
	Notes        string  `json:"notes" gorm:"column:notes;type:text"`
}

type Config struct {
	ConfigID     int     `gorm:"column:config_id;primaryKey"`
	UbcName      string  `gorm:"column:ubc_name;size:100;not null"`
	UbcAddress   string  `gorm:"column:ubc_address;size:255;not null"`
	QimeiName    string  `gorm:"column:qimei_name;size:100;not null"`
	QimeiAddress string  `gorm:"column:qimei_address;size:255;not null"`
	QimeiTaxno   string  `gorm:"column:qimei_taxno;size:100;not null"`
	QimeiCompno  string  `gorm:"column:qimei_compno;size:100;not null"`
	RmbRate      float64 `gorm:"column:rmb_rate;not null"`
	EuroRate     float64 `gorm:"column:euro_rate"`
	JpyRate      float64 `gorm:"column:jpy_rate"`
	GbpRate      float64 `gorm:"column:gbp_rate"`
}

type ConfigHist struct {
	ChgID        int       `gorm:"column:chg_id;primaryKey;autoIncrement"`
	ChgDt        time.Time `gorm:"column:chg_dt;not null"`
	ConfigID     int       `gorm:"column:config_id;not null"`
	UbcName      string    `gorm:"column:ubc_name;size:100;not null"`
	UbcAddress   string    `gorm:"column:ubc_address;size:255;not null"`
	QimeiName    string    `gorm:"column:qimei_name;size:100;not null"`
	QimeiAddress string    `gorm:"column:qimei_address;size:255;not null"`
	QimeiTaxno   string    `gorm:"column:qimei_taxno;size:100;not null"`
	QimeiCompno  string    `gorm:"column:qimei_compno;size:100;not null"`
	RmbRate      float64   `gorm:"column:rmb_rate;not null"`
	EuroRate     float64   `gorm:"column:euro_rate"`
	JpyRate      float64   `gorm:"column:jpy_rate"`
	GbpRate      float64   `gorm:"column:gbp_rate"`
}

type Customer struct {
	CustomerID      int    `gorm:"column:customer_id;primaryKey;autoIncrement"`
	CustomerCode    string `gorm:"column:customer_code;size:255;unique;not null"`
	CustomerEmail   string `gorm:"column:customer_email;size:255"`
	CustomerName    string `gorm:"column:customer_name;size:255"`
	BillingContact  string `gorm:"column:billing_contact;type:text"`
	NotifyContact   string `gorm:"column:notify_contact;type:text"`
	PaymentTerm     string `gorm:"column:payment_term;size:25"`
	ShipTo          string `gorm:"column:ship_to;type:text"`
	SalesPerson     string `gorm:"column:sales_person;size:255"`
	UbcMerchandiser string `gorm:"column:ubc_merchandiser;size:255"`
	Country         string `gorm:"column:country;size:255"`
	DischargeLoc    string `gorm:"column:discharge_loc;size:255"`
	Status          string `gorm:"column:status;size:25"`
	DueDateGap      int    `gorm:"column:due_date_gap"`
}

type HSCode struct {
	HsID          int    `gorm:"column:hs_id;primaryKey;autoIncrement"`
	HsCode        string `gorm:"column:hs_code;size:100;not null"`
	HtsCode       string `gorm:"column:hts_code;size:100;not null"`
	DescrEn       string `gorm:"column:descr_en;size:255;not null"`
	DescrCn       string `gorm:"column:descr_cn;size:255;not null"`
	CustomFactors string `gorm:"column:custom_factors;size:255;not null"`
	Notes         string `gorm:"column:notes;type:text"`
}

type HSCodeHist struct {
	ChgID         int       `gorm:"column:chg_id;primaryKey;autoIncrement"`
	ChgDt         time.Time `gorm:"column:chg_dt;not null"`
	HsCode        string    `gorm:"column:hs_code;size:100;not null"`
	HtsCode       string    `gorm:"column:hts_code;size:100;not null"`
	DescrEn       string    `gorm:"column:descr_en;size:255;not null"`
	DescrCn       string    `gorm:"column:descr_cn;size:255;not null"`
	CustomFactors string    `gorm:"column:custom_factors;size:255;not null"`
	Notes         string    `gorm:"column:notes;type:text"`
}

type Invoice struct {
	InvoiceID                 int     `gorm:"column:invoice_id;primaryKey;autoIncrement"`
	ShipID                    int     `gorm:"column:ship_id"`
	InvoiceCode               string  `gorm:"column:invoice_code;size:100;unique;not null"`
	UbcPi                     string  `gorm:"column:ubc_pi;size:100;not null"`
	CustomerCode              string  `gorm:"column:customer_code;size:255;not null"`
	InvoiceAmt                float64 `gorm:"column:invoice_amt"`
	ReceivedAmt               float64 `gorm:"column:received_amt"`
	AdditionalCost            float64 `gorm:"column:additional_cost"`
	AdditionalCostDescription string  `gorm:"column:additional_cost_description"`
	InvoiceDt                 string  `gorm:"column:invoice_dt"`
	InvoiceDue                string  `gorm:"column:invoice_due"`
	InvoiceCurrency           string  `gorm:"column:invoice_currency;size:100;default:USD;not null"`
	Notes                     string  `gorm:"column:notes;type:text"`
}

type PO struct {
	PoID        int     `gorm:"column:po_id;primaryKey;autoIncrement"`
	CustomerPo  string  `gorm:"column:customer_po;size:255;unique;not null"`
	StyleCode   string  `gorm:"column:style_code;size:100;not null"`
	StyleColor  string  `gorm:"column:style_color;size:100;not null"`
	StyleSize   string  `gorm:"column:style_size;size:100;not null"`
	OrderDt     string  `gorm:"column:order_dt"`
	OrderQty    int     `gorm:"column:order_qty"`
	ShipQty     int     `gorm:"column:ship_qty"`
	UnitPrice   float64 `gorm:"column:unit_price"`
	SalesPrice  float64 `gorm:"column:sales_price"`
	CustomPrice float64 `gorm:"column:custom_price"`
	CostPrice   float64 `gorm:"column:cost_price"`
	Notes       string  `gorm:"column:notes;type:text"`
}

type Packing struct {
	PackID      int     `gorm:"column:pack_id;primaryKey;autoIncrement"`
	ShipID      int     `gorm:"column:ship_id"`
	InvoiceID   int     `gorm:"column:invoice_id"`
	PackName    string  `gorm:"column:pack_name;size:255;not null"`
	StyleCode   string  `gorm:"column:style_code;size:255;not null"`
	PoCode      string  `gorm:"column:po_code;size:250;not null"`
	CartonCode  string  `gorm:"column:carton_code;size:50;not null"`
	BeginNum    int     `gorm:"column:begin_num;not null"`
	EndNum      int     `gorm:"column:end_num;not null"`
	CartonCnt   int     `gorm:"column:carton_cnt;not null"`
	PackColor   string  `gorm:"column:pack_color;size:50;not null"`
	PackSize    string  `gorm:"column:pack_size;size:50;not null"`
	ItemCnt     int     `gorm:"column:item_cnt"`
	GrossWeight float64 `gorm:"column:gross_weight"`
	NetWeight   float64 `gorm:"column:net_weight"`
	CartonSize  string  `gorm:"column:carton_size;size:255"`
	LineCnt     int16   `gorm:"column:line_cnt"`
}

type PackingList struct {
	ListID        int     `gorm:"column:list_id;primaryKey;autoIncrement"`
	ShipID        int     `gorm:"column:ship_id;not null"`
	ProjID        int     `gorm:"column:proj_id"`
	SalePrice     float64 `json:"sale_price"`
	CartonCnt     int     `json:"carton_cnt"`
	TotalQuantity int     `json:"total_quantity"`
}

type Projection struct {
	ProjID          int     `gorm:"column:proj_id"`
	ArriveDt        string  `gorm:"column:arrive_dt;not null"`
	UbcPi           string  `gorm:"column:ubc_pi;size:100;not null"`
	FobLdp          string  `gorm:"column:fob_ldp;size:25;not null"`
	CustomerCode    string  `gorm:"column:customer_code;size:255;not null"`
	Country         string  `gorm:"column:country;size:100;not null"`
	CustomerPo      string  `gorm:"column:customer_po;size:100;not null;primaryKey"`
	MasterPo        string  `gorm:"column:master_po;size:100;not null"`
	StyleCode       string  `gorm:"column:style_code;size:100;not null;primaryKey"`
	StyleName       string  `gorm:"column:style_name;size:255;not null"`
	Fabrication     string  `gorm:"column:fabrication;size:255;not null"`
	Color           string  `gorm:"column:color;size:255;not null;primaryKey"`
	Size            string  `gorm:"column:size;size:255;not null"`
	PoQty           int     `gorm:"column:po_qty"`
	ShipQty         int     `gorm:"column:ship_qty"`
	SalePrice       float64 `gorm:"column:sale_price"`
	TtlBuy          float64 `gorm:"column:ttl_buy"`
	TtlSell         float64 `gorm:"column:ttl_sell"`
	SaleCustPrice   float64 `gorm:"column:sale_cust_price"`
	SaleCurrency    string  `gorm:"column:sale_currency;size:100;default:USD;not null"`
	InvoiceCode     string  `gorm:"column:invoice_code;size:100;not null"`
	Receiving       string  `gorm:"column:receiving;size:255;not null"`
	Notes           string  `gorm:"column:notes;size:255;not null"`
	CostPrice       float64 `gorm:"column:cost_price"`
	CostCurrency    string  `gorm:"column:cost_currency;size:100;default:RMB;not null"`
	RmbInv          string  `gorm:"column:rmb_inv;size:100;not null"`
	Exporter        string  `gorm:"column:exporter;size:100;not null"`
	UbcPayable      float64 `gorm:"column:ubc_payable"`
	PayPeriod       string  `gorm:"column:pay_period;size:100;not null"`
	SalesPerson     string  `gorm:"column:sales_person;size:100;not null"`
	SalesCommission float64 `gorm:"column:sales_commission"`
	CommPaid        float64 `gorm:"column:comm_paid"`
}

type Style struct {
	StyleID      int     `gorm:"column:style_id;primaryKey;autoIncrement"`
	CustomerCode string  `gorm:"column:customer_code;size:255;not null"`
	StyleCode    string  `gorm:"column:style_code;size:255;not null"`
	StyleName    string  `gorm:"column:style_name;size:255"`
	SizeType     string  `gorm:"column:size_type;size:255;default:S-M-L;not null"`
	Fabrication  string  `gorm:"column:fabrication;size:255"`
	RmbPrice     float64 `gorm:"column:rmb_price"`
	HsID         int     `gorm:"column:hs_id"`
	Notes        string  `gorm:"column:notes;type:text"`
}

type User struct {
	UserID    int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Email     string `gorm:"column:email;size:255;not null;unique"`
	FirstName string `gorm:"column:first_name;size:255;not null"`
	LastName  string `gorm:"column:last_name;size:255;not null"`
	Password  string `gorm:"column:password;size:64;not null"`
	Role      string `gorm:"column:role;type:enum('admin','ch_user','us_user','user');default:user"`
}

type SearchShipment struct {
	ShipID       int     `json:"ship_id"`
	CustomerCode string  `json:"customer_code"`
	RmbInv       string  `json:"rmb_inv"`
	InvoiceTtl   float64 `json:"invoice_ttl"`
	ShipFrom     string  `json:"ship_from"`
	HouseBlNum   string  `json:"house_bl_num"`
	Exporter     string  `json:"exporter"`
	ShipName     string  `json:"ship_name"`
	ShipDt       string  `json:"ship_dt"`
	Notes        string  `json:"notes"`
	GrossWeight  float64 `json:"gross_weight"`
	ItemCnt      int     `json:"item_cnt"`
	CartonCnt    int     `json:"carton_cnt"`
	CartonSize   float64 `json:"carton_size"`
}

func (s *Shipment) Save(tx *gorm.DB) (int, error) {
	if tx == nil {
		tx = mysqlDb
	}

	// 保存 Shipment
	if err := tx.Table("Shipment").Save(&s).Error; err != nil {
		return -1, err
	}
	return s.ShipID, nil
}

func (s *Shipment) Remove(id string) error {
	return mysqlDb.Table("Shipment").Delete(&Shipment{}, id).Error
}

func (s *Shipment) Search(searchValue, order string, page, size int) ([]res_models.SearchShipment, int64, error) {

	var (
		shipments []res_models.SearchShipment
		total     int64
	)

	query := mysqlDb.Table("Shipment s").
		Select("s.*, SUM(p.gross_weight) AS gross_weight, SUM(p.item_cnt) AS item_cnt, SUM(p.carton_cnt) AS carton_cnt, SUM(p.meas_vol) AS carton_size").
		Joins("LEFT JOIN PackingList p ON s.ship_id = p.ship_id").
		Group("s.ship_id, s.customer_code, s.rmb_inv, s.invoice_ttl, s.ship_from, s.house_bl_num, s.exporter, s.ship_name, s.ship_dt, s.notes")

	if searchValue != "" {
		query = query.Where("s.rmb_inv LIKE ? OR s.customer_code LIKE ? OR s.house_bl_num LIKE ? OR s.exporter LIKE ? OR s.ship_name LIKE ? OR s.notes LIKE ?",
			"%"+searchValue+"%", "%"+searchValue+"%", "%"+searchValue+"%", "%"+searchValue+"%", "%"+searchValue+"%", "%"+searchValue+"%")
	}

	//if order != "" {
	//	query = query.Order(order)
	//} else {
	//	query = query.Order("s.ship_id DESC")
	//}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("ship_dt DESC").Offset((page - 1) * size).Limit(size).Find(&shipments).Error; err != nil {
		return nil, 0, err
	}

	return shipments, total, nil

}

func (s *Shipment) SearchByIdWithInvoice(shipId int) (results res_models.ShipmentAndInvoice, err error) {
	query := `SELECT
				s.*,
				i.* ,
				c.*
			FROM
				Shipment s
			JOIN Invoice i ON s.ship_id = i.ship_id 
			JOIN Customer c ON c.customer_code = s.customer_code
			WHERE
				s.ship_id = ?`
	err = mysqlDb.Raw(query, shipId).Scan(&results).Error
	return
}

func (s *Packing) SearchList(shipmentId, page, size int) ([]res_models.PackingListRes, int, error) {
	offset := (page - 1) * size
	sql := fmt.Sprintf(`SELECT
                        j.ubc_pi,
                        p.*,
                        j.sale_price
						
                    FROM
                        PackingList p
                    JOIN Projection j ON p.proj_id = j.proj_id 
                    WHERE
                        p.ship_id = ? 
                    LIMIT %d OFFSET %d`, size, offset)

	countSql := `SELECT COUNT(*)
             FROM PackingList p
             JOIN Projection j ON p.proj_id = j.proj_id 
             WHERE p.ship_id = ?`
	var totalRecords int
	err := mysqlDb.Raw(countSql, shipmentId).Scan(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}
	var res []res_models.PackingListRes
	if err := mysqlDb.Raw(sql, shipmentId).Scan(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, totalRecords, nil
}

func (p *PackingList) DeleteAndSave(shipId int, ps []PackingList) error {
	return mysqlDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("PackingList").Where("ship_id = ?", shipId).Delete(&PackingList{}).Error; err != nil {
			return err
		}
		if err := tx.Table("PackingList").Save(&ps).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p *PackingList) SaveBatch(ps []PackingList, tx *gorm.DB) error {
	if len(ps) == 0 {
		return nil
	}
	if tx == nil {
		tx = mysqlDb
	}

	return tx.Table("PackingList").Save(&ps).Error
}

func (p *PackingList) SearchByShipId(shipId int) ([]res_models.PackingsRes, error) {
	query := `SELECT
   				 p.list_id as id,
   				 p.list_id as list_id,
   				 p.proj_id as proj_id,
   				 p.sale_price as sale_price,
   				 p.total_quantity as total_quantity,
   				 p.carton_cnt as carton_cnt,
   				 j.fabrication as fabrication,
   				 j.color as color,
   				 j.style_code as style_code,
   				 j.style_name as style_name,
   				 j.customer_po as customer_po
		    FROM
				PackingList p
			JOIN Projection j ON p.proj_id = j.proj_id 
			WHERE
				p.ship_id = ?`

	var results []res_models.PackingsRes
	err := mysqlDb.Raw(query, shipId).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (i *Invoice) SaveOrUpdate(tx *gorm.DB) error {

	if tx == nil {
		tx = mysqlDb
	}
	return tx.Table("Invoice").Save(&i).Error
}

func (i *Invoice) FindById(invoiceId string) error {
	return mysqlDb.Table("Invoice").Where("invoice_id = ?", invoiceId).First(&i).Error
}

func (i *Invoice) Remove(id string) error {
	return mysqlDb.Table("Invoice").Delete(&Shipment{}, id).Error
}

func (i *Invoice) SearchList(searchValue string, page, size int) ([]Invoice, int64, error) {

	var invoices []Invoice
	var totalRecords int64

	// Calculate the offset
	offset := (page - 1) * size

	// Create the query
	query := mysqlDb.Table("Invoice")

	// Get the total count
	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	// Get the paginated results
	if err := query.Offset(offset).Limit(size).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalRecords, nil
}

func (c *Customer) SearchAll() (cus []Customer, err error) {
	err = mysqlDb.Table("Customer").Find(&cus).Error
	return

}

func (c *Customer) SearchList(searchValue string, page, size int) ([]Customer, int64, error) {

	var invoices []Customer
	var totalRecords int64

	offset := (page - 1) * size

	query := mysqlDb.Table("Customer")

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(size).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalRecords, nil
}

func (c *Customer) Remove(id string) error {
	return mysqlDb.Table("Customer").Delete(&Customer{}, id).Error
}

func (c *Customer) SaveOrUpdate() error {

	return mysqlDb.Table("Customer").Save(&c).Error
}

func (p *Projection) SearchAll() (pro []Projection, err error) {
	err = mysqlDb.Table("Projection").Find(&pro).Error
	return

}

func (p *Projection) Remove(id string) error {
	return mysqlDb.Table("Projection").Delete(&Projection{}, id).Error
}

func (p *Projection) SearchList(searchValue string, page, size int) ([]Projection, int64, error) {

	var invoices []Projection
	var totalRecords int64

	offset := (page - 1) * size

	query := mysqlDb.Table("Projection")

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(size).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalRecords, nil
}

func (p *Projection) SaveOrUpdate() error {
	return mysqlDb.Table("Projection").Save(&p).Error
}

func (p *Projection) SaveAll(ps []Projection) error {
	return mysqlDb.Table("Projection").Save(&ps).Error
}

func SaveShipmentAndPackingAndInvoice(shipment *Shipment, list []PackingList, invoice *Invoice) (shipmentId, invoiceId int, err error) {

	err = mysqlDb.Transaction(func(tx *gorm.DB) error {
		sId, errs := shipment.Save(tx)
		if errs != nil {
			return errs
		}
		pl := PackingList{}
		for i := range list {
			list[i].ShipID = sId
		}
		if errs = pl.SaveBatch(list, tx); errs != nil {
			return errs
		}

		invoice.ShipID = sId
		if errs = invoice.SaveOrUpdate(tx); errs != nil {
			return errs
		}
		return nil
	})

	return shipment.ShipID, invoice.InvoiceID, err
}
