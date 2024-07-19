package usernoauth

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/usernoauth"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SigninHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := usernoauth.NewSigninLogic(r.Context(), svcCtx)
		resp, err := l.Signin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
