package gaodeng

var testClient = NewClient(EnvTest, Config{
	AppKey:    TestAppKey,
	AppSecret: TestAppSecret,
})

const (
	TestAppKey         = "EgDjckWzyGxwIi7e9J1A8LdruWMidFFH"
	TestAppSecret      = "9Q8744Oe0nv8aw738b3HkjdylYZzNeZOcTz53KI4pchKpqIi"
	TestTaxPayerNumber = "91440300661005378A"
	TestSellerName     = "深圳市卓越物业管理股份有限公司卓越时代广场地下停车场"
	TestSellerAddress  = "深圳市福田区益田路与福华路交汇处卓越时代广场地下"

	GoodsName             = "计时停车"
	TaxCodeParking        = "3040502020200000000"
	TaxRateParking uint32 = 90
)
