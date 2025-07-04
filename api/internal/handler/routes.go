// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	common "UBC/api/internal/handler/common"
	customer "UBC/api/internal/handler/customer"
	packing "UBC/api/internal/handler/packing"
	projection "UBC/api/internal/handler/projection"
	projectionPo "UBC/api/internal/handler/projectionPo"
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
				Path:    "/shipment/download",
				Handler: shipment.DownloadHandler(serverCtx),
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
				Method:  http.MethodGet,
				Path:    "/customer/searchAllCustomerAndProjection",
				Handler: shipment.SearchAllCustomerAndProjectionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shipment/createInvoice",
				Handler: shipment.CreateInoviceHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/projection/saveOrUpdate",
				Handler: projection.SaveOrUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/projection/remove",
				Handler: projection.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projection/batch_remove",
				Handler: projection.BatchRemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projection/search",
				Handler: projection.SearchHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/saveOrUpdate",
				Handler: projectionPo.SaveOrUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/updateFields",
				Handler: projectionPo.UpdateFieldsHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/projectionPo/remove",
				Handler: projectionPo.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/batch_remove",
				Handler: projectionPo.BatchRemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/search",
				Handler: projectionPo.SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/download",
				Handler: projectionPo.DownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/projectionPo/download/pdf",
				Handler: projectionPo.DownloadPdfHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/customer/saveOrUpdate",
				Handler: customer.SaveOrUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/customer/remove",
				Handler: customer.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/search",
				Handler: customer.SearchHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ubc/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/packing/saveOrUpdate",
				Handler: packing.SaveOrUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/packing/search",
				Handler: packing.SearchHandler(serverCtx),
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
