package gaodeng

import "os"

// 整体配置的构造函数
func NewConfig() (conf Config) {
	conf = Config{
		AppKey:    os.Getenv("AppKey"),
		AppSecret: os.Getenv("AppSecret"),
	}
	return conf
}

// 高灯发票的整体配置
type Config struct {
	AppKey    string // 高灯开放平台id
	AppSecret string // 高灯开放平台secret
}
