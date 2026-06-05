package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	dsn := cfg.Dsn

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	} else {
		fmt.Println("Database connected successfully")
	}

	return db
}
