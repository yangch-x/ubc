package shipment

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/shipment"
	"UBC/api/internal/svc"
)

func SearchAllCustomerAndProjectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shipment.NewSearchAllCustomerAndProjectionLogic(r.Context(), svcCtx)
		resp, err := l.SearchAllCustomerAndProjection()
		result.HttpResult(r, w, resp, err)
	}
}
