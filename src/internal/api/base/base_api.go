package base

import (
	"KimJin/src/internal/service/base"
	"KimJin/src/pkg/request"
	"KimJin/src/pkg/response"
	"github.com/gin-gonic/gin"
)

// PublicAPI 公共接口
type PublicAPI struct {
	publicService *base.PublicRepoService
}

// NewPublicController 创建表单控制器实例
func NewPublicController() *PublicAPI {
	return &PublicAPI{
		publicService: base.NewPublicRepoService(),
	}
}

func (p *PublicAPI) Login(c *gin.Context) {
	// 绑定请求参数
	var logReq request.LoginRequest
	if err := c.ShouldBindJSON(&logReq); err != nil {
		response.FailWithMessage(c, "参数错误："+err.Error())
		return
	}
	// 调用服务层登录方法
	token, err := p.publicService.Login(logReq.Username, logReq.Password)
	if err != nil {
		response.FailWithMessage(c, "登录失败："+err.Error())
		return
	}
	response.OkWithDetailed(c, "登录成功", map[string]string{"token": token})
}
