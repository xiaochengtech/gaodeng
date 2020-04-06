package gaodeng

import (
	"encoding/json"
)

// 查询发票信息
func (c *Client) SendEmail(body SendEmailRequest) (rsp SendEmailResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/invoice/amount", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type SendEmailRequest struct {
	OrderId           string `json:"order_id,omitempty"`  // 商户订单号
	OrderSn           string `json:"order_sn,omitempty"`  // 高灯方商户订单号
	SellerTaxpayerNum string `json:"seller_taxpayer_num"` // 销方纳税人识别号
	IsRed             int    `json:"is_red,omitempty"`    // 是否为红票(0：否，1：是，默认为0)
	Email             string `email`                      // 收票人邮箱
}

type SendEmailResponse struct {
	status int `json:"number"` // 邮件状态(1：成功，2：失败)
}
