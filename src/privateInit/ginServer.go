package privateInit

import (
	"KimJin/src/config"
	"KimJin/src/internal/api"
	"KimJin/src/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinRun(addr string, port int) {
	if addr == "" {
		addr = "0.0.0.0"
	}
	if port <= 0 {
		port = 8888
	}
	// 设置Gin运行模式
	if config.GlobalConfig.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎，注册中间件
	r := gin.Default()
	r.Use(middleware.Cors())     // 跨域中间件
	r.Use(middleware.Recovery()) // 异常恢复中间件

	// 注册控制器和路由
	formCtrl := api.NewFormController()
	apiGroup := r.Group("/api")
	{
		formGroup := apiGroup.Group("/form")
		{
			formGroup.GET("/config/:formId", formCtrl.GetFormConfig)
			formGroup.POST("/submit", formCtrl.SubmitForm)
			// 新增查询路由
			formGroup.GET("/submissions/:formId", formCtrl.GetFormSubmissions)
			formGroup.GET("/submission/:id", formCtrl.GetFormSubmissionByID)
		}
	}

	fmt.Printf("服务启动成功，监听端口：%d\n", port)
	if err := r.Run(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}
