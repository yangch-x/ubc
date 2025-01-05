package shipment

import (
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"UBC/models"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInoviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInoviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInoviceLogic {
	return &CreateInoviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInoviceLogic) CreateInovice(req *types.CreateInvoiceReq, w http.ResponseWriter) (res *types.CreateInvoiceRes, err error) {

	invoiceDate, _ := utils.ConvertTimeFormat(req.Shipment.InvoiceDt)
	dueData, _ := utils.ConvertTimeFormat(req.Shipment.InvoiceDue)
	invoiceOne := []string{req.Shipment.InvoiceCode, invoiceDate, dueData}
	var billTo, shipTo []string
	// 修改 sub total ，Total PCs

	invoice := &models.Invoice{InvoiceCode: req.Invoice.InvoiceCode, TotalPCs: req.Invoice.TotalPCs, SubTotal: req.Invoice.SubTotal}
	if err = invoice.UpdateByInvoiceCode(); err != nil {
		l.Errorf("[CreateInvoice] update totalPsc and sub total err:%v", err)
		return nil, xerr.CreateInvoiceErr
	}

	if len(req.Shipment.BillingContact) != 0 {
		billTo = strings.Split(req.Shipment.BillingContact, "|")
	}

	if len(req.Shipment.ShipTo) != 0 {
		shipTo = strings.Split(req.Shipment.ShipTo, "|")
	}

	table1Data := []utils.Table1Row{
		{req.Shipment.ShipFrom, "", req.Shipment.UbcPi, req.Shipment.ShipMethod, req.Shipment.Term},
	}
	req.Shipment.EtdDt, _ = utils.ConvertToDateOnly(req.Shipment.EtdDt)

	table2Data := []utils.Table2Row{
		{req.Shipment.CountryOfOrigin, req.Shipment.VesselFlight, req.Shipment.BillOfLanding, req.Shipment.EtdDt},
	}
	table3Data := make([]utils.Table3Row, len(req.Packings)+1)
	total := 0
	for i := range req.Packings {
		table3Data[i] = utils.Table3Row{
			PO:          req.Packings[i].CustomerPo,
			StyleName:   req.Packings[i].StyleName,
			StyleCode:   req.Packings[i].StyleCode,
			Description: req.Packings[i].Fabrication,
			Color:       req.Packings[i].Color,
			Size:        req.Packings[i].Size,
			Qty:         fmt.Sprintf("%d", req.Packings[i].TotalQuantity),
			UPrice:      fmt.Sprintf("%.2f", req.Packings[i].SalePrice),
			TotalUSD:    fmt.Sprintf("%.2f", float64(req.Packings[i].TotalQuantity)*req.Packings[i].SalePrice),
		}
		total += req.Packings[i].TotalQuantity
	}
	// 新增空行
	for i := 0; i < 25-len(table3Data); i++ {
		t := utils.Table3Row{UPrice: "", TotalUSD: "-"}
		table3Data = append(table3Data, t)
	}
	// 计算Qty
	var totalQty int64
	for _, t := range table3Data {
		q, _ := strconv.ParseInt(t.Qty, 10, 32)
		totalQty += q
	}

	subStr := fmt.Sprintf("%.2f", req.Invoice.SubTotal)

	// 小计
	table3Data = append(table3Data, utils.Table3Row{Description: "TOTAL AMOUNT", Qty: strconv.FormatInt(totalQty, 10), TotalUSD: subStr})

	// Deposit
	table3Data = append(table3Data, utils.Table3Row{
		Description: "Deposit",
		TotalUSD:    fmt.Sprintf("%.2f", req.Shipment.DepositAmt),
	})

	// 判断是否有附加费用
	if req.Shipment.AdditionalCost > 0 && len(req.Shipment.AdditionalCostDescription) != 0 {
		table3Data = append(table3Data, utils.Table3Row{
			Description: req.Shipment.AdditionalCostDescription,
			TotalUSD:    fmt.Sprintf("%.2f", req.Shipment.AdditionalCost),
		})
	}

	for i := 0; i < 2; i++ {
		t := utils.Table3Row{UPrice: "", TotalUSD: "-"}
		table3Data = append(table3Data, t)
	}

	totals := req.Invoice.SubTotal + req.Shipment.AdditionalCost + req.Shipment.DepositAmt

	subStr = fmt.Sprintf("%.2f", req.Invoice.SubTotal+req.Shipment.AdditionalCost+req.Shipment.DepositAmt)

	totalStr := fmt.Sprintf("%d", total)
	cnStr := utils.ConvertFloatToWords(totals)
	lastStr := fmt.Sprintf("TOTAL %d CTNS\n%s", req.Invoice.TotalCartons, cnStr)
	pdfBuffer, err := utils.BuildInvoicePdf(table1Data, table2Data, table3Data, l.svcCtx.Config.Address, l.svcCtx.Config.Invoice,
		invoiceOne, billTo, shipTo, lastStr, totalStr, subStr)
	if err != nil {
		l.Errorf("[CreateInvoice] build pdf err:%v", err)
		return nil, xerr.CreateInvoiceErr
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf")
	_, err = w.Write(pdfBuffer.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return nil, nil
}
