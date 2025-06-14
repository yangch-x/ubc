package projectionPo

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"
	"gorm.io/datatypes"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateLogic {
	return &SaveOrUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateLogic) SaveOrUpdate(req *types.SaveProection) error {
	projection := &models.ProjectionPo{
		Id:                  req.Id,
		ArriveDt:            req.ArriveDt,
		PoDate:              req.PoDate,
		UbcPi:               req.UbcPi,
		FobLdp:              req.FobLdp,
		CustomerCode:        req.CustomerCode,
		Country:             req.Country,
		CustomerPo:          req.CustomerPo,
		MasterPo:            req.MasterPo,
		StyleCode:           req.StyleCode,
		StyleName:           req.StyleName,
		Fabrication:         req.Fabrication,
		Color:               req.Color,
		Size:                req.Size,
		PoQty:               req.PoQty,
		ShipQty:             req.ShipQty,
		SalePrice:           req.SalePrice,
		SaleCustPrice:       req.SaleCustPrice,
		SaleCurrency:        req.SaleCurrency,
		InvoiceCode:         req.InvoiceCode,
		Receiving:           req.Receiving,
		Notes:               req.Notes,
		CostPrice:           req.CostPrice,
		CostCurrency:        req.CostCurrency,
		RmbInv:              req.RmbInv,
		Exporter:            req.Exporter,
		UbcPayable:          req.UbcPayable,
		PayPeriod:           req.PayPeriod,
		SalesPerson:         req.SalesPerson,
		SalesCommission:     req.SalesCommission,
		CommPaid:            req.CommPaid,
		TtlSell:             req.TtlSell,
		TtlBuy:              req.TtlBuy,
		PoItems:             datatypes.JSON(req.PoItems),
		ShipTo:              req.ShipTo,
		ShipFrom:            req.ShipFrom,
		ShipTerms:           req.ShipTerms,
		PaymentTerms:        req.PaymentTerms,
		LastRevised:         req.LastRevised,
		PoTotal:             req.PoTotal,
		PageInfo:            req.PageInfo,
		ShipVia:             req.ShipVia,
		SpecialInstructions: req.SpecialInstructions,
	}

	err := projection.SaveOrUpdate()
	if err != nil {
		l.Errorf("[SaveOrUpdate] err:%v", err)
		return xerr.ServerCommonError
	}
	return nil
}
