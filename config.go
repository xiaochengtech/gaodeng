package gaodeng

import "os"

// 整体配置的构造函数
func NewConfig(isProd bool) (conf Config) {
	if isProd {
		conf = Config{
			AppKey:    os.Getenv("AppKey"),
			AppSecret: os.Getenv("AppSecret"),
		}
	} else {
		conf = Config{
			AppKey:    TestAppKey,
			AppSecret: TestAppSecret,
		}
	}
	return conf
}

// 高灯发票的整体配置
type Config struct {
	AppKey    string // 高灯开放平台id
	AppSecret string // 高灯开放平台secret
}
