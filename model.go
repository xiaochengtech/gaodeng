package gaodeng

// 返回结果的通信标识
type ResponseModel struct {
	Code uint16 `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 错误信息
	Data string `json:"data"` // 业务逻辑返回
}
