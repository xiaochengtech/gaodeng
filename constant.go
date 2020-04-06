package gaodeng

const (
	EnvProd = "prod" // (生产环境)
	EnvTest = "test" // (沙盒环境)

	AppVersion = "1.0.0"

	InvoiceTypeBlue uint8 = 0 // 发票种类-蓝票
	InvoiceTypeRed  uint8 = 1 // 发票种类-红票

	InvoiceStatusDoing   uint8 = 1 // 开票中
	InvoiceStatusSuccess uint8 = 2 // 开票成功
	InvoiceStatusFailure uint8 = 3 // 开票失败

	InvoiceTitleTypePerson     uint8 = 1 // 个人
	InvoiceTitleTypeEnterprise uint8 = 2 // 企业

	TradeTypeElse             = "0" // 其他
	TradeTypeCommunication    = "1" // 通信
	TradeTypeFood             = "2" // 餐饮
	TradeTypeTransportation   = "3" // 交通
	TradeTypePayment          = "4" // 支付平台
	TradeTypeTicketAndTourist = "5" // 票务/旅游

	InvoiceTypeCodeZZSZY = "004" // 增值税专用发票
	InvoiceTypeCodeZZSPT = "007" // 增值税普通发票
	InvoiceTypeCodeZZSJS = "025" // 增值税卷式发票
	InvoiceTypeCodeZZSDZ = "026" // 增值税电子发票(默认)
	InvoiceTypeCodeQKL   = "032" // 区块链发票

	ZeroTaxNo     = ""  // 非零税率
	ZeroTaxExport = "0" // 出口零税
	ZeroTaxFree   = "1" // 免税
	ZeroTaxNone   = "2" // 不征税
	ZeroTaxCommon = "3" // 普通零税率

	PreferentialPolicyNo  = ""  // 不使用
	PreferentialPolicyUse = "1" // 使用

	RedStateSuccess uint8 = 1 // 调用成功
	RedStateFailure uint8 = 0 // 调用失败

	TicketDoing   = "1" // 开票中
	TicketSuccess = "2" // 开票成功
	TicketFailure = "3" // 开票失败

	VerifyCodeTrue uint8 = 0 // 成功（发票验证为真）
	VerifyCodeFail uint8 = 1 // 成功（发票验证为假）

	SendEmailStatusSuccess = 1 // 发送邮件成功
	SendEmailStatusFailure = 2 // 发送邮件失败

	CallbackNotifyTypeBlue = "invoice.blue" // 蓝票通知类型
	CallbackNotifyTypeRed  = "invoice.red"  // 红票通知类型

	CallbackTicketStatusSuccess uint8 = 2 // 开票成功，其他都是失败

	StatusCodeNormal               uint16 = 0    // 正常
	StatusCodeRequestError         uint16 = 400  // 请求错误,服务器无法理解http请求
	StatusCodeNotPermission        uint16 = 403  // 禁止访问
	StatusCodeServerInnerError     uint16 = 500  // 服务器内部错误
	StatusCodeVersionError         uint16 = 1001 // 版本错误
	StatusCodeLackCommonParameters uint16 = 1002 // 缺少公共参数
	StatusCodeTimestampError       uint16 = 1003 // 时间戳错误
	StatusCodeAppKeyNotExists      uint16 = 1004 // appkey不存在
	StatusCodeAppClosed            uint16 = 1005 // app已关闭
	StatusCodeDeveloperClosed      uint16 = 1006 // 开发者已被关闭
	StatusCodeSignError            uint16 = 1007 // 签名错误
	StatusCodeDeveloperUnAudit     uint16 = 1008 // 开发者资质待审核
	StatusCodeDeveloperNoAudit     uint16 = 1009 // 开发者资质审核未通过
	StatusCodeParameterInValid     uint16 = 1010 // 参数检验错误
	StatusCodeDirtyData            uint16 = 1011 // 脏数据
	StatusCodeCallNumberNotEnough  uint16 = 1012 // 接口调用次数不足
	StatusCodeInterfaceFeeError    uint16 = 1013 // 接口计费错误
	StatusCodearameterError        uint16 = 1014 // 业务参数错误
	StatusCodeDependDeny           uint16 = 1015 // 依赖服务返回拒绝
	StatusCodeDependError          uint16 = 1016 // 依赖服务故障
)
