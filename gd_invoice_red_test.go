package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票冲红
func testInvoiceRed(t *testing.T, c *Client, orderSn string) (err error) {
	fmt.Println("----------发票冲红----------")
	// 初始化参数
	body := InvoiceRedRequest{
		Invoices: []Invoice{Invoice{
			SellerTaxPayerNumber: TestTaxPayerNumber,
			CallbackUrl:          "https://www.xiaochengtech.cn/",
			OrderSn:              orderSn,
		}},
	}
	// 请求接口
	rsp, err := c.InvoiceRed(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
