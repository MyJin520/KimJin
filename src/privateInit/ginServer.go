package privateInit

import (
	"KimJin/src/config"
	"KimJin/src/internal/middleware"
	"KimJin/src/internal/router"
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

	baseGroup := r.Group("/")
	{
		router.FormRouter(baseGroup)
	}

	// todo 后续新增登陆路由，登陆路由为私有路由
	//publicGroup := r.Group("/public")
	//{
	//	router.FormRouter(publicGroup)
	//}
	//
	//privateGroup := r.Group("/private")
	//{
	//	router.FormRouter(privateGroup)
	//}

	fmt.Printf("服务启动成功，监听端口：%d\n", port)
	if err := r.Run(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}
