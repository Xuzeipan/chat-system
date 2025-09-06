package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// GenerateRandomToken 生成指定长度的安全随机令牌
func GenerateRandomToken(length int) (string, error) {
	// 计算需要生成的随机字节数
	// base64编码会使数据长度增加约 1/3，所以需要生成更少的字节
	bytes := make([]byte, length*3/4)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}
	// 使用 URL 安全的 base64 编码，去掉尾部的填充字符
	token := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
	// 确保返回指定长度的令牌
	if len(token) > length {
		token = token[:length]
	}
	return token, nil
}
