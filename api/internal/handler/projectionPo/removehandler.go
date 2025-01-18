package projectionPo

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/projectionPo"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveProjection
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := projectionPo.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove(&req)
		result.HttpResult(r, w, nil, err)
	}
}