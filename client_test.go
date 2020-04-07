package gaodeng

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	TestAppKey         = os.Getenv("GDTestAppKey")
	TestAppSecret      = os.Getenv("GDTestAppSecret")
	TestTaxPayerNumber = os.Getenv("GDTestTaxPayerNumber")
	TestSellerName     = os.Getenv("GDTestSellerName")
	TestSellerAddress  = os.Getenv("GDTestSellerAddress")
	TestEmail          = os.Getenv("GDTestEmail")
)

var testClient = NewClient(EnvTest, Config{
	AppKey:    TestAppKey,
	AppSecret: TestAppSecret,
})

// 测试完整流程
func TestClient(t *testing.T) {
	now := time.Now()
	orderId := fmt.Sprintf("%s%09d", now.Format("060102150405"), now.UnixNano()%1e9)
	// 发票开具
	blueRsp, err := testInvoiceBlue(t, testClient, orderId)
	if err != nil {
		t.Error(err)
		return
	}
	// 查询发票状态
	err = testInvoiceStatus(t, testClient, blueRsp.OrderSn)
	if err != nil {
		t.Error(err)
		return
	}
	// 发送邮件
	err = testSendEmail(t, blueRsp.OrderSn, 0)
	// 发票红冲
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second) // 加延时保证发票已经开具
	err = testInvoiceRed(t, testClient, blueRsp.OrderSn)
	if err != nil {
		t.Error(err)
		return
	}
}
