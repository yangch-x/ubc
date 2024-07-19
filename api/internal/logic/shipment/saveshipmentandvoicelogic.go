package shipment

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"
	"time"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveShipmentAndVoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveShipmentAndVoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveShipmentAndVoiceLogic {
	return &SaveShipmentAndVoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveShipmentAndVoiceLogic) SaveShipmentAndVoice(req *types.ShipmentAndInvoice) (resp *types.ShipmentAndInvoiceRes, err error) {
	resp = &types.ShipmentAndInvoiceRes{}
	shipment := &models.Shipment{
		ShipID:       req.ShipmentId,
		HouseBlNum:   req.BillOfLanding,
		CustomerCode: req.CustomerCode,
		ShipFrom:     req.ShipFrom,
		ShipMethod:   req.ShipMethod,
		OrigCountry:  req.CountryOfOrigin,
		Exporter:     req.Manufacture,
		ShipName:     req.VesselFlight,
		ShipDt:       time.Unix(req.ETDDt/1000, 0).Format("2006-01-02 15:04:05"),
		UbcPi:        req.UBCPI,
		Markurl:      req.MarksAndNumbers,
	}

	var plireq []models.PackingList
	invoice := &models.Invoice{
		InvoiceID:    req.InvoiceId,
		UbcPi:        req.UBCPI,
		InvoiceCode:  req.InvoiceCode,
		CustomerCode: req.CustomerCode,
		InvoiceAmt:   req.AdditionalCost,
		ReceivedAmt:  req.DepositAmt,
		InvoiceDt:    time.Unix(req.InvoiceDt/1000, 0).Format("2006-01-02 15:04:05"),
		InvoiceDue:   time.Unix(req.InvoiceDue/1000, 0).Format("2006-01-02 15:04:05"),
	}
	sId, iId, err := models.SaveShipmentAndPackingAndInvoice(shipment, plireq, invoice)
	if err != nil {
		l.Error("SaveShipmentAndVoice err:%v", err)
		return resp, xerr.DbError
	}
	resp.ShipmentId = sId
	resp.InvoiceId = iId
	return
}
