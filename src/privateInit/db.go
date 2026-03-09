package privateInit

import (
	"KimJin/src/config"
	"KimJin/src/internal/model"
	"KimJin/src/pkg/db"
	"KimJin/src/pkg/logger"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

	// 初始化默认用户
	InitDefaultUser()
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	return db.DB.AutoMigrate(
		&model.FormConfig{},
		&model.FormSubmission{},
		&model.User{},
	)
}

// InitDefaultUser 初始化默认用户
func InitDefaultUser() {
	// 检查默认用户是否存在
	var count int64
	db.DB.Model(&model.User{}).Where("name = ? and id = ?", config.GlobalConfig.App.DefaultUser, 1).Count(&count)
	if count > 0 {
		return
	}

	// 加密密码
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(config.GlobalConfig.App.DefaultPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("密码加密失败: %v", zap.Error(err))
		return
	}

	// 创建默认用户
	defaultUser := model.User{
		Name:     config.GlobalConfig.App.DefaultUser,
		Password: string(encryptedPassword),
	}
	if err := db.DB.Create(&defaultUser).Error; err != nil {
		logger.Error("创建默认用户失败: %v", zap.Error(err))
	}
}
