package gaodeng

import "encoding/json"

// 开具发票
func (c *Client) InvoiceBlue(body InvoiceBlueRequest) (rsp InvoiceBlueResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/blue", body)
	if err != nil {
		return
	}

	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceBlueRequest struct {
	SellerName           string `json:"seller_name,omitempty"`          // 销方名称
	SellerTaxPayerNumber string `json:"seller_taxpayer_num"`            // 销货方税号
	SellerAddress        string `json:"seller_address,omitempty"`       // 销货方地址
	SellerTel            string `json:"seller_tel,omitempty"`           // 销货方电话
	SellerBankName       string `json:"seller_bank_name,omitempty"`     // 销货方开户行
	SellerBankAccount    string `json:"seller_bank_account,omitempty"`  // 销货方银行账号
	TitleType            uint8  `json:"title_type"`                     // 发票抬头类型(见constant定义)
	BuyerTitle           string `json:"buyer_title"`                    // 购方抬头
	BuyerTaxPayerNumber  string `json:"buyer_taxpayer_num,omitempty"`   // 购方税号(企业类型的抬头必填）
	BuyerAddress         string `json:"buyer_address,omitempty"`        // 购方地址
	BuyerBankName        string `json:"buyer_bank_name,omitempty"`      // 购方开户行
	BuyerBankAccount     string `json:"buyer_bank_account,omitempty"`   // 购方银行账号
	BuyerPhone           string `json:"buyer_phone,omitempty"`          // 购方电话
	BuyerEMail           string `json:"buyer_email,omitempty"`          // 购方邮箱
	TakerPhone           string `json:"taker_phone,omitempty"`          // 收票人手机
	TakerName            string `json:"taker_name,omitempty"`           // 收票人名称
	OrderId              string `json:"order_id"`                       // 商户订单号
	InvoiceTypeCode      string `json:"invoice_type_code,omitempty"`    // 发票种类编码(见constant定义)
	CallbackUrl          string `json:"callback_url"`                   // 接收开票平台推送的消息地址
	Drawer               string `json:"drawer,omitempty"`               // 开票人
	Payee                string `json:"payee,omitempty"`                // 收款人
	Checker              string `json:"checker,omitempty"`              // 复核人
	TradeType            string `json:"trade_type,omitempty"`           // 行业分类(见constant定义)
	MachineNo            string `json:"machine_no,omitempty"`           // 税盘号
	UserOpenId           string `json:"user_openid,omitempty"`          // 三方用户id
	SpecialInvoiceKind   string `json:"special_invoice_kind,omitempty"` // 特殊票种标识（成品油票必传：08，其他票种可以为空）
	TerminalCode         string `json:"terminal_code,omitempty"`        // 开票终端代码,使用百望税控服务器时必填
	AmountHasTax         uint32 `json:"amount_has_tax"`                 // 含税金额(单分)
	TaxAmount            uint32 `json:"tax_amount"`                     // 税额(单分)
	AmountWithoutTax     uint32 `json:"amount_without_tax"`             // 不含税金额(单分)
	Remark               string `json:"remark,omitempty"`               // 发票备注
	Items                []Item `json:"items"`                          // 商品数组
}

type Item struct {
	Name                   string `json:"name"`                               // 商品名(见constant定义)
	TaxCode                string `json:"tax_code"`                           // 税目编码(见constant定义)
	TaxType                string `json:"tax_type,omitempty"`                 // 税目类别
	Models                 string `json:"models,omitempty"`                   // 商品规格
	Unit                   string `json:"unit,omitempty"`                     // 计量单位
	TotalPrice             uint32 `json:"total_price"`                        // 不含税商品总金额（精确到2位）
	Total                  string `json:"total"`                              // 商品数量
	Price                  string `json:"price"`                              // 不含税商品单价
	TaxRate                uint32 `json:"tax_rate"`                           // 税率(千分位,税率*1000)
	TaxAmount              uint32 `json:"tax_amount"`                         // 税额(单分)
	Discount               int32  `json:"discount,omitempty"`                 // 总的折扣金额；金额必须是负数；无折扣时不传
	ZeroTaxFlag            string `json:"zero_tax_flag,omitempty"`            // 零税率标识，默认为空(见constant定义)
	PreferentialPolicyFlag string `json:"preferential_policy_flag,omitempty"` // 优惠政策标志(见constant定义)
	VatSpecialManagement   string `json:"vat_special_management,omitempty"`   // 增值税特殊管理
}

type InvoiceBlueResponse struct {
	InvoiceId string `json:"invoice_id"` // 高灯发票唯一识别号
	OrderSn   string `json:"order_sn"`   // 高灯订单号(红冲需要用到)
	Status    uint8  `json:"status"`     // 发票状态(见constant定义)
}
