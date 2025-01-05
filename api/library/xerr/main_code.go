package xerr

var (
	// 基础
	OK                    = add(200)    // "SUCCESS" //成功返回
	ServerCommonError     = add(100001) // "服务器开小差啦,稍后再来试一试"
	RequestParamError     = add(100002) // "参数错误"
	TokenExpireError      = add(100003) // "token失效 //请重新登陆"
	DbError               = add(100004) // "数据库繁忙 //请稍后再试"
	ErrGenerateTokenError = add(100005) // "生成token失败"
	Unauthorized          = add(100006) // "签名验证错误" // 未认证
	SaveShipmentError     = add(100007) // 保存shipment失败
	ServerErr             = add(100008) // 服务器错误
	LimitExceed           = add(100009) // 超出限制
	AccessDenied          = add(100010) // 访问权限不足
	RequestIllegal        = add(100011) // 非法请求
	RequestThirdErr       = add(100012) // 请求第三方失败
	DeleteChatInfoErr     = add(100016) // 删除错误
	CreateInvoiceErr      = add(100017) // 创建Invoice失败
	EmptyErr              = add(100018)
)
