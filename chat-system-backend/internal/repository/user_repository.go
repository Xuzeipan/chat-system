package repository

import (
	"chat-system-backend/internal/model"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	// 创建用户
	Create(user *model.User) error
	// 根据ID 获取用户
	GetByID(id uint) (*model.User, error)
	// 根据用户名获取用户
	GetByUsername(username string) (*model.User, error)
	// 根据手机号获取用户
	GetByPhone(phone string) (*model.User, error)
	// 根据邮箱获取用户
	GetByEmail(email string) (*model.User, error)
	// 更新用户信息
	Update(*model.User) error
	// 删除用户
	Delete(id uint) error
}

// userRepository 实现 UserRepository 接口
type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository 创建新的用户仓库实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// Create 创建用户
func (r *userRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// GetByID 根据ID 获取用户
func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetByPhone 根据手机号获取用户
func (r *userRepository) GetByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update 更新用户信息
func (r *userRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

// Delete 删除用户
func (r *userRepository) Delete(id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}
