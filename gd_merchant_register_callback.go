package gaodeng

type MerchantRegisterCallbackBody struct {
	AppKey            string `json:"appkey"`                       // 服务商appkey
	Code              string `json:"code"`                         // 返回状态码
	Message           string `json:"message"`                      // 返回描述信息
	SerialNo          string `json:"serial_no"`                    // 请求流水号
	TaxpayerName      string `json:"taxpayer_name"`                // 销方名称
	TaxpayerNum       string `json:"taxpayer_num"`                 // 销方纳税人识别号
	PlatformCode      string `json:"platform_code,omitempty"`      // 企业平台编码
	RegistrationCode  string `json:"registration_code,omitempty"`  // 企业注册码
	AuthorizationCode string `json:"authorization_code,omitempty"` // 企业授权码
	ExpressName       string `json:"express_name,omitempty"`       // 快递名称
	ExpressNo         string `json:"express_no,omitempty"`         // 快递单号
}
