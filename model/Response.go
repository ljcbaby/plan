package model

// 响应体
type Response struct {
	Code int         `json:"code"`           // 返回码
	Data interface{} `json:"data,omitempty"` // 返回数据
	Msg  string      `json:"msg"`            // 信息
}

type Page struct {
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
