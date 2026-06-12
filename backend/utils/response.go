package utils

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 业务状态码常量
const (
	Unauthorized = 1001 // 未授权
	Forbidden    = 1002 // 禁止访问
	SUCCESS      = 2000 // 成功
	ERROR        = 2001 // 失败
	ServerError  = 2002 // 服务器错误
)
