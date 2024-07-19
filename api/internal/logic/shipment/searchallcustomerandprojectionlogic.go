package shipment

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAllCustomerAndProjectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchAllCustomerAndProjectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllCustomerAndProjectionLogic {
	return &SearchAllCustomerAndProjectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchAllCustomerAndProjectionLogic) SearchAllCustomerAndProjection() (resp *types.SearchAllCustomerAndPrijection, err error) {
	c := models.Customer{}
	cus, err := c.SearchAll()
	if err != nil {
		l.Error("[SearchAllCustomerAndProjection] search customer err:%v", err)
		return nil, xerr.DbError
	}
	p := models.Projection{}
	pro, err := p.SearchAll()
	if err != nil {
		l.Error("[SearchAllCustomerAndProjection] search projection err:%v", err)
		return nil, xerr.DbError
	}
	return &types.SearchAllCustomerAndPrijection{Customers: cus, Projections: pro}, nil
}
