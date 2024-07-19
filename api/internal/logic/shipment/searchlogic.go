package shipment

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

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

func (l *SearchLogic) Search(req *types.QueryShipment) (resp *types.ListRes, err error) {
	s := models.Shipment{}
	search, i, err := s.Search(req.SearchParams, req.Order, req.PageNo, req.PageSize)
	if err != nil {
		l.Error("Search err:%v", err)
		return nil, xerr.DbError
	}
	resp = &types.ListRes{Res: search, Total: int(i), PageNo: req.PageNo, PageSize: req.PageSize}
	return
}
