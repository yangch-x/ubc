package result

import (
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, utils.JsonCamelCase{Value: r})
	} else {
		// 错误返回
		errcode := xerr.ServerCommonError
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)          // err类型
		if e, ok := causeErr.(xerr.Code); ok { // 自定义错误类型
			// 自定义CodeError
			errcode = e
			errmsg = e.Message()

		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { // 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = xerr.Int(grpcCode)
					errmsg = gstatus.Message()
				}
			} else {
				errmsg = err.Error()
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusOK, Error(errcode.Code(), errmsg))
	}
}

// AuthHttpResult 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errcode := xerr.ServerCommonError
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)          // err类型
		if e, ok := causeErr.(xerr.Code); ok { // 自定义错误类型
			// 自定义CodeError
			errcode = e
			errmsg = e.Message()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { // 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = xerr.Int(grpcCode)
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode.Code(), errmsg))
	}
}

// ParamErrorResult  http 参数错误返回
func ParamErrorResult(_ *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.RequestParamError.Message(), err.Error())
	httpx.WriteJson(w, http.StatusOK, Error(xerr.RequestParamError.Code(), errMsg))
}

// HttpCacheOK writes json string load from redis into w.
func HttpCacheOK(w http.ResponseWriter, resp []byte) {
	w.Header().Set(httpx.ContentType, httpx.JsonContentType)
	w.WriteHeader(200)

	if n, err := w.Write(resp); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(resp) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(resp), n)
	}
}
