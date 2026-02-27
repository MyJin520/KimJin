package db

import (
	"KimJin/src/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DB 全局数据库连接实例
var DB *gorm.DB

// InitSQLite 初始化SQLite连接
func InitSQLite() error {
	// 自定义日志配置（根据环境控制SQL日志）
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
		logger.Config{
			SlowThreshold: time.Second,   // 慢查询阈值
			LogLevel:      logger.Silent, // 默认静音
			Colorful:      true,          // 彩色输出
		},
	)

	// 开发环境开启SQL日志
	if config.GlobalConfig.Database.Debug {
		newLogger.LogMode(logger.Info)
	}

	// 连接SQLite
	db, err := gorm.Open(sqlite.Open(config.GlobalConfig.Database.Path), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	// 赋值给全局变量
	DB = db

	return nil
}
