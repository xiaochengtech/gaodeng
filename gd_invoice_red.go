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
	CallbackUrl string    `json:"callback_url"` // 接收开票平台推送的消息地址
	Invoices    []Invoice `json:"invoice"`      // 发票组
}

type Invoice struct {
	TaxPayerNumber string `json:"taxpayer_num"`         // 开票商户的税号
	BTradeNo       string `json:"b_trade_no,omitempty"` // 商户交易流水号
	GTradeNo       string `json:"g_trade_no,omitempty"` // 高灯开票平台开票流水号
}

type InvoiceRedResponse = []InvoiceNoRed

type InvoiceNoRed struct {
	TaxPayerNumber string `json:"taxpayer_num"`   // 开票商户的税号
	BTradeNo       string `json:"b_trade_no"`     // 商户交易流水号
	GTradeNo       string `json:"g_trade_no"`     // 高灯开票平台开票流水号
	State          uint8  `json:"state"`          // 冲红状态(见constant定义)
	Message        string `json:"msg"`            // 消息提示
	RedGTradeNo    string `json:"red_g_trade_no"` // 红票对应的平台唯一流水号
}
