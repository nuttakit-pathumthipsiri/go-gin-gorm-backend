package model

import (
	"time"
)

// Topic represents a topic entity
// @Description Topic entity
type Topic struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
	Name      string    `gorm:"size:255;not null;uniqueIndex" json:"name" example:"ยา" binding:"required"` // ชื่อ topic
	Order     int       `gorm:"not null" json:"order" example:"1" binding:"required"`                      // ลำดับ topic
	CreatedBy string    `gorm:"size:100;not null" json:"created_by" example:"admin" binding:"required"`    // ผู้สร้าง
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"` // วันที่สร้าง
	UpdatedBy string    `gorm:"size:100" json:"updated_by" example:"admin"`                                // ผู้อัพเดท
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"` // วันที่อัพเดท
}

// TopicRequest represents a topic request (without auto-generated fields)
// @Description Topic request object
type CreateTopicRequest struct {
	Name string `json:"name" example:"ยา" binding:"required"` // ชื่อ topic
}

type UpdateTopicRequest struct {
	Name  *string `json:"name,omitempty" example:"ยา"` // ชื่อ topic (optional)
	Order *int    `json:"order,omitempty" example:"1"` // ลำดับ topic (optional)
}
