package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票冲红
func TestInvoiceRed(t *testing.T) {
	fmt.Println("----------发票冲红 start----------")
	// 初始化参数
	body := InvoiceRedRequest{
		CallbackUrl: "https://www.cuckoopark.com/",
		Invoices: []Invoice{Invoice{
			TaxPayerNumber: TestTaxPayerNumber,
			BTradeNo:       "cuckoo20190813180503001",
			GTradeNo:       "6566985919163970279",
		}},
	}

	// 请求接口
	wxRsp, err := testClient.InvoiceRed(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	fmt.Println("----------发票冲红 end----------")
}
