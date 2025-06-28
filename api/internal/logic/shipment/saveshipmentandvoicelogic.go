package shipment

import (
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"UBC/models"
	"context"

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
		ShipFrom:     req.ShipFrom,
		Exporter:     req.Manufacture,
		OrigCountry:  req.CountryOfOrigin,
		ShipName:     req.VesselFlight,
		UbcPi:        req.UBCPI,
		ShipDt:       utils.FormatDateToYMD(req.ETDDt),
		CustomerCode: req.CustomerCode,
		ShipMethod:   req.ShipMethod,
	}

	var plireq []models.PackingList
	invoice := &models.Invoice{
		InvoiceID:                 req.InvoiceId,
		UbcPi:                     req.UBCPI,
		InvoiceCode:               req.InvoiceCode,
		CustomerCode:              req.CustomerCode,
		InvoiceAmt:                req.AdditionalCost,
		ReceivedAmt:               req.ReceivedAmt,
		AdditionalCost:            req.AdditionalCost,
		AdditionalCostDescription: req.AdditionalCostDescription,
		InvoiceDt:                 utils.FormatDateToYMD(req.InvoiceDt),
		InvoiceDue:                utils.FormatDateToYMD(req.InvoiceDue),
	}
	sId, iId, err := models.SaveShipmentAndPackingAndInvoice(shipment, plireq, invoice)
	if err != nil {
		l.Error("SaveShipmentAndVoice err:%v", err)
		return resp, xerr.SaveShipmentError
	}
	resp.ShipmentId = sId
	resp.InvoiceId = iId
	return
}
