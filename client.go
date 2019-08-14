package gaodeng

import (
	"encoding/json"
	"errors"
	"fmt"
	"gaodeng/golden"
	"io/ioutil"
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
	client.sdk = golden.NewSdk(client.config.AppKey, client.config.AppSecret, client.env)
	return client
}

type Client struct {
	env    string      // 是否是生产环境
	config Config      // 配置信息
	sdk    *golden.Sdk // 高灯开票sdk
}

// 向高灯sdk发送请求
func (c *Client) post(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body := make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)

	// 发起请求
	resp, err := c.sdk.HttpPost(relativeUrl, body)
	if err != nil {
		return
	}

	// 读取返回body的内容
	defer resp.Body.Close()
	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//rst := string(byteBody)
	//fmt.Println(rst)

	// 读取业务内容并转为字节流返回
	var model ResponseModel
	_ = json.Unmarshal(byteBody, &model)
	if model.Code == StatusCodeNormal {
		bytes, _ = json.Marshal(model.Data)
	} else {
		fmt.Printf("error code : %d, error message : %s\n", model.Code, model.Msg)
		err = errors.New(model.Msg)
	}
	return
}
