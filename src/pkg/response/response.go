package response

import "github.com/gin-gonic/gin"

// Response 统一响应结构体
type Response struct {
	Code int         `json:"code"` // 状态码：200成功，其他失败
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 业务数据
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "操作成功",
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{ // HTTP状态码统一200，业务码区分
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// NotFound 资源不存在响应
func NotFound(c *gin.Context, msg string) {
	Fail(c, 404, msg)
}

// BadRequest 参数错误响应
func BadRequest(c *gin.Context, msg string) {
	Fail(c, 400, msg)
}

// ServerError 服务器内部错误响应
func ServerError(c *gin.Context, msg string) {
	Fail(c, 500, msg)
}
