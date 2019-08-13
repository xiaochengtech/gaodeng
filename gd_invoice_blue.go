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

// 关闭订单的参数
type InvoiceBlueRequest struct {
	TaxPayerNumber     string  `json:"taxpayer_num"`                   // 销货方纳税人识别号(15~20位数字或者大写字母)
	SellerAddress      string  `json:"seller_address,omitempty"`       // 销货方地址
	SellerTel          string  `json:"seller_tel,omitempty"`           // 销货方电话
	SellerBankName     string  `json:"seller_bank_name,omitempty"`     // 销货方开户行
	SellerBankAccount  string  `json:"seller_bank_account,omitempty"`  // 销货方银行账号
	BuyerTitleType     uint8   `json:"buyer_title_type"`               // 发票抬头类型(见constant定义)
	OrderId            string  `json:"order_id"`                       // 商户订单号
	UserId             string  `json:"b_user_id,omitempty"`            // 用户在服务商测标识
	BuyerTaxCode       string  `json:"buyer_taxcode,omitempty"`        // 购货方识别号（当开票抬头类型为 2时：必填）
	BuyerTitle         string  `json:"buyer_title"`                    // 购方名称
	BuyerBankName      string  `json:"buyer_bank_name,omitempty"`      // 购方开户行
	BuyerBankAccount   string  `json:"buyer_bank_account,omitempty"`   // 购方银行账号
	BuyerPhone         string  `json:"buyer_phone,omitempty"`          // 购方电话
	BuyerEMail         string  `json:"buyer_email,omitempty"`          // 购方邮箱
	BuyerAddress       string  `json:"buyer_address,omitempty"`        // 购方地址
	Extra              string  `json:"extra,omitempty"`                // 发票备注
	CallbackUrl        string  `json:"callback_url"`                   // 接收开票平台推送的消息地址
	Cashier            string  `json:"cashier,omitempty"`              // 收银员名称
	Checker            string  `json:"checker,omitempty"`              // 复核员名称
	Invoicer           string  `json:"invoicer,omitempty"`             // 开票员名称
	TradeType          string  `json:"trade_type,omitempty"`           // 行业分类(见constant定义)
	MachineNo          string  `json:"machine_no,omitempty"`           // 税盘号
	EtrData            string  `json:"etr_data,omitempty"`             // 预留字段
	InvoiceTypeCode    string  `json:"invoice_type_code,omitempty"`    // 发票种类编码(见constant定义)
	SpecialInvoiceKind string  `json:"special_invoice_kind,omitempty"` // 特殊票种标识（成品油票必传：08，其他票种可以为空）
	TerminalCode       string  `json:"terminal_code,omitempty"`        // 开票终端代码
	TotalAmountHasTax  string  `json:"total_amount_has_tax"`           // 总含税金额(最多保留两位小数)
	TotalTaxAmount     string  `json:"total_tax_amount"`               // 总税额(最多保留两位小数)
	TotalAmountNoTax   string  `json:"total_amount_no_tax"`            // 总不含税金额(最多保留两位小数)
	GoodsInfo          []Goods `json:"goods_info"`                     // 商品数组
}

type Goods struct {
	Name                   string `json:"name"`                               // 商品名(见constant定义)
	TaxCode                string `json:"tax_code"`                           // 税目编码(见constant定义)
	TaxType                string `json:"tax_type,omitempty"`                 // 税目类别
	Models                 string `json:"models,omitempty"`                   // 商品规格
	Unit                   string `json:"unit,omitempty"`                     // 计量单位
	TotalPrice             string `json:"total_price"`                        // 不含税商品总金额（精确到2位）
	Total                  string `json:"total"`                              // 商品数量
	Price                  string `json:"price"`                              // 不含税商品单价
	TaxRate                string `json:"tax_rate"`                           // 税率，范围0-1(见constant定义)
	TaxAmount              string `json:"tax_amount"`                         // 税额
	Discount               string `json:"discount,omitempty"`                 // 总的折扣金额；金额必须是负数；无折扣时不传
	ZeroTaxFlag            string `json:"zero_tax_flag,omitempty"`            // 零税率标识，默认为空(见constant定义)
	PreferentialPolicyFlag string `json:"preferential_policy_flag,omitempty"` // 优惠政策标志(见constant定义)
	VatSpecialManagement   string `json:"vat_special_management,omitempty"`   // 增值税特殊管理
	EtrData                string `json:"etr_data,omitempty"`                 // 预留字段
}

// 关闭订单的返回值
type InvoiceBlueResponse struct {
	UniqueId string `json:"g_unique_id"` // 平台交易流水号(由平台维护唯一性)
	OrderId  string `json:"order_id"`    // 商户交易流水号(由商户维护唯一性)
	State    uint8  `json:"state"`       // 发票状态(见constant定义)
}
