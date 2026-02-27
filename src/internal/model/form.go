package model

// FormConfig 动态表单配置模型
type FormConfig struct {
	BaseModel `gorm:"embedded"`
	FormID    string `gorm:"index;size:64" json:"form_id"` // 表单ID
	FormName  string `gorm:"size:128" json:"form_name"`    // 表单名称
	Fields    string `gorm:"type:text" json:"fields"`      // 字段配置（JSON字符串）
}

type FormSubmission struct {
	BaseModel `gorm:"embedded"`
	FormID    string `gorm:"index;size:64" json:"form_id"` // 关联表单ID
	Data      string `gorm:"type:text" json:"data"`        // 提交数据（JSON字符串）
}
