package config

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN string
}

func LoadDBConfig() *DBConfig {
	dsn := os.Getenv("DB_DSN")
	return &DBConfig{DSN: dsn}
}

func ConnectDB(cfg *DBConfig) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, fmt.Errorf("database DSN is not set in environment variable DB_DSN")
	}
	db, err := gorm.Open(sqlserver.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
