package usernoauth

import (
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigninLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigninLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigninLogic {
	return &SigninLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SigninLogic) Signin(req *types.SignInRequest) (resp *types.SignInResponse, err error) {
	resp = &types.SignInResponse{}
	token, expire, refreshExpire, err := utils.GetJwtToken("414685046@qq.com", "aaasss", l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		l.Errorf("[SignIn] GetJwtToken err:%v\n", err)
		return resp, xerr.ServerCommonError
	}
	resp.Token = token
	resp.AccessExpire = expire
	resp.RefreshAfter = refreshExpire
	return
}
