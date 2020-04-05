package gaodeng

import (
	"testing"
)

// 测试发票冲红
func TestInvoiceRed(t *testing.T) {
	t.Log("----------发票冲红----------")
	// 初始化参数
	body := InvoiceRedRequest{
		Invoices: []Invoice{Invoice{
			SellerTaxPayerNumber: TestTaxPayerNumber,
			CallbackUrl:          "https://www.xiaochengtech.cn/",
			OrderSn:              "6566985919163970279",
		}},
	}
	// 请求接口
	wxRsp, err := testClient.InvoiceRed(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v", wxRsp)
}
