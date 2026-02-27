package router

import (
	"KimJin/src/internal/api"
	"github.com/gin-gonic/gin"
)

func FormRouter(r *gin.RouterGroup) {
	formGroup := r.Group("form")

	formGroup.GET("/config/:formId", api.FormAPI.GetFormConfig)
	formGroup.POST("/submit", api.FormAPI.SubmitForm)
	formGroup.GET("/submissions/:formId", api.FormAPI.GetFormSubmissions)
	formGroup.GET("/submission/:id", api.FormAPI.GetFormSubmissionByID)
}
