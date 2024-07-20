package projection

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/projection"
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

		l := projection.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove(&req)
		result.HttpResult(r, w, nil, err)
	}
}
