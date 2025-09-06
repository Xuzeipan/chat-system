package api

import (
	"chat-system-backend/internal/middleware"
	"chat-system-backend/internal/repository"

	"chat-system-backend/internal/handler"
	"chat-system-backend/internal/service"
	"chat-system-backend/pkg/db"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
// func SetupRouter(db *gorm.DB) *gin.Engine {
func SetupRouter() *gin.Engine {
	// 创建路由器
	r := gin.Default()

	// 全局中间件
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// 初始化仓库和服务
	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 初始化密码重置相关仓库、服务和处理器
	passwordResetTokenRepo := repository.NewPasswordResetTokenRepository(db.DB)
	passwordResetService := service.NewPasswordResetService(
		userRepo,
		passwordResetTokenRepo,
		userService,
	)
	passwordResetHandler := handler.NewPasswordResetHandler(passwordResetService)

	// API 路由 （不需要认证）
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// 用户相关路由
			users := v1.Group("/user")
			{
				users.POST("/register", userHandler.Register)
				users.POST("/login", userHandler.Login)

				// 需要认证的路由
				auth := users.Group("", middleware.JWTAuth())
				{
					auth.GET("/info", userHandler.GetUserInfo)
				}
			}

			// 密码重置相关路由
			passwordReset := v1.Group("/password-reset")
			{
				// 发送密码重置邮件
				passwordReset.POST("/send", passwordResetHandler.SendPasswordReset)
				// 验证密码重置令牌
				passwordReset.GET("/validate", passwordResetHandler.ValidateToken)
				// 重置密码
				passwordReset.POST("/reset", passwordResetHandler.ResetPassword)
			}

			// 测试路由
			v1.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}
	}

	return r
}
