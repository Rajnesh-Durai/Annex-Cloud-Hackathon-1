package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, _ := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}