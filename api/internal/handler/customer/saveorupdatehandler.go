package customer

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/customer"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveOrUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveOrUpdateCustomer
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := customer.NewSaveOrUpdateLogic(r.Context(), svcCtx)
		err := l.SaveOrUpdate(&req)
		result.HttpResult(r, w, nil, err)

	}
}
