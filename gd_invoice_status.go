package gaodeng

import "encoding/json"

// 查询发票信息
func (c *Client) InvoiceStatus(body InvoiceStatusRequest) (rsp InvoiceStatusResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/query", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceStatusRequest struct {
	SellerTaxPayerNumber string `json:"seller_taxpayer_num"` // 销货方税号
	OrderId              string `json:"order_id,omitempty"`  // 三方订单号(和order_sn二选一必填)
	OrderSn              string `json:"order_sn,omitempty"`  // 平台订单号(和order_id二选一必填)
	IsRed                uint8  `json:"is_red,omitempty"`    // 发票种类(见constant定义)
}

type InvoiceStatusResponse struct {
	OrderId          string `json:"order_id"`           // 三方订单号
	OrderSn          string `json:"order_sn,omitempty"` // 平台订单号
	Status           uint8  `json:"status"`             // 冲红状态(见constant定义)
	Message          string `json:"message"`            // 开票错误信息
	TicketDate       string `json:"ticket_date"`        // 开票日期
	TicketSn         string `json:"ticket_sn"`          // 发票号码
	TicketCode       string `json:"ticket_code"`        // 发票代码
	CheckCode        string `json:"check_code"`         // 发票校验码
	AmountWithTax    string `json:"amount_with_tax"`    // 含税金额(元)
	AmountWithoutTax string `json:"amount_without_tax"` // 不含税金额(元)
	TaxAmount        string `json:"tax_amount"`         // 税额(元)
	IsRed            uint8  `json:"is_red_washed"`      // 是否被红冲
	PdfUrl           string `json:"pdf_url"`            // 发票pdf文件url
}
