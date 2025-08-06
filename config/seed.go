package config

import (
	"go-gin-gorm-backend/model"
	"log"

	"gorm.io/gorm"
)

func SeedTopics(db *gorm.DB) {
	topics := []model.Topic{
		{Name: "ยา", Order: 1, CreatedBy: "system"},
		{Name: "วิตามิน", Order: 2, CreatedBy: "system"},
		{Name: "จุลินทรีย์", Order: 3, CreatedBy: "system"},
		{Name: "ยี่ห้อ", Order: 4, CreatedBy: "system"},
	}
	for _, t := range topics {
		var count int64
		db.Model(&model.Topic{}).Where("name = ?", t.Name).Count(&count)
		if count == 0 {
			db.Create(&t)
		}
	}
}

func SeedTopicDetails(db *gorm.DB) {
	// Get the "ยา" topic
	var medicineTopic model.Topic
	if err := db.Where("name = ?", "ยา").First(&medicineTopic).Error; err != nil {
		log.Printf("Could not find 'ยา' topic: %v", err)
		return
	}

	// Topic details for "ยา"
	details := []model.TopicDetail{
		{Name: "ยาแก้ปวด", Order: 1, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้ไข้", Order: 2, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้ไอ", Order: 3, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้ท้องเสีย", Order: 4, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้ท้องผูก", Order: 5, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้แพ้", Order: 6, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยานอนหลับ", Order: 7, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาคลายกล้ามเนื้อ", Order: 8, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาแก้อักเสบ", Order: 9, TopicID: medicineTopic.ID, CreatedBy: "system"},
		{Name: "ยาฆ่าเชื้อ", Order: 10, TopicID: medicineTopic.ID, CreatedBy: "system"},
	}

	for _, d := range details {
		var count int64
		db.Model(&model.TopicDetail{}).Where("name = ? AND topic_id = ?", d.Name, d.TopicID).Count(&count)
		if count == 0 {
			db.Create(&d)
		}
	}
} 