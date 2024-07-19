package shipment

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"

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

func (l *SaveOrUpdateLogic) SaveOrUpdate(req *types.SaveOrUpdateShipment) error {

	shipment := &models.Shipment{
		ShipID:       req.ID,
		RmbInv:       req.RMBInv,
		InvoiceTtl:   req.InvoiceTTL,
		CustomerCode: req.CustomerCode,
		MasterPo:     req.MasterPO,
		ShipFrom:     req.ShipFrom,
		UbcPi:        req.UBCPI,
		OrigCountry:  req.OrigCountry,
		ShipMethod:   req.ShipMethod,
		ShipTerm:     req.ShipTerm,
		MasterBlNum:  req.MasterBLNum,
		HouseBlNum:   req.HouseBLNum,
		Exporter:     req.Exporter,
		ShipName:     req.ShipName,
		ShipDt:       req.ShipDT,
		ArriveDt:     req.ArriveDT,
		Notes:        req.Notes,
	}

	_, err := shipment.Save(nil)
	if err != nil {
		l.Errorf("[SaveOrUpdate] err:%v", err)
		return xerr.ServerCommonError
	}
	return nil
}
