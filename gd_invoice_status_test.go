package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票查询
func testInvoiceStatus(t *testing.T, c *Client, orderSn string) (err error) {
	fmt.Println("----------发票查询----------")
	// 初始化参数
	body := InvoiceStatusRequest{
		SellerTaxPayerNumber: TestTaxPayerNumber,
		OrderSn:              orderSn,
	}
	// 请求接口
	rsp, err := c.InvoiceStatus(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
