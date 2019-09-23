package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票余量
func TestInvoiceAmount(t *testing.T) {
	fmt.Println("----------发票余量 start----------")
	// 初始化参数
	body := InvoiceAmountRequest{
		TaxPayerNumber:  TestTaxPayerNumber,
		InvoiceTypeCode: InvoiceTypeCodeQKL,
	}
	// 请求接口
	wxRsp, err := testClient.InvoiceAmount(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	fmt.Println("----------发票余量 end----------")
}
