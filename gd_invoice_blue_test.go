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
		TaxPayerNumber:    TestTaxPayerNumber,
		SellerAddress:     TestSellerAddress,
		BuyerTitleType:    InvoiceTitleTypePerson,
		OrderId:           "cuckoo20190813180503001",
		BuyerTitle:        "张三",
		CallbackUrl:       "https://www.cuckoopark.com/",
		InvoiceTypeCode:   InvoiceTypeCodeQKL,
		TotalAmountHasTax: "15.00",
		TotalTaxAmount:    "1.35",
		TotalAmountNoTax:  "13.65",
		GoodsInfo: []Goods{Goods{
			Name:       GoodsName,
			TaxCode:    TaxCodeParking,
			TotalPrice: "15.00",
			Total:      "1",
			Price:      "15.00",
			TaxRate:    TaxRateParking,
			TaxAmount:  "1.35",
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
