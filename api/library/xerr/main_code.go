package xerr

var (
	// 基础
	OK                        = add(200)    // "SUCCESS" //成功返回
	ServerCommonError         = add(100001) // "服务器开小差啦,稍后再来试一试"
	RequestParamError         = add(100002) // "参数错误"
	TokenExpireError          = add(100003) // "token失效 //请重新登陆"
	DbError                   = add(100004) // "数据库繁忙 //请稍后再试"
	ErrGenerateTokenError     = add(100005) // "生成token失败"
	Unauthorized              = add(100006) // "签名验证错误" // 未认证
	ServerErr                 = add(100008) // 服务器错误
	LimitExceed               = add(100009) // 超出限制
	AccessDenied              = add(100010) // 访问权限不足
	RequestIllegal            = add(100011) // 非法请求
	RequestThirdErr           = add(100012) // 请求第三方失败
	SpeechToTextErr           = add(100013) // 语音解析失败
	GptChatNotFoundErr        = add(100014) // chat不存在
	GptChatMappingNotFoundErr = add(100015) // mapping不存在
	DeleteChatInfoErr         = add(100016) // 删除错误
	UpdateChatTitleErr        = add(100017) // 修改标题错误
	NotFoundProductHeaderErr  = add(100018) // 产品key不存在

	//用户模块
	ErrUsernamePwd         = add(200001) // "账号或密码不正确"
	ErrUserNoExists        = add(200002) // "用户不存在"
	ErrUserExists          = add(200003) // "该手机号或邮箱已被注册"
	ErrAccountFormat       = add(200004) // "账号格式错误"
	ErrAccountForbidden    = add(200005) // "账号被禁用"
	ErrMobileFormat        = add(200006) // "手机格式错误"
	ErrAccountExists       = add(200007) // "账号已存在"
	ErrEmailFormat         = add(200008) // "邮箱格式错误"
	ErrSignup              = add(200009) // "注册失败"
	ErrGoogleSignIn        = add(200010) // "google登录失败"
	ErrGoogleEmailNotFound = add(200011) // "google用户信息获取失败"
	UpdateUserInfoErr      = add(200012) // "修改用户信息失败"
	ResetPwdUrlErr         = add(200013) // "重置密码链接错误"
	ResetPwdUrlTimeOutErr  = add(200014) // "链接失效，请重新发送邮箱"
	ResetPwdErr            = add(200015) // "重置密码错误"
	OldPwdErr              = add(200016) // "旧密码错误"
	FreeTokenInsufficient  = add(200017) // "免费token余额不足"
	FreeDocInsufficient    = add(200018) // "免费文档数余额不足"
	TokenInsufficient      = add(200019) // "token余额不足"
	ModelUnavailable       = add(200020) // "高级模式功能不可用"
	ModuleUnavailable      = add(200021) // "高级模板功能不可用"
	ErrUpdateAvatar        = add(200022) // "头像修改失败"
	ErrGetAvatar           = add(200023) // "头像获取失败"
	CreditInsufficient     = add(200024) // "Credit额度不足"
	ErrAppleAuthorization  = add(200025) // "Apple授权失败"
	ErrDeleteUser          = add(200026) // "删除账号失败"
	ErrWechatSignIn        = add(200027) // "微信登录失败"
	ErrBindMobile          = add(200028) // "绑定手机号失败"
	ErrCheckIn             = add(200029) // "签到失败"
	ErrGetCheckIn          = add(200030) // "获取签到信息失败"

	//验证码模块
	ErrCptTplCodeNoFound   = add(300001) // "未找到该验证码模板"
	ErrCptNoMatch          = add(300002) // "验证码不正确,请重新输入"
	ErrCptTooFrequent24    = add(300003) // "该手机获取验证码过于频繁，请24小时后重试"
	ErrCptTooFrequent48    = add(300004) // "该手机获取验证码过于频繁，请48小时后重试"
	ErrCptFrequent         = add(300005) // "操作过于频繁，请稍后重试"
	ErrCptExpire           = add(300006) // "验证码失效，请重新获取"
	ErrSendEmail           = add(300007) // "验证码发送失败"
	ErrSendEmailOutOfLimit = add(300008) // "每小时发送次数超出上限"

	// doc 模块
	ErrSaveDoc           = add(400001) // "新增文档错误"
	ErrDeleteDoc         = add(400002) // "删除文档错误"
	ErrGetDoc            = add(400003) // "查询文档错误"
	ErrDeleteCard        = add(400004) // "删除卡片错误"
	ErrSaveCard          = add(400005) // "新增卡片错误"
	ErrUpdateCard        = add(400006) // "修改卡片错误"
	ErrNotFoundDoc       = add(400007) // "文档不存在"
	ErrUpdateCommentCard = add(400008) // "修改卡片批注错误"
	ErrInsertCommentCard = add(400009) // "新增卡片批注错误"
	ErrDeleteCommentCard = add(400010) // "删除卡片批注错误"
	ErrCommentMaxNum     = add(400011) // "批注数量超过最大限"
	ErrTemplate          = add(400012) // "获取模板错误"
	ErrAddFigurePen      = add(400013) // "新增数字笔错误"
	ErrUpdateFigurePen   = add(400014) // "修改数字笔错误"
	ErrNotFoundFigurePen = add(400015) // "数字笔不存在"
	ErrFigurePenNum      = add(400016) // "数字笔数量超过限制"
	ErrFigurePenExist    = add(400017) // "数字笔已创建，请勿重复创建"

	// gpt模块
	ErrMethodNotFound = add(500001) // "方法未找到"
	ErrParam          = add(500002) // "参数错误"

	// 支付模块
	UserSubscription = add(600001) // "用户已订阅"
	ErrBuyAppVerify  = add(600002) // "订阅失败"
	// wap-view
	ErrSaveUserInfo           = add(700001) // "设置用户信息失败"
	ErrStartInterview         = add(700002) // "开始面试失败"
	ErrSaveMapping            = add(700003) // "报错会话信息失败"
	ErrNotFoundChat           = add(700004) // "记录不存在"
	ErrGenerateFeedback       = add(700005) // "面试反馈生成失败"
	ErrNotFoundRandomQuestion = add(700006) // "该场景没有配置随机问题"
	ErrNotFoundIntroduction   = add(700007) // "该场景没有配置介绍问题"
	ErrNotFoundPrompt         = add(700008) // "该场景没有配置提示词"
)
