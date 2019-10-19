package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票冲红
func TestInvoiceRed(t *testing.T) {
	fmt.Println("----------发票冲红----------")
	// 初始化参数
	body := InvoiceRedRequest{
		Invoices: []Invoice{Invoice{
			SellerTaxPayerNumber: TestTaxPayerNumber,
			CallbackUrl:          "https://www.cuckoopark.com/",
			OrderSn:              "6566985919163970279",
		}},
	}
	// 请求接口
	wxRsp, err := testClient.InvoiceRed(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
