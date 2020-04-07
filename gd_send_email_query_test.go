package gaodeng

import (
	"fmt"
	"testing"
)

// 测试邮件发送查询
func testSendEmailQuery(t *testing.T, orderSn string) (err error) {
	fmt.Println("----------邮件发送查询----------")
	// 初始化参数
	body := SendEmailQueryRequest{
		SellerTaxpayerNum: TestTaxPayerNumber,
		OrderSn:           orderSn,
	}
	// 请求接口
	rsp, err := testClient.SendEmailQuery(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
