package service

import (
	"chat-system-backend/internal/model"
	"chat-system-backend/internal/repository"
	"chat-system-backend/pkg/utils"
	"errors"
	"time"
)

// PasswordResetService 密码重置服务
type PasswordResetService interface {
	// 发送密码重置邮件/短信
	SendPasswordReset(email string) error
	// 验证密码重置令牌
	ValidateToken(token string) (*model.PasswordResetToken, error)
	// 重置用户密码
	ResetPassword(token string, newPassword string) error
}

// passwordResetService 是 PasswordResetService 的实例
type passwordResetService struct {
	userRepo               repository.UserRepository
	passwordResetTokenRepo repository.PasswordResetTokenRepository
	userService            UserService
}

// NewPasswordResetService 创建密码重置服务实例
func NewPasswordResetService(
	userRepo repository.UserRepository,
	passwordResetTokenRepo repository.PasswordResetTokenRepository,
	userService UserService,
) PasswordResetService {
	return &passwordResetService{
		userRepo:               userRepo,
		passwordResetTokenRepo: passwordResetTokenRepo,
		userService:            userService,
	}
}

// SendPasswordReset 实现发送密码重置邮件/短信方法
func (s *passwordResetService) SendPasswordReset(email string) error {
	// 1.根据邮箱查找用户
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	// 2.生成密码重置令牌（32 位随机字符串）
	tokenStr, err := utils.GenerateRandomToken(32)
	if err != nil {
		return err
	}

	// 3.创建密码重置令牌记录，有效期 15 分钟
	resetToken := &model.PasswordResetToken{
		UserID:    int64(user.ID),
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		IsUsed:    false,
	}

	// 4.保存令牌到数据库
	if err := s.passwordResetTokenRepo.Create(resetToken); err != nil {
		return err
	}

	// 5. 这里应该发送邮件或短信，包含重置链接
	// 由于没有邮件服务，这里仅打印日志
	// 实际项目中，应该使用邮件服务发送包含重置链接的邮件

	return nil
}

// ValidateToken 实现验证密码重置令牌方法
func (s *passwordResetService) ValidateToken(token string) (*model.PasswordResetToken, error) {
	// 1.根据令牌查找记录
	resetToken, err := s.passwordResetTokenRepo.FindByToken(token)
	if err != nil {
		return nil, errors.New("无效的重置令牌")
	}

	// 2.检查令牌是否已使用
	if resetToken.IsUsed {
		return nil, errors.New("该重置令牌已被使用")
	}

	// 3.检查令牌是否过期
	if time.Now().After(resetToken.ExpiresAt) {
		return nil, errors.New("重置令牌已过期")
	}

	return resetToken, nil
}

// ResetPassword 实现重置用户密码方法
func (s *passwordResetService) ResetPassword(token string, newPassword string) error {
	// 1.验证令牌
	resetToken, err := s.ValidateToken(token)
	if err != nil {
		return err
	}

	// 2.查找用户
	user, err := s.userRepo.GetByID(uint(resetToken.UserID))
	if err != nil {
		return err
	}

	// 3.更新用户密码
	err = s.userService.UpdatePassword(user.ID, newPassword)
	if err != nil {
		return err
	}

	// 4.标记令牌为已使用
	return s.passwordResetTokenRepo.MarkAsUsed(resetToken)
}
