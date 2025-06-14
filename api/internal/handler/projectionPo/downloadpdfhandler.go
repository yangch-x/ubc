package projectionPo

import (
	"net/http"

	"UBC/api/internal/logic/projectionPo"
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadPdfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := projectionPo.NewDownloadPdfLogic(r.Context(), svcCtx)
		_ = l.DownloadPdf(&req, w)

	}
}
