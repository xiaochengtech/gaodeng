package gaodeng

import "encoding/json"

// 发票查验
func (c *Client) InvoiceVerify(body InvoiceVerifyRequest) (rsp InvoiceVerifyResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/verify", body)
	if err != nil {
		return
	}

	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceVerifyRequest struct {
	TicketCode string `json:"ticket_code"`          // 发票代码
	TicketSn   string `json:"ticket_sn"`            // 发票号码
	TicketDate string `json:"ticket_date"`          // 开票日期
	CheckCode  string `json:"check_code,omitempty"` // 发票校验码（增值税专用发票，增值税机动车发票，二手车统一发票可以不传）
	TicketTax  string `json:"ticket_tax,omitempty"` // 开具金额、不含税价（增值税普通发票，增值税电子发票，卷式发票，电子普通[通行费]发票可以不传）
}

type InvoiceVerifyResponse struct {
	TicketCode              string        `json:"ticket_code"`                 // 发票代码
	TicketSn                string        `json:"ticket_sn"`                   // 发票号码
	TicketDate              string        `json:"ticket_date"`                 // 开票日期
	TicketTax               string        `json:"ticket_tax"`                  // 不含税金额
	CheckCode               string        `json:"check_code"`                  // 发票校验码
	InvoiceTypeName         string        `json:"invoice_type_name"`           // 发票类型
	HasRed                  string        `json:"has_red"`                     // 是否冲红
	InvoiceTypeCode         string        `json:"invoice_type_code"`           // 发票种类编码(见constant定义)
	TicketTotalAmountHasTax string        `json:"ticket_total_amount_has_tax"` // 发票含税总价
	TicketTotalAmountNoTax  string        `json:"ticket_total_amount_no_tax"`  // 发不含税金额
	TotalTaxAmount          string        `json:"total_tax_amount"`            // 总税额(最多保留两位小数)
	SellName                uint8         `json:"sell_name"`                   // 销方名称
	SellTaxCode             string        `json:"sell_tax_code"`               // 购方税号
	SellAddressPhone        string        `json:"sell_address_phone"`          // 销方地址和电话
	SellBankNameAccount     string        `json:"sell_bank_name_account"`      // 销方银行账号
	BuyName                 string        `json:"buy_name"`                    // 购方名称
	BuyTaxCode              string        `json:"buy_tax_code"`                // 购方税号
	BuyAddress              string        `json:"buy_address"`                 // 购方地址
	BuyBankAccount          string        `json:"buy_bank_account"`            // 购方银行账号
	Extra                   string        `json:"extra"`                       // 发票备注
	GoodsInfo               []GoodsDetail `json:"goods_info"`                  // 商品数组
}

type GoodsDetail struct {
	Name        string `json:"name"`          // 商品名
	Models      string `json:"models"`        // 商品规格
	Unit        string `json:"unit"`          // 计量单位
	TotalPrice  string `json:"total_price"`   // 含税商品总金额
	Total       string `json:"total"`         // 商品数量
	Price       string `json:"per_price"`     // 商品单价
	TaxCode     string `json:"tax_code"`      // 税目编码(见constant定义)
	TaxType     string `json:"tax_type"`      // 税目类别
	TaxRate     string `json:"tax_rate"`      // 税率，范围0-1(见constant定义)
	CTotalPrice string `json:"c_total_price"` // 不含税商品总金额
	TaxAmount   string `json:"tax_amount"`    // 税额
	HasTax      string `json:"has_tax"`       // 是否含税。0否，1是
	ZeroTaxFlag string `json:"zero_tax_flag"` // 零税率标识，默认为空(见constant定义)
	Property    string `json:"fp_property"`   // 发票行性质
}
