package gaodeng

import "encoding/json"

// 发票冲红
func (c *Client) InvoiceRed(body InvoiceRedRequest) (rsp InvoiceRedResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/red", body)
	if err != nil {
		return
	}

	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceRedRequest struct {
	Invoices []Invoice `json:"invoices"` // 发票组
}

type Invoice struct {
	SellerTaxPayerNumber string `json:"seller_taxpayer_num"` // 销方税号
	CallbackUrl          string `json:"callback_url"`        // 红票接收地址
	OrderSn              string `json:"order_sn,omitempty"`  // 需要红冲的高灯订单号(蓝票开具时获得),和order_id不能同时为空
	OrderId              string `json:"order_id,omitempty"`  // 需要红冲的三方自有订单号(蓝票开具时传入),和order_sn不能同时为空
}

type InvoiceRedResponse = []InvoiceNoRed

type InvoiceNoRed struct {
	Code      uint8  `json:"code"`       // 冲红状态(见constant定义)
	Message   string `json:"message"`    // 提交红冲描述
	OrderSn   string `json:"order_sn"`   // 高灯红票订单号(常与蓝票订单号一致)
	InvoiceId string `json:"invoice_id"` // 高灯发票唯一识别号
}
