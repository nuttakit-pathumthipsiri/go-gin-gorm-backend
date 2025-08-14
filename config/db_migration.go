package config

import (
	"go-gin-gorm-backend/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// Drop existing tables if they exist (for SQL Server compatibility)
	db.Migrator().DropTable(&model.TopicDetail{})
	db.Migrator().DropTable(&model.Topic{})
	db.Migrator().DropTable(&model.User{})

	// Auto migrate the basic structure
	if err := db.AutoMigrate(&model.User{}, &model.Topic{}, &model.TopicDetail{}); err != nil {
		return err
	}

	return nil
}
