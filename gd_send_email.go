package gaodeng

import (
	"encoding/json"
)

// 查询发票信息
func (c *Client) SendEmail(body SendEmailRequest) (rsp SendEmailResponse, err error) {
	// 业务逻辑
	bytes, err := c.post("/private/send-email", body)
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
	IsRed             uint8  `json:"is_red,omitempty"`    // 是否为红票(0：蓝票，1：红票，默认为0, 参照constant中关于发票种类的定义)
	Email             string `json:"email"`               // 收票人邮箱
}

type SendEmailResponse struct {
	Status int `json:"status"` // 邮件状态(1：成功，2：失败)
}
