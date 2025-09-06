package middleware

import (
	"chat-system-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取 token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 token"})
			c.Abort()
			return
		}

		// 移除 Bearer 前缀
		const bearerPrefix = "Bearer "
		if len(tokenString) > len(bearerPrefix) && tokenString[:len(bearerPrefix)] == bearerPrefix {
			tokenString = tokenString[len(bearerPrefix):]
		}

		// 解析和验证token
		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			if err == utils.ErrExpiredToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token 已过期"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效"})
			}
			c.Abort()
			return
		}

		// 将用户 ID 存储在上下文
		c.Set("userID", userID)
		c.Next()
	}
}
