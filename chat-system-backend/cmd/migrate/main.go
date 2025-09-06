package main

import (
	"chat-system-backend/config"
	"chat-system-backend/pkg/db"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// 初始化配置
	config.InitConfig("./config/config.yaml")

	// 初始化数据库连接
	if err := db.InitPostgreSQL(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取原始数据库连接
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB instance: %v", err)
	}

	// 创建迁移驱动
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{SchemaName: "public"})
	if err != nil {
		log.Fatalf("Failed to create postgres driver: %v", err)
	}

	// 创建迁移源
	source, err := (&file.File{}).Open("file://migrations/migrations")
	if err != nil {
		log.Fatalf("Failed to open migration source: %v", err)
	}

	// 创建迁移实例
	m, err := migrate.NewWithInstance(
		"file",
		source,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// 执行迁移
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply")
		} else {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
