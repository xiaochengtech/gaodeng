package gaodeng

import "gaodeng/golden"

// 高灯客户端的构造函数
func NewClient(isProd bool) (client *Client) {
	client = new(Client)
	if isProd {
		client.env = EnvProd
	} else {
		client.env = EnvTest
	}
	client.config = NewConfig()
	client.sdk = golden.NewSdk(client.config.AppKey, client.config.AppSecret, client.env)
	return client
}

type Client struct {
	env    string      // 是否是生产环境
	config Config      // 配置信息
	sdk    *golden.Sdk // 高灯开票sdk
}
