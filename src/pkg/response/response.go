package response

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	EXPIRE_CODE = 401
	ERROR       = 7
	SUCCESS     = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func DownloadFile(filepath, filename string, c *gin.Context) {
	c.FileAttachment(filepath, filename)
}

func FailWithMessage(c *gin.Context, message string) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func OkWithMessage(c *gin.Context, message string) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(SUCCESS, data, message, c)
}

// JWT使用
func NoAuth(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		EXPIRE_CODE,
		nil,
		message,
	})
}
func FailWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(ERROR, data, message, c)
}

func OkSendByteDate(rawFileName string, data []byte, c *gin.Context) {
	encodedFileName := url.QueryEscape(rawFileName)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"; filename*=UTF-8''%s", encodedFileName, encodedFileName))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	_, err := c.Writer.Write(data)
	if err != nil {
		FailWithMessage(c, "文件下载响应失败")
	}
}
