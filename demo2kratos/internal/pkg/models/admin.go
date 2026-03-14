package models

import "gorm.io/gorm"

// Admin model stores admin user info
// Admin 模型存储管理员用户信息
type Admin struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:64" json:"username"` // Unique username // 唯一用户名
	Password string `gorm:"size:128" json:"-"`                   // Encrypted password // 加密密码
	Mailbox  string `gorm:"size:128" json:"mailbox"`             // Email address // 邮箱地址
	Status   int    `gorm:"default:1" json:"status"`             // Status: 1=active, 0=disabled // 状态：1=启用，0=禁用
}

// TableName sets custom table name
// TableName 设置自定义表名
func (*Admin) TableName() string {
	return "tb_admin"
}
