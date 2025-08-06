package model

import (
	"time"
)

// TopicDetail represents a topic detail entity
// @Description Topic detail entity
type TopicDetail struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
	TopicID   uint      `gorm:"not null;index" json:"topic_id" example:"1"`                                      // รหัส topic
	Name      string    `gorm:"size:255;not null;uniqueIndex" json:"name" example:"ยาแก้ปวด" binding:"required"` // ชื่อ topic_detail
	Order     int       `gorm:"not null" json:"order" example:"1" binding:"required"`                            // ลำดับ topic_detail
	CreatedBy string    `gorm:"size:100;not null" json:"created_by" example:"admin" binding:"required"`          // ผู้สร้าง
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`       // วันที่สร้าง
	UpdatedBy string    `gorm:"size:100" json:"updated_by" example:"admin"`                                      // ผู้อัพเดท
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`       // วันที่อัพเดท
	Topic     Topic     `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
}

// TopicDetailRequest represents a topic detail request (without auto-generated fields)
// @Description Topic detail request object
type CreateTopicDetailRequest struct {
	Name string `json:"name" example:"ยาแก้ปวด" binding:"required"` // ชื่อ topic_detail
}

// UpdateTopicDetailRequest represents an update topic detail request (with optional fields)
// @Description Update topic detail request object
type UpdateTopicDetailRequest struct {
	Name  *string `json:"name,omitempty" example:"ยาแก้ปวด"` // ชื่อ topic_detail (optional)
	Order *int    `json:"order,omitempty" example:"1"`       // ลำดับ topic_detail (optional)
}
