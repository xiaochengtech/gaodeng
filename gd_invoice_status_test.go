package gaodeng

import (
	"testing"
)

// 测试发票查询
func TestInvoiceStatus(t *testing.T) {
	t.Log("----------发票查询----------")
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
	t.Logf("返回值: %+v", wxRsp)
}
