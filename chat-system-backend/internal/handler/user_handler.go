package handler

import (
	"chat-system-backend/internal/handler/dto"
	"chat-system-backend/internal/model"
	"chat-system-backend/internal/service"
	"chat-system-backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建新的用户处理器
func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{userService: svc}
}

// 将用户模型转换为响应 DTO
func convertUserToResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Avatar:        user.Avatar,
		Bio:           user.Bio,
		Phone:         user.Phone,
		Email:         user.Email,
		PhoneVerified: user.PhoneVerified,
		EmailVerified: user.EmailVerified,
		Status:        user.Status,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层注册用户
	user, err := h.userService.Register(req.Username, req.Password)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user":    convertUserToResponse(user),
	})
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层登录
	user, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 更新登录信息
	ip := c.ClientIP()
	if err := h.userService.UpdateLoginInfo(user.ID, ip); err != nil {
		// 记录错误但不阻止登录
		log.Printf("Warning: Update login info failed: %v", err)
	}

	// 这里应该生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
		return
	}
	response := &dto.LoginResponse{
		Token: token,
		User:  convertUserToResponse(user),
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "data": response})
}

// GetUserInfo 获取用户信息
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// 从上下文中获取用户 ID（在中间件中设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的用户 ID"})
		return
	}

	user, err := h.userService.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "获取用户信息成功", "data": convertUserToResponse(user)})
}
