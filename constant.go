package gaodeng

const (
	EnvProd = "prod" // (生产环境)
	EnvTest = "test" // (沙盒环境)

	TestAppKey         = "EgDjckWzyGxwIi7e9J1A8LdruWMidFFH"
	TestAppSecret      = "9Q8744Oe0nv8aw738b3HkjdylYZzNeZOcTz53KI4pchKpqIi"
	TestTaxPayerNumber = "91440300661005378A"
	TestSellerName     = "深圳市卓越物业管理股份有限公司卓越时代广场地下停车场"
	TestSellerAddress  = "深圳市福田区益田路与福华路交汇处卓越时代广场地下"

	InvoiceStatusDoing   int8 = 1  // 开票中
	InvoiceStatusSuccess int8 = 2  // 开票成功
	InvoiceStatusFailure int8 = -2 // 开票失败

	InvoiceTitleTypePerson     uint8 = 1 // 个人
	InvoiceTitleTypeEnterprise uint8 = 2 // 企业

	TradeTypeElse             string = "0" // 其他
	TradeTypeCommunication    string = "1" // 通信
	TradeTypeFood             string = "2" // 餐饮
	TradeTypeTransportation   string = "3" // 交通
	TradeTypePayment          string = "4" // 支付平台
	TradeTypeTicketAndTourist string = "5" // 票务/旅游

	InvoiceTypeCodeZZSZY string = "004" // 增值税专用发票
	InvoiceTypeCodeZZSPT string = "007" // 增值税普通发票
	InvoiceTypeCodeZZSJS string = "025" // 增值税卷式发票
	InvoiceTypeCodeZZSDZ string = "026" // 增值税电子发票(默认)
	InvoiceTypeCodeQKL   string = "032" // 区块链发票

	GoodsName      string = "计时停车"
	TaxCodeParking string = "3040502020200000000"
	TaxRateParking string = "0.09"

	ZeroTaxNo     string = ""  // 非零税率
	ZeroTaxExport string = "0" // 出口零税
	ZeroTaxFree   string = "1" // 免税
	ZeroTaxNone   string = "2" // 不征税
	ZeroTaxCommon string = "3" // 普通零税率

	PreferentialPolicyNo  string = ""  // 不使用
	PreferentialPolicyUse string = "1" // 使用
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
