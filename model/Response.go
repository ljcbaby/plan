package model

// 响应体
type Response struct {
	Code int64       `json:"code"` // 返回码
	Data interface{} `json:"data"` // 返回数据
	Msg  string      `json:"msg"`  // 信息
}
