package packing

import (
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"UBC/api/library/xerr"
	"UBC/models"
	"UBC/models/res_models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.QueryPacking) (resp *types.QueryPackingResp, err error) {
	p := models.PackingList{}

	packings, err := p.SearchByShipId(req.ShipId)
	if err != nil {
		l.Errorf("search packing list by ship id err:%v", err)
		return nil, xerr.DbError
	}

	s := models.Shipment{}
	sv, err := s.SearchByIdWithInvoice(req.ShipId)
	if err != nil {
		l.Errorf("search shipment and invoice by ship id err:%v", err)
		return nil, xerr.DbError
	}
	sRes := res_models.ShipAndInvoice{
		ShipID:                    sv.ShipID,
		InvoiceId:                 sv.InvoiceID,
		MasterPO:                  sv.MasterPO,
		CustomerCode:              sv.CustomerCode,
		UbcPi:                     sv.UbcPi,
		OrigCountry:               sv.OrigCountry,
		ShipMethod:                sv.ShipMethod,
		ShipTerm:                  sv.ShipTerm,
		InvoiceTtl:                sv.InvoiceTtl,
		ShipFrom:                  sv.ShipFrom,
		HouseBlNum:                sv.HouseBLNum,
		Manufacture:               sv.Exporter,
		ShipName:                  sv.ShipName,
		InvoiceCode:               sv.InvoiceCode,
		InvoiceAmt:                sv.InvoiceAmt,
		ReceivedAmt:               sv.ReceivedAmt,
		ShipDt:                    sv.ShipDt,
		InvoiceDt:                 sv.InvoiceDt,
		InvoiceDue:                sv.InvoiceDue,
		ShipTo:                    sv.ShipTo,
		BillingContact:            sv.BillingContact,
		Term:                      sv.Term,
		AdditionalCost:            sv.AdditionalCost,
		AdditionalCostDescription: sv.AdditionalCostDescription,
	}
	//
	//sRes.ShipDt, _ = time.Parse("2006-01-02", sv.ShipDt)
	//sRes.InvoiceDt, _ = time.Parse("2006-01-02", sv.InvoiceDt)
	//sRes.InvoiceDue, _ = time.Parse("2006-01-02", sv.InvoiceDue)

	return &types.QueryPackingResp{Packings: packings, Shipment: sRes}, nil
}
