package shipment

import (
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"context"
	"fmt"
	"net/http"
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

	if len(req.Shipment.BillingContact) != 0 {
		billTo = strings.Split(req.Shipment.BillingContact, "|")
	}

	if len(req.Shipment.ShipTo) != 0 {
		shipTo = strings.Split(req.Shipment.ShipTo, "|")
	}

	table1Data := []utils.Table1Row{
		{req.Shipment.ShipFrom, "", req.Shipment.UbcPi, req.Shipment.VesselFlight, req.Shipment.Term},
	}
	req.Shipment.EtdDt, _ = utils.ConvertToDateOnly(req.Shipment.EtdDt)

	table2Data := []utils.Table2Row{
		{req.Shipment.CountryOfOrigin, req.Shipment.VesselFlight, req.Shipment.BillOfLanding, req.Shipment.EtdDt},
	}
	table3Data := make([]utils.Table3Row, len(req.Packings))
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
	totalStr := fmt.Sprintf("%d", total)
	subStr := fmt.Sprintf("%.2f", req.Invoice.SubTotal)
	cnStr := utils.ConvertFloatToWords(req.Invoice.SubTotal)
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
