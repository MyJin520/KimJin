package model

type User struct {
	BaseModel `gorm:"embedded"`
	Name      string
	Password  string
	Avatar    string `gorm:"comment:头像地址"`
	Email     string `gorm:"comment:邮箱地址"`
}
