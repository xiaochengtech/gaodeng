package gaodeng

import (
	"fmt"
	"testing"
)

// 测试商户注册
func TestMerchantRegister(t *testing.T) {
	fmt.Println("----------测试商户注册----------")
	// 初始化参数
	body := MerchantRegisterRequest{
		BankName:                   "中国工商银行",
		BankAccount:                "621281240200093200",
		Drawer:                     "drawer",
		Payee:                      "payee",
		Reviewer:                   "reviewer",
		State:                      "1",
		TaxpayerNum:                "9876543210123456780",
		TaxpayerName:               "罗定市泗纶镇美森竹制品厂",
		LegalPersonName:            "方志欢",
		ContactsName:               "方志",
		Email:                      "eline.Liang@gaopeng.com",
		Phone:                      "18285162583",
		RegionCode:                 RegionCodeGuangDong,
		CityName:                   "深圳市",
		Address:                    "泗纶镇美森竹制品厂",
		CallbackUrl:                "http://www.baidu.com",
		TaxRegistrationCertificate: "营业执照base64编码文件",
	}
	// 请求接口
	rsp, err := testClient.MerchantRegister(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", rsp)
	return
}
