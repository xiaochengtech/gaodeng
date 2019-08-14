package gaodeng

import (
	"encoding/json"
	"errors"
	"fmt"
	"gaodeng/goland"
)

// 高灯客户端的构造函数
func NewClient(isProd bool) (client *Client) {
	client = new(Client)
	if isProd {
		client.env = EnvProd
	} else {
		client.env = EnvTest
	}
	client.config = NewConfig(isProd)
	client.sdk = goland.NewSdk(client.config.AppKey, client.config.AppSecret, AppVersion, client.env)
	return client
}

type Client struct {
	env    string      // 是否是生产环境
	config Config      // 配置信息
	sdk    *goland.Sdk // 高灯开票sdk
}

// 向高灯sdk发送请求
func (c *Client) post(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body := make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)

	// 发起请求
	byteBody, err := c.sdk.HttpPost(relativeUrl, body)
	if err != nil {
		return
	}
	//raw := string(byteBody)
	//fmt.Printf("raw body : %s\n", raw)

	// 读取业务内容并转为字节流返回
	var model ResponseModel
	err = json.Unmarshal(byteBody, &model)
	if err != nil {
		fmt.Printf("json.Unmarshal error : %v\n", err)
		return
	}

	// 解析业务对象
	if model.Code == StatusCodeNormal {
		bytes, _ = json.Marshal(model.Data)
	} else {
		fmt.Printf("error code : %d, error message : %s\n", model.Code, model.Msg)
		err = errors.New(model.Msg)
	}
	return
}
