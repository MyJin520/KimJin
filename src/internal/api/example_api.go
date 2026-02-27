package api

import (
	"KimJin/src/internal/service"
	"KimJin/src/pkg/logger"
	"KimJin/src/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// FormController 表单控制器
type FormController struct {
	formService *service.FormService
}

// NewFormController 创建表单控制器实例
func NewFormController() *FormController {
	return &FormController{
		formService: service.NewFormService(),
	}
}

// GetFormConfig 获取表单配置
func (c *FormController) GetFormConfig(ctx *gin.Context) {
	formID := ctx.Param("formId")
	config, err := c.formService.GetFormConfig(formID)
	if err != nil {
		logger.Error("获取表单配置失败", zap.Error(err))
		response.FailWithMessage(ctx, "表单配置不存在")
		return
	}
	response.OkWithDetailed(ctx, "表单配置获取成功", config)
}

// SubmitForm 提交表单数据
func (c *FormController) SubmitForm(ctx *gin.Context) {
	// 定义接收参数的结构体
	type SubmitRequest struct {
		FormID string `json:"form_id" binding:"required"`
		Data   string `json:"data" binding:"required"`
	}

	var req SubmitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(ctx, "参数错误："+err.Error())
		return
	}

	// 调用服务层提交数据
	submissionID, err := c.formService.SubmitForm(req.FormID, req.Data)
	if err != nil {
		response.FailWithMessage(ctx, "提交失败："+err.Error())
		return
	}

	response.OkWithDetailed(ctx, "提交成功", gin.H{"submission_id": submissionID})
}

// GetFormSubmissions 查询表单提交记录
func (c *FormController) GetFormSubmissions(ctx *gin.Context) {
	formID := ctx.Param("formId")
	submissions, err := c.formService.GetFormSubmissions(formID)
	if err != nil {
		response.FailWithMessage(ctx, "查询失败："+err.Error())
		return
	}
	response.OkWithDetailed(ctx, "查询成功", submissions)
}

// GetFormSubmissionByID 查询单条提交记录
func (c *FormController) GetFormSubmissionByID(ctx *gin.Context) {
	// 解析ID参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage(ctx, "ID格式错误")
		return
	}

	submission, err := c.formService.GetFormSubmissionByID(uint(id))
	if err != nil {
		response.FailWithMessage(ctx, "提交记录不存在")
		return
	}
	response.OkWithDetailed(ctx, "查询成功", submission)
}
