package repository

import (
	"chat-system-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

// PasswordResetTokenRepository 定义密码重置令牌仓库接口
type PasswordResetTokenRepository interface {
	// 创建密码重置令牌
	Create(token *model.PasswordResetToken) error
	// 根据令牌查找
	FindByToken(token string) (*model.PasswordResetToken, error)
	// 标记令牌为已使用
	MarkAsUsed(token *model.PasswordResetToken) error
	// 删除过期的令牌
	DeleteExpired() error
}

// passwordResetTokenRepository 实现 PasswordResetTokenRepository 接口
type passwordResetTokenRepository struct {
	db *gorm.DB
}

// NewPasswordResetTokenRepository 创建密码重置令牌仓库实例
func NewPasswordResetTokenRepository(db *gorm.DB) PasswordResetTokenRepository {
	return &passwordResetTokenRepository{db: db}
}

// Create 实现创建密码重置令牌方法
func (r *passwordResetTokenRepository) Create(token *model.PasswordResetToken) error {
	return r.db.Create(token).Error
}

// FindByToken 实现根据令牌查找方法
func (r *passwordResetTokenRepository) FindByToken(token string) (*model.PasswordResetToken, error) {
	var resetToken model.PasswordResetToken
	err := r.db.Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

// MarkAsUsed 实现标记令牌为已使用方法
func (r *passwordResetTokenRepository) MarkAsUsed(token *model.PasswordResetToken) error {
	token.IsUsed = true
	token.UpdatedAt = time.Now()
	return r.db.Save(token).Error
}

// DeleteExpired 实现删除过去的令牌方法
func (r *passwordResetTokenRepository) DeleteExpired() error {
	return r.db.Where("expries_at < ?", time.Now()).Delete(&model.PasswordResetToken{}).Error
}
