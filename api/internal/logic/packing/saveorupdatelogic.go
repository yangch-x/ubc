package packing

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

func (l *SaveOrUpdateLogic) SaveOrUpdate(req *types.CreatePackings) error {
	p := models.PackingList{}
	var plireq []models.PackingList
	for _, packing := range req.CreatePackings {
		plireq = append(plireq, models.PackingList{
			ShipID:        req.ShipId,
			ListID:        packing.ListId,
			ProjID:        packing.ProjId,
			CartonCnt:     packing.CartonCnt,
			SalePrice:     packing.SalePrice,
			TotalQuantity: packing.TotalQuantity,
		})
	}
	err := p.DeleteAndSave(req.ShipId, plireq)
	if err != nil {
		l.Errorf("[SaveOrUpdate] save batch err:%v", err)
		return xerr.DbError
	}
	return nil
}
