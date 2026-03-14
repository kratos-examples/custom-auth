package models

import (
	"time"

	"gorm.io/gorm"
)

// Token model stores auth tokens
// Token 模型存储认证令牌
type Token struct {
	gorm.Model
	AdminID   uint      `gorm:"index" json:"admin_id"`         // Related admin ID // 关联的管理员ID
	Token     string    `gorm:"uniqueIndex;size:128" json:"-"` // Token value // 令牌值
	Type      string    `gorm:"size:32" json:"type"`           // Token type: access, refresh // 令牌类型：access、refresh
	ExpiresAt time.Time `json:"expires_at"`                    // Expire time // 过期时间
}

// TableName sets custom table name
// TableName 设置自定义表名
func (*Token) TableName() string {
	return "tb_token"
}
