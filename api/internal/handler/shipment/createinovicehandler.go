package shipment

import (
	"net/http"

	"UBC/api/internal/logic/shipment"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateInoviceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateInvoiceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := shipment.NewCreateInoviceLogic(r.Context(), svcCtx)
		_, _ = l.CreateInovice(&req, w)
		//result.HttpResult(r, w, resp, err)
	}
}
