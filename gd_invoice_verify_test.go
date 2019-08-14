package gaodeng

import (
	"fmt"
	"testing"
)

// 测试发票查验
func TestInvoiceVerify(t *testing.T) {
	fmt.Println("----------发票查验 start----------")
	// 初始化参数
	body := InvoiceVerifyRequest{
		TicketCode: "011001800304",
		TicketSn:   "02703183",
		TicketDate: "2019-07-23",
	}

	// 请求接口
	wxRsp, err := testClient.InvoiceVerify(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	fmt.Println("----------发票查验 end----------")
}
