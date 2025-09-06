package model

import (
    "time"
)

// PasswordResetToken 表示密码重置令牌模型
type PasswordResetToken struct {
    ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID    int64     `gorm:"not null;index:idx_password_reset_tokens_user_id" json:"user_id"`
    Token     string    `gorm:"not null;uniqueIndex:idx_password_reset_tokens_token" json:"token"`
    ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
    IsUsed    bool      `gorm:"not null;default:false" json:"is_used"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    
    // 关联用户
    User      User      `gorm:"foreignKey:UserID" json:"user"`
}

// 自定义表名
func (PasswordResetToken) TableName() string {
    return "password_reset_tokens"
}
