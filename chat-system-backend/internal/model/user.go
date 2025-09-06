package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
// 适配PostgreSQL数据库
// 目前使用账号密码登录，预留手机号和邮箱字段用于后期扩展
type User struct {
	// gorm.Model 包含了ID, CreatedAt, UpdatedAt, DeletedAt字段
	gorm.Model
	// 基本信息
	Username string `gorm:"size:50;not null;uniqueIndex:idx_username" json:"username"` // 用户名，唯一索引
	Nickname string `gorm:"size:50;default:''" json:"nickname"`                        // 昵称
	Avatar   string `gorm:"size:255;default:''" json:"avatar"`                         // 头像URL
	Bio      string `gorm:"type:text;default:''" json:"bio"`                           // 个人简介

	// 认证相关
	Password string `gorm:"size:255;not null" json:"-"` // 密码哈希，JSON序列化时忽略
	Salt     string `gorm:"size:50;not null" json:"-"`  // 密码盐值，JSON序列化时忽略

	// 预留字段 - 后期扩展手机号和邮箱登录
	Email         string `gorm:"size:100;index:idx_email,unique,class:btree,optional" json:"email"`
	Phone         string `gorm:"size:20;index:idx_phone,unique,class:btree,optional" json:"phone"`
	PhoneVerified bool   `gorm:"default:false" json:"phone_verified"` // 手机号是否已验证
	EmailVerified bool   `gorm:"default:false" json:"email_verified"` // 邮箱是否已验证

	// 状态字段
	Status int `gorm:"default:1" json:"status"` // 状态：1-正常，2-禁用

	// 扩展字段
	LastLoginAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_login_at"` // 最后登录时间
	LastLoginIP string    `gorm:"size:50;default:''" json:"last_login_ip"`        // 最后登录IP
}

// 自定义表名
func (User) TableName() string {
	return "users"
}
