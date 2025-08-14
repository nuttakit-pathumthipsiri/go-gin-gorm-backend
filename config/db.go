package config

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN      string
	Database string
	Server   string
	Port     string
	User     string
	Password string
}

func LoadDBConfig() *DBConfig {
	dsn := os.Getenv("DB_DSN")
	database := os.Getenv("DB_NAME")
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if database == "" {
		database = "go_gin_gorm_db"
	}
	if server == "" {
		server = "localhost"
	}
	if port == "" {
		port = "1433"
	}
	if user == "" {
		user = "sa"
	}
	if password == "" {
		password = "StrongP@ssw0rd"
	}

	return &DBConfig{
		DSN:      dsn,
		Database: database,
		Server:   server,
		Port:     port,
		User:     user,
		Password: password,
	}
}

func ConnectDB(cfg *DBConfig) (*gorm.DB, error) {
	// First connect to master database to create our database
	masterDSN := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=master",
		cfg.User, cfg.Password, cfg.Server, cfg.Port)

	masterDB, err := gorm.Open(sqlserver.Open(masterDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to master database: %v", err)
	}

	// Create database if it doesn't exist
	createDBQuery := fmt.Sprintf("IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = '%s') CREATE DATABASE [%s]",
		cfg.Database, cfg.Database)

	if err := masterDB.Exec(createDBQuery).Error; err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	// Close master connection
	sqlDB, err := masterDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %v", err)
	}
	sqlDB.Close()

	// Now connect to our specific database
	var targetDSN string
	if cfg.DSN != "" {
		targetDSN = cfg.DSN
	} else {
		targetDSN = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
	}

	db, err := gorm.Open(sqlserver.Open(targetDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target database: %v", err)
	}

	return db, nil
}
