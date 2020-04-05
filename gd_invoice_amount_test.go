package gaodeng

import (
	"testing"
)

// 测试发票余量
func TestInvoiceAmount(t *testing.T) {
	t.Log("----------发票余量----------")
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
	t.Logf("返回值: %+v", wxRsp)
}
