package main

import (
	"chat-system-backend/config"
	"chat-system-backend/internal/api"
	"chat-system-backend/pkg/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 初始化配置
	config.InitConfig("./config/config.yaml")

	// 初始化数据库连接
	if err := db.InitPostgreSQL(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 调用 SetupRouter 函数获取 Gin 引擎实例
	router := api.SetupRouter()

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(config.GetConfig().Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.GetConfig().Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.GetConfig().Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.GetConfig().Server.IdleTimeout) * time.Second,
	}

	// 启动服务器
	log.Println("Server is running on :8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

}
