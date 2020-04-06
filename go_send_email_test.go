package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票余量
func TestSendEmail(t *testing.T) {
	fmt.Println("----------发票余量----------")
	// 初始化参数
	body := SendEmailRequest{
		SellerTaxpayerNum: TestSellerTaxpayerNum,
		Email:             TestEmail,
	}
	// 请求接口
	rsp, err := testClient.SendEmail(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
}
