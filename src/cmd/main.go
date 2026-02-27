package main

import (
	"KimJin/src/config"
	"KimJin/src/pkg/logger"
	"KimJin/src/privateInit"
)

func main() {
	// 初始化日志
	logger.InitZapLogger("./src/loggs")
	defer logger.Sync()
	// 初始化配置文件
	config.Init()
	// 初始化数据库
	privateInit.DBInit()
	// 启动gin服务
	privateInit.GinRun(config.GlobalConfig.App.Addr, config.GlobalConfig.App.Port)
}
