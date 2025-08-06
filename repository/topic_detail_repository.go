package repository

import (
	"go-gin-gorm-backend/model"

	"gorm.io/gorm"
)

type TopicDetailRepository interface {
	Create(detail *model.TopicDetail) error
	FindAllByTopicID(topicID uint) ([]model.TopicDetail, error)
	FindByID(id uint) (*model.TopicDetail, error)
	FindByName(name string) (*model.TopicDetail, error)
	Update(detail *model.TopicDetail) error
	Delete(id uint) error
}

type topicDetailRepository struct {
	db *gorm.DB
}

func NewTopicDetailRepository(db *gorm.DB) TopicDetailRepository {
	return &topicDetailRepository{db}
}

func (r *topicDetailRepository) Create(detail *model.TopicDetail) error {
	return r.db.Create(detail).Error
}

func (r *topicDetailRepository) FindAllByTopicID(topicID uint) ([]model.TopicDetail, error) {
	var details []model.TopicDetail
	err := r.db.Where("topic_id = ?", topicID).Order("[order] ASC").Find(&details).Error
	return details, err
}

func (r *topicDetailRepository) FindByID(id uint) (*model.TopicDetail, error) {
	var detail model.TopicDetail
	err := r.db.First(&detail, "id = ?", id).Error
	return &detail, err
}

func (r *topicDetailRepository) FindByName(name string) (*model.TopicDetail, error) {
	var detail model.TopicDetail
	err := r.db.First(&detail, "name = ?", name).Error
	return &detail, err
}

func (r *topicDetailRepository) Update(detail *model.TopicDetail) error {
	return r.db.Save(detail).Error
}

func (r *topicDetailRepository) Delete(id uint) error {
	return r.db.Delete(&model.TopicDetail{}, "id = ?", id).Error
} 