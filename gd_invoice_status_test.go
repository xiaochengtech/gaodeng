package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票查询
func TestInvoiceStatus(t *testing.T) {
	fmt.Println("----------发票查询 start----------")
	// 初始化参数
	body := InvoiceStatusRequest{
		SellerTaxPayerNumber: TestTaxPayerNumber,
		OrderSn:              "6566985919163970279",
	}

	// 请求接口
	wxRsp, err := testClient.InvoiceStatus(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	fmt.Println("----------发票查询 end----------")
}
