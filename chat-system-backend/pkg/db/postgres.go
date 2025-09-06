package db

import (
	"fmt"
	"log"
	"time"

	"chat-system-backend/config"
	"chat-system-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitPostgreSQL 初始化PostgreSQL连接
func InitPostgreSQL() error {
	cfg := config.GetConfig().Database

	// 构建连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	// 配置GORM日志级别
	logLevel := logger.Info
	if config.GetConfig().App.Env == "production" {
		logLevel = logger.Warn
	}

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// 配置连接池参数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	DB = db

	log.Println("Successfully connected to PostgreSQL database")

	// 自动迁移表结构（开发环境下）
	if config.GetConfig().App.Env == "development" {
		// 执行自动迁移
		if err := db.AutoMigrate(&model.User{}); err != nil {
			log.Printf("Warning: Auto migrate failed: %v", err)
		} else {
			log.Printf("Database auto migrate completed")
		}
	}

	return nil
}
