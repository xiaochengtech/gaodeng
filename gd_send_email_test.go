package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票余量
func testSendEmail(t *testing.T, orderId string, orderSn string, isRed uint8) (err error) {
	fmt.Println("----------发票余量----------")
	// 初始化参数
	body := SendEmailRequest{
		SellerTaxpayerNum: TestTaxPayerNumber,
		Email:             TestEmail,
		IsRed:             isRed,
		OrderId:           orderId,
		OrderSn:           orderSn,
	}
	// 请求接口
	rsp, err := testClient.SendEmail(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
