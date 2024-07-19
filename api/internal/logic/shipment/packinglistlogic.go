package shipment

import (
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PackingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPackingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PackingListLogic {
	return &PackingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PackingListLogic) PackingList(req *types.PackingListReq) (resp *types.ListRes, err error) {
	s := models.Packing{}
	search, count, err := s.SearchList(req.ShipmentId, req.PageNo, req.PageSize)
	if err != nil {
		return nil, err
	}

	resp = &types.ListRes{Res: search, Total: count, PageNo: req.PageNo, PageSize: req.PageSize}
	return
}
