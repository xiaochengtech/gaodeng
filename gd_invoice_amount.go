package gaodeng

import "encoding/json"

// 查询发票信息
func (c *Client) InvoiceAmount(body InvoiceAmountRequest) (rsp InvoiceAmountResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/amount", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type InvoiceAmountRequest struct {
	TaxPayerNumber  string `json:"taxpayer_num"`                // 开票商户的税号
	MachineNo       string `json:"machine_no,omitempty"`        // 税盘号,使用云票儿或百望税控服务器时必填
	TerminalCode    string `json:"terminal_code,omitempty"`     // 开票终端代码,使用百望税控服务器时必填
	InvoiceTypeCode string `json:"invoice_type_code,omitempty"` // 发票种类编码(见constant定义)
}

type InvoiceAmountResponse struct {
	Result string `json:"result"` // 查询结果(见constant定义)
	Number uint32 `json:"number"` // 剩余发票数量
}
