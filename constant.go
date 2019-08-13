package gaodeng

const (
	EnvProd = "prod" // (生产环境)
	EnvTest = "test" // (沙盒环境)
)

const (
	StatusCodeNormal               uint16 = 0    // 正常
	StatusCodeLackCommonParameters uint16 = 1001 // 缺少公共参数
	StatusCodeAppKeyNotExists      uint16 = 1002 // appkey不存在
	StatusCodeSignError            uint16 = 1003 // 签名错误
	StatusCodeParameterError       uint16 = 1004 // 业务逻辑校验参数错误
	StatusCodePlatformOuterError   uint16 = 1005 // 开放平台异常(外部依赖)
	StatusCodePlatformInnerError   uint16 = 1006 // 开放平台异常(内部)
	StatusCodeOrderNotExists       uint16 = 2001 // 订单不存在
	StatusCodeOrderCanceled        uint16 = 2002 // 订单取消
	// 开票相关
	StatusCodeRequestInValid        uint16 = 9001 // 请求校验出错
	StatusCodeParameterInValid      uint16 = 9002 // 业务逻辑参数校验错误，修改参数后重试
	StatusCodeCallInvoiceError      uint16 = 9003 // 请求开票失败，调用失败，可重试
	StatusCodeServerConnectionError uint16 = 9004 // 服务器连接错误，可重试
	StatusCodeServerInnerError      uint16 = 9006 // 服务器内部错误，联系技术人员
	StatusCodeSellerNotExists       uint16 = 9010 // 服务商对应商户不存在，销方税号错误
	StatusCodeSellerNotAuthorize    uint16 = 9011 // 商户暂无开票权限
	StatusCodeInvoiceNotExists      uint16 = 9102 // 发票不存在
	// 插卡相关(不对接此功能)
	StatusCodeCardServiceError    uint16 = 7000 // 插卡服务错误
	StatusCodeCardParameterError  uint16 = 7001 // 参数错误，参数校验不通过
	StatusCodeServiceNotAuthorize uint16 = 7002 // 公众号未授权给高灯
	StatusCodeOrderNotAuthorize   uint16 = 7101 // 订单没有授权，可能是授权appid和插卡appid不一致，修改后重试
	StatusCodeCardSoon            uint16 = 7102 // 创建card_id与插卡时间间隔过短，可重试。注意不要每次插卡都创建card_id
	StatusCodeInvoiceCodeError    uint16 = 7103 // 发票号码和发票代码错误,修改后可重试
	StatusCodePdfInvalid          uint16 = 7104 // Pdf 无效，请提供真实有效的 pdf,修改后可重试。或者是pdf暂未生效， 稍后可重试
	StatusCodeWxError             uint16 = 7105 // 微信服务器处理失败，可直接重试
	StatusCodeInvoiceUpdating     uint16 = 7106 // 发票正在被修改状态，请稍后再试，可直接重试
	StatusCodeWxUnknownError      uint16 = 7200 // 微信其它未知错误
	StatusCodeInvoiceCodeRepeated uint16 = 7201 // 发票号码和发票代码重复，插卡重复，重试无效
	StatusCodeWxReturnError       uint16 = 7202 // 微信返回参数错误
)
