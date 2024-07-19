// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	common "UBC/api/internal/handler/common"
	shipment "UBC/api/internal/handler/shipment"
	usernoauth "UBC/api/internal/handler/usernoauth"
	"UBC/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: usernoauth.SigninHandler(serverCtx),
			},
		},
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/shipment/saveOrUpdate",
				Handler: shipment.SaveOrUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/shipment/remove",
				Handler: shipment.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shipment/search",
				Handler: shipment.SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/packing/list",
				Handler: shipment.PackingListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shipment/saveShipmentAndIVoice",
				Handler: shipment.SaveShipmentAndVoiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shipment/savePacking",
				Handler: shipment.SavePackingHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/customer/searchAllCustomerAndProjection",
				Handler: shipment.SearchAllCustomerAndProjectionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/common/upload",
				Handler: common.UploadFileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)
}
