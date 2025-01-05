package projection

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchRemoveLogic {
	return &BatchRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchRemoveLogic) BatchRemove(req *types.IdsReq) error {
	s := models.Projection{}

	err := s.BatchRemove(req.Ids)
	if err != nil {
		l.Errorf("[BatchRemove] err:%", err)
		return xerr.DeleteChatInfoErr
	}
	return nil
}
