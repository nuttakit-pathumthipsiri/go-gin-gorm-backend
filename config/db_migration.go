package config

import (
	"go-gin-gorm-backend/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// Auto migrate the basic structure
	if err := db.AutoMigrate(&model.Topic{}, &model.TopicDetail{}); err != nil {
		return err
	}

	return nil
}
