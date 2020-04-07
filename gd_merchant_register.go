package gaodeng

import (
	"encoding/json"
)

// 商户注册
func (c *Client) MerchantRegister(body MerchantRegisterRequest) (rsp MerchantRegisterResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/merchant/register", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type MerchantRegisterRequest struct {
	SerialNo                   string `json:"serial_no,omitempty"`          // 请求流水号（强烈建议传该字段，回调通知的时候会通过该字段标识）
	TaxpayerName               string `json:"taxpayer_name"`                // 企业名称
	TaxpayerNum                string `json:"taxpayer_num"`                 // 纳税人识别号(税号)
	LegalPersonName            string `json:"legal_person_name"`            // 注册企业法人代表名称
	ContactsName               string `json:"contacts_name"`                // 联系人姓名
	Email                      string `json:"email,omitempty"`              // 联系人邮箱地址
	BusinessMobile             string `json:"business_mobile,omitempty"`    // 企业电话
	Phone                      string `json:"phone"`                        // 联系人手机号
	BankName                   string `json:"bank_name,omitempty"`          // 银行名称
	BankAccount                string `json:"bank_account,omitempty"`       // 银行账号
	Address                    string `json:"address"`                      // 不包含省市名称的地址
	RegionCode                 uint8  `json:"region_code"`                  // 地区编码(见constant定义, 此处官方api文档中的类型有误，应该是int，不是string)
	CityName                   string `json:"city_name"`                    // 市(地区)名称
	Drawer                     string `json:"Drawer"`                       // 开票人
	Reviewer                   string `json:"reviewer,omitempty"`           // 复核人
	Payee                      string `json:"payee,omitempty"`              // 收款人
	RegisterCode               string `json:"register_code,omitempty"`      // 注册邀请码
	TaxRegistrationCertificate string `json:"tax_registration_certificate"` // 税务登记证图片
	State                      string `json:"state,omitempty"`              // 开票状态(见constant定义)
	CallbackUrl                string `json:"callback_url"`                 // 接收推送的消息地址
}

type MerchantRegisterResponse struct {
	TaxpayerName string `json:"taxpayer_name"` // 纳税人识别号
	SerialNo     int    `json:"serial_no"`     // 请求流水号
}
