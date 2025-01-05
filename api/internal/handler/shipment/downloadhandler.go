package shipment

import (
	"UBC/api/library/result"
	"net/http"

	"UBC/api/internal/logic/shipment"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := shipment.NewDownloadLogic(r.Context(), svcCtx)
		file, err := l.Download(&req)
		if err != nil {
			result.HttpResult(r, w, nil, err)
			return
		}

		// 设置响应头
		w.Header().Set("Content-Disposition", "attachment; filename=发票数据.xlsx")
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(file)
	}
}
