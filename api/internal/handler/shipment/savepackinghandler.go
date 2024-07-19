package shipment

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/shipment"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SavePackingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveShipment
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := shipment.NewSavePackingLogic(r.Context(), svcCtx)
		err := l.SavePacking(&req)
		result.HttpResult(r, w, nil, err)
	}
}
