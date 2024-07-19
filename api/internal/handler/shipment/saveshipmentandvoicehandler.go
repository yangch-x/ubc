package shipment

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/shipment"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveShipmentAndVoiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShipmentAndInvoice
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := shipment.NewSaveShipmentAndVoiceLogic(r.Context(), svcCtx)
		resp, err := l.SaveShipmentAndVoice(&req)
		result.HttpResult(r, w, resp, err)
	}
}
