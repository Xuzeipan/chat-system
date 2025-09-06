package utils

import (
	"chat-system-backend/config"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

// GenerateToken 生成 JWT token
func GenerateToken(userID uint, username string) (string, error) {
	jwtConfig := config.GetConfig().JWT

	// 创建 claims
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Duration(jwtConfig.ExpireHours) * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	// 创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的编码后的字符串 token
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ParseToken 解析 JWT token
func ParseToken(tokenString string) (uint, error) {
	jwtConfig := config.GetConfig().JWT

	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		// 检查是否是token 过期错误
		if strings.Contains(err.Error(), "token is expired") {
			return 0, ErrExpiredToken
		}
		return 0, ErrInvalidToken
	}

	// 验证 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 从 claims 中获取用户 ID
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, ErrInvalidToken
		}
		return uint(userIDFloat), nil
	}

	return 0, ErrInvalidToken
}
