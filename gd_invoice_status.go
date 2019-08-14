package gaodeng

import "encoding/json"

// 查询发票信息
func (c *Client) InvoiceStatus(body InvoiceStatusRequest) (rsp InvoiceStatusResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/status", body)
	if err != nil {
		return
	}

	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceStatusRequest struct {
	TaxPayerNumber string `json:"taxpayer_num"`          // 开票商户的税号
	OrderId        string `json:"order_id"`              // 商户交易流水号(由商户维护唯一性)
	UniqueId       string `json:"g_unique_id,omitempty"` // 平台交易流水号(由平台维护唯一性)
	IsRed          uint8  `json:"is_red,omitempty"`      // 是否为红票(见constant定义)
}

type InvoiceStatusResponse struct {
	OrderId                 string `json:"order_id"`                    // 商户交易流水号(由商户维护唯一性)
	UniqueId                string `json:"g_unique_id"`                 // 平台交易流水号(由平台维护唯一性)
	Status                  uint8  `json:"status"`                      // 冲红状态(见constant定义)
	FailMessage             string `json:"fail_msg"`                    // 开票错误信息
	AuditPayState           string `json:"audit_pay_state"`             // 审核-支付状态
	AuditPayMessage         string `json:"audit_pay_msg"`               // 审核-支付提示消息
	TicketDate              string `json:"ticket_date"`                 // 开票日期
	TicketSn                string `json:"ticket_sn"`                   // 发票号码
	TicketCode              string `json:"ticket_code"`                 // 发票代码
	TicketTotalAmountHasTax string `json:"ticket_total_amount_has_tax"` // 发票含税总价,开票失败时值为空
	TicketTotalAmountNoTax  string `json:"ticket_total_amount_no_tax"`  // 发票去税总价,开票失败时值为空
	TicketTaxAmount         string `json:"ticket_tax_amount"`           // 发票税额,开票失败时值为空
	IsRed                   string `json:"is_red"`                      // 是否为红票(见constant定义)
	HasRed                  string `json:"has_red"`                     // 是否冲红
	GpPdfName               string `json:"gp_pdf_name"`                 // 发票pdf文件url
	CheckCode               string `json:"check_code"`                  // 发票校验码
}
