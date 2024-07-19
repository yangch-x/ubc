package common

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/common"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFile
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := common.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req, r)
		result.HttpResult(r, w, resp, err)
	}
}
