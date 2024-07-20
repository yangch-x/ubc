package projection

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/projection"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveOrUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveProection
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := projection.NewSaveOrUpdateLogic(r.Context(), svcCtx)
		err := l.SaveOrUpdate(&req)
		result.HttpResult(r, w, nil, err)

	}
}
