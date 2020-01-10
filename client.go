package gaodeng

import (
	"encoding/json"
	"errors"
	"fmt"
	golden "gitee.com/cuckoopark/gaodeng/golden/1.0.0"
)

// 高灯客户端的构造函数
func NewClient(env string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.sdk = golden.NewSdk(config.AppKey, config.AppSecret, AppVersion, env)
	return client
}

type Client struct {
	config Config      // 配置信息
	sdk    *golden.Sdk // 高灯开票SDK
}

// 向高灯SDK发送请求
func (c *Client) post(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, err := json.Marshal(bodyObj)
	if err != nil {
		return
	}
	body := make(map[string]interface{})
	if err = json.Unmarshal(bodyJson, &body); err != nil {
		return
	}
	// 发起请求
	byteBody, err := c.sdk.HttpPost(relativeUrl, body)
	if err != nil {
		return
	}
	// 读取业务内容并转为字节流返回
	var model ResponseModel
	if err = json.Unmarshal(byteBody, &model); err != nil {
		return
	}
	// 解析业务对象
	if model.Code == StatusCodeNormal {
		bytes, err = json.Marshal(model.Data)
	} else {
		err = errors.New(fmt.Sprintf("%d:%s", model.Code, model.Msg))
	}
	return
}
