package service

import (
	"chat-system-backend/internal/model"
	"chat-system-backend/internal/repository"
	"crypto/rand"
	"encoding/hex"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	// 用户注册
	Register(username, password string) (*model.User, error)
	// 用户登录
	Login(username, password string) (*model.User, error)
	// 获取用户信息
	GetUserInfo(id uint) (*model.User, error)
	// 更新用户信息
	UpdateUserInfo(*model.User) error
	// 更新用户登录信息
	UpdateLoginInfo(userID uint, ip string) error
	// 更新用户密码
	UpdatePassword(userID uint, newPassword string) error
}

// userService 实现 UserService 接口
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建新的用户服务实例
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

// 生成随机盐值
func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// 哈希密码
func hashPassword(password, salt string) (string, error) {
	// 将盐值与密码组合后进行哈希
	combined := []byte(salt + password)
	hash, err := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Register 用户注册
func (s *userService) Register(username, password string) (*model.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.GetByUsername(username)
	if err == nil && existingUser != nil {
		return nil, gorm.ErrDuplicatedKey
	}

	// 生成盐值
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}

	// 哈希密码
	passwordHash, err := hashPassword(password, salt)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		Username:    username,
		Password:    passwordHash,
		Salt:        salt,
		Nickname:    "",
		Avatar:      "",
		Bio:         "",
		Status:      1,
		LastLoginAt: time.Now(),
		LastLoginIP: "",
	}

	// 保存用户信息
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *userService) Login(username, password string) (*model.User, error) {
	// 获取用户信息
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	combined := []byte(user.Salt + password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), combined)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserInfo 获取用户信息
func (s *userService) GetUserInfo(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// UpdateUserInfo 更新用户信息
func (s *userService) UpdateUserInfo(user *model.User) error {
	return s.userRepo.Update(user)
}

// UpdateLoginInfo 更新用户登录信息
func (s *userService) UpdateLoginInfo(userID uint, ip string) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	user.LastLoginAt = time.Now()
	user.LastLoginIP = ip

	return s.userRepo.Update(user)
}

// UpdatePassword 更新用户密码
func (s *userService) UpdatePassword(userID uint, newPassword string) error {
	// 获取用户信息
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// 生成新的盐值
	newSalt, err := generateSalt()
	if err != nil {
		return err
	}

	// 哈希新密码
	passwordHash, err := hashPassword(newPassword, newSalt)
	if err != nil {
		return err
	}

	// 更新用户密码和盐值
	user.Password = passwordHash
	user.Salt = newSalt

	return s.userRepo.Update(user)
}
