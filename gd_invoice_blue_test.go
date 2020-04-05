package gaodeng

import (
	"fmt"
	"testing"
)

// 测试开具发票
func testInvoiceBlue(t *testing.T, c *Client, orderId string) (rsp InvoiceBlueResponse, err error) {
	fmt.Println("----------开具发票----------")
	// 初始化参数
	body := InvoiceBlueRequest{
		SellerTaxPayerNumber: TestTaxPayerNumber,
		SellerAddress:        TestSellerAddress,
		TitleType:            InvoiceTitleTypePerson,
		BuyerTitle:           "张三",
		OrderId:              orderId,
		InvoiceTypeCode:      InvoiceTypeCodeZZSDZ,
		TerminalCode:         "1",
		CallbackUrl:          "https://www.xiaochengtech.cn/",
		AmountHasTax:         1500,
		TaxAmount:            124,
		AmountWithoutTax:     1376,
		Items: []Item{{
			Name:       "计时停车",
			TaxCode:    "3040502020200000000",
			TotalPrice: 1376,
			Total:      "1",
			Price:      "13.76",
			TaxRate:    90,
			TaxAmount:  124,
		}},
	}
	// 请求接口
	rsp, err = c.InvoiceBlue(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
