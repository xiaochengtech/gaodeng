package gaodeng

import "encoding/json"

// 发票查验
func (c *Client) InvoiceVerify(body InvoiceVerifyRequest) (rsp InvoiceVerifyResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/tool/invoice-check", body)
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
	Additional string `json:"additional,omitempty"` // 发票校验码后6位（增值税专用发票，增值税机动车发票，二手车统一发票可以不传）
}

type InvoiceVerifyResponse struct {
	Status       uint8         `json:"status"`        // 销方名称
	OriginStatus string        `json:"origin_status"` // 详细状态(详见附录)
	Message      string        `json:"message"`       // 查询描述
	Invoice      InvoiceDetail `json:"invoice"`       // 发票信息
}

type InvoiceDetail struct {
	TicketCode         string       `json:"ticket_code"`          // 发票代码
	TicketSn           string       `json:"ticket_sn"`            // 发票号码
	TicketDate         string       `json:"ticket_date"`          // 开票日期
	BuyerName          string       `json:"buyer_name"`           // 销方名称
	BuyerTaxCode       string       `json:"buyer_tax_code"`       // 购方税号
	BuyerAddressPhone  string       `json:"buyer_address_phone"`  // 销方地址和电话
	BuyerBankAccount   string       `json:"buyer_bank_account"`   // 销方银行账号
	SellerName         string       `json:"seller_name"`          // 销方名称
	SellerTaxCode      string       `json:"seller_tax_code"`      // 购方税号
	SellerAddressPhone string       `json:"seller_address_phone"` // 销方地址和电话
	SellerBankAccount  string       `json:"seller_bank_account"`  // 销方银行账号
	Remark             string       `json:"remark"`               // 发票备注
	MachineNo          string       `json:"machine_no"`           // 税盘号
	InvoiceType        string       `json:"invoice_type"`         // 发票种类编码(见constant定义)
	CheckCode          string       `json:"check_code"`           // 发票校验码
	IsAbandoned        string       `json:"is_abandoned"`         // 是否作废(红冲)
	HasSellerList      string       `json:"has_seller_list"`      // 是否有销货清单
	SellerListTitle    string       `json:"seller_list_title"`    // 销货清单标题
	SellerListTax      string       `json:"seller_list_tax"`      // 销货清单税额
	AmountWithoutTax   string       `json:"amount_without_tax"`   // 不含税金额(元)
	TaxAmount          string       `json:"tax_amount"`           // 税额(元)
	AmountWithTax      string       `json:"amount_with_tax"`      // 含税金额(元)
	Items              []ItemDetail `json:"items"`                // 项目明细
}

type ItemDetail struct {
	LineNo           string `json:"line_no"`            // 项目行号
	Name             string `json:"name"`               // 商品名
	Model            string `json:"model"`              // 商品规格
	Unit             string `json:"unit"`               // 计量单位
	Number           string `json:"num"`                // 商品数量
	Price            string `json:"per_price"`          // 商品单价
	AmountWithoutTax string `json:"amount_without_tax"` // 不含税金额
	TaxRate          string `json:"tax_rate"`           // 税率
	TaxAmount        string `json:"tax_amount"`         // 税额
}
