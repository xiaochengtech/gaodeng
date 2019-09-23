package gaodeng

import (
	"fmt"
	"testing"
)

// 测试开具发票
func TestInvoiceBlue(t *testing.T) {
	fmt.Println("----------开具发票 start----------")
	// 初始化参数
	body := InvoiceBlueRequest{
		SellerTaxPayerNumber: TestTaxPayerNumber,
		SellerAddress:        TestSellerAddress,
		TitleType:            InvoiceTitleTypePerson,
		BuyerTitle:           "张三",
		OrderId:              "cuckoo20190813180503001",
		InvoiceTypeCode:      InvoiceTypeCodeQKL,
		CallbackUrl:          "https://www.cuckoopark.com/",
		AmountHasTax:         1500,
		TaxAmount:            135,
		AmountWithoutTax:     1365,
		Items: []Item{Item{
			Name:       GoodsName,
			TaxCode:    TaxCodeParking,
			TotalPrice: 1500,
			Total:      "1",
			Price:      "15.00",
			TaxRate:    TaxRateParking,
			TaxAmount:  135,
		}},
	}
	// 请求接口
	wxRsp, err := testClient.InvoiceBlue(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	fmt.Println("----------开具发票 end----------")
}
