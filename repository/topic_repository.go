package repository

import (
	"go-gin-gorm-backend/model"

	"gorm.io/gorm"
)

type TopicRepository interface {
	Create(topic *model.Topic) error
	FindAll() ([]model.Topic, error)
	FindByID(id uint) (*model.Topic, error)
	FindByName(name string) (*model.Topic, error)
	Update(topic *model.Topic) error
	Delete(id uint) error
}

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{db}
}

func (r *topicRepository) Create(topic *model.Topic) error {
	return r.db.Create(topic).Error
}

func (r *topicRepository) FindAll() ([]model.Topic, error) {
	var topics []model.Topic
	err := r.db.Order("[order] ASC").Find(&topics).Error
	return topics, err
}

func (r *topicRepository) FindByID(id uint) (*model.Topic, error) {
	var topic model.Topic
	err := r.db.First(&topic, "id = ?", id).Error
	return &topic, err
}

func (r *topicRepository) FindByName(name string) (*model.Topic, error) {
	var topic model.Topic
	err := r.db.First(&topic, "name = ?", name).Error
	return &topic, err
}

func (r *topicRepository) Update(topic *model.Topic) error {
	return r.db.Save(topic).Error
}

func (r *topicRepository) Delete(id uint) error {
	return r.db.Delete(&model.Topic{}, "id = ?", id).Error
}
