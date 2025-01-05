package xerr

var (
	messageZh = map[uint32]string{
		OK.Code():                    "SUCCESS",
		ServerCommonError.Code():     "服务器开小差啦,稍后再来试一试",
		RequestParamError.Code():     "参数错误",
		TokenExpireError.Code():      "token失效,请重新登陆",
		DbError.Code():               "数据库繁忙,请稍后再试",
		ErrGenerateTokenError.Code(): "生成token失败",
		Unauthorized.Code():          "签名验证错误",
		SaveShipmentError.Code():     "保存shipment失败",
		DeleteChatInfoErr.Code():     "删除对话框错误",
		CreateInvoiceErr.Code():      "创建发票失败",
		EmptyErr.Code():              "未查询到数据",
	}

	messageEn = map[uint32]string{
		OK.Code():                    "SUCCESS",
		ServerCommonError.Code():     "Server encountered a glitch, please try again later",
		RequestParamError.Code():     "Parameter error",
		TokenExpireError.Code():      "Token expired, please log in again",
		DbError.Code():               "Database busy, please try again later",
		ErrGenerateTokenError.Code(): "Failed to generate token",
		Unauthorized.Code():          "Signature validation error",
		SaveShipmentError.Code():     "new shipment failed",
		ServerErr.Code():             "Server error",
		LimitExceed.Code():           "Limit exceeded",
		AccessDenied.Code():          "Insufficient access permissions",
		RequestIllegal.Code():        "Illegal request",
		RequestThirdErr.Code():       "Failed to request third-party",
		DeleteChatInfoErr.Code():     "Error deleting dialogue box",
		CreateInvoiceErr.Code():      "Error Create Invoice",
		EmptyErr.Code():              "No data is queried",
	}
)

func Init(lang string) {
	switch lang {
	case "zh":
		Register(messageZh)
	case "en":
		Register(messageEn)
	default:
		panic("unknown lang type")
	}

}
