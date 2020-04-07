package gaodeng

import (
	"encoding/json"
)

// 邮件发送查询
func (c *Client) SendEmailQuery(body SendEmailQueryRequest) (rsp SendEmailQueryResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/private/send-email-query", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = json.Unmarshal(bytes, &rsp)
	return
}

type SendEmailQueryRequest struct {
	OrderId           string `json:"order_id,omitempty"`  // 商户订单号
	OrderSn           string `json:"order_sn,omitempty"`  // 高灯方商户订单号
	SellerTaxpayerNum string `json:"seller_taxpayer_num"` // 销方纳税人识别号
}

type SendEmailQueryResponse struct {
	Email    string `json:"email"`     // 邮箱
	PushedAt int    `json:"pushed_at"` // 邮件发送时间戳, 为0表示发送邮件失败或未发送
}
