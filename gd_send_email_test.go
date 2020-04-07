package gaodeng

import (
	"fmt"
	"testing"
)

// 测试邮件发送
func testSendEmail(t *testing.T, orderSn string, isRed uint8) (err error) {
	fmt.Println("----------邮件发送----------")
	// 初始化参数
	body := SendEmailRequest{
		SellerTaxpayerNum: TestTaxPayerNumber,
		Email:             TestEmail,
		IsRed:             isRed,
		OrderSn:           orderSn,
	}
	// 请求接口
	rsp, err := testClient.SendEmail(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
