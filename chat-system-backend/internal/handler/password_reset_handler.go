package handler

import (
	"chat-system-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PasswordResetHandler 处理密码重置相关请求
type PasswordResetHandler struct {
	passwordResetSrevice service.PasswordResetService
}

// NewPasswordResetHandler 创建密码重置处理器实例
func NewPasswordResetHandler(passwordResetService service.PasswordResetService) *PasswordResetHandler {
	return &PasswordResetHandler{
		passwordResetSrevice: passwordResetService,
	}
}

// sendPasswordReset 处理发送密码重置请求
func (h *PasswordResetHandler) SendPasswordReset(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding: "required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.passwordResetSrevice.SendPasswordReset(request.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送重置邮件失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "重置邮件已发送，请查收"})
}

// ValidateToken 验证密码重置令牌
func (h *PasswordResetHandler) ValidateToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "令牌不能为空"})
		return
	}

	resetToken, err := h.passwordResetSrevice.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"user_id": resetToken.UserID,
	})
}

// ResetPassword 处理密码重置请求
func (h *PasswordResetHandler) ResetPassword(c *gin.Context) {
	var request struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6,max=32"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.passwordResetSrevice.ResetPassword(request.Token, request.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "密码重置成功"})
}
