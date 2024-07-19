package shipment

import (
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePackingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSavePackingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePackingLogic {
	return &SavePackingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SavePackingLogic) SavePacking(req *types.SaveShipment) error {
	// todo: add your logic here and delete this line

	return nil
}
