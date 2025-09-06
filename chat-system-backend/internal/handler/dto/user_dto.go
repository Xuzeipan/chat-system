package dto

import (
	"time"
)

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username" binding: "required,min=3,max=50"`
	Password string `json:"password" binding: "required,min=6,max=30"`
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserResponse 用户响应
type UserResponse struct {
	ID            uint      `json:"id"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	Bio           string    `json:"bio"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	PhoneVerified bool      `json:"phone_verified"`
	EmailVerified bool      `json:"email_verified"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}
