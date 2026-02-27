package privateInit

import (
	"KimJin/src/internal/model"
	"KimJin/src/internal/service"
	"KimJin/src/pkg/db"
	"fmt"
)

func DBInit() {
	var err error
	// 初始化数据库
	if err = db.InitSQLite(); err != nil {
		panic(fmt.Sprintf("数据库初始化失败: %v", err))
	}

	// 自动迁移表结构
	if err = AutoMigrate(); err != nil {
		panic(fmt.Sprintf("表结构迁移失败: %v", err))
	}

	// 初始化默认表单配置 --> 后续删除
	formService := service.NewFormService()
	if err := formService.InitDefaultConfig(); err != nil {
		panic(fmt.Sprintf("初始化默认配置失败: %v", err))
	}
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	return db.DB.AutoMigrate(
		&model.FormConfig{},
		&model.FormSubmission{},
	)
}
