package projectionPo

import (
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.IdsReq) ([]byte, error) {
	s := models.ProjectionPo{}
	search, err := s.SearchByIds(req.Ids)
	if err != nil {
		l.Error("Search err:%v", err)
		return nil, nil
	}
	if search == nil {
		l.Error("未查询到数据")
		return nil, xerr.EmptyErr
	}
	export, err := utils.Export(search, "Po", "Po")
	if err != nil {
		l.Errorf("导出文件失败：%s", err)
		return nil, xerr.ServerCommonError
	}

	return export, nil
}
