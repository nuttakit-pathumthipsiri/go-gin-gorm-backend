package service

import (
	"errors"
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/repository"
	"go-gin-gorm-backend/utils"
	"strconv"
	"strings"
)

type TopicService interface {
	CreateTopic(topic *model.Topic) error
	CreateTopicWithValidation(topicRequest *model.CreateTopicRequest) (*model.Topic, error)
	GetAllTopics() ([]model.Topic, error)
	GetTopicByID(id string) (*model.Topic, error)
	UpdateTopic(topic *model.Topic) error
	UpdateTopicWithValidation(id string, topicRequest *model.UpdateTopicRequest) (*model.Topic, error)
	DeleteTopic(id string) error
	GetNextOrder() (int, error)
	MoveTopicToPosition(topicID uint, newOrder int) error
	ValidateTopicName(name string, excludeID uint) error
}

type topicService struct {
	topicRepo repository.TopicRepository
}

func NewTopicService(topicRepo repository.TopicRepository) TopicService {
	return &topicService{topicRepo}
}

func (s *topicService) handleDuplicateOrderError(err error) error {
	if err == nil {
		return nil
	}

	// Check for SQL Server unique constraint violation
	if strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "UNIQUE constraint") ||
		strings.Contains(err.Error(), "Cannot insert duplicate key") {
		return errors.New("order number already exists")
	}

	return err
}

// Topic Service

func (s *topicService) handleDuplicateNameError(err error) error {
	if err == nil {
		return nil
	}

	// Check for SQL Server unique constraint violation on name field
	if strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "UNIQUE constraint") ||
		strings.Contains(err.Error(), "Cannot insert duplicate key") {
		return errors.New("topic name already exists")
	}

	return err
}

func (s *topicService) GetNextOrder() (int, error) {
	topics, err := s.topicRepo.FindAll()
	if err != nil {
		return 0, err
	}

	return utils.GetNextOrder(topics, func(t model.Topic) int { return t.Order }), nil
}

func (s *topicService) CreateTopic(topic *model.Topic) error {
	err := s.topicRepo.Create(topic)
	if err != nil {
		// Check for duplicate name error first
		if s.handleDuplicateNameError(err).Error() == "topic name already exists" {
			return errors.New("topic name already exists")
		}
		return s.handleDuplicateOrderError(err)
	}
	return nil
}

// CreateTopicWithValidation handles all business logic for creating a topic
func (s *topicService) CreateTopicWithValidation(topicRequest *model.CreateTopicRequest) (*model.Topic, error) {
	// Validate topic name uniqueness
	if err := s.ValidateTopicName(topicRequest.Name, 0); err != nil {
		return nil, err
	}

	// Get the next order number
	nextOrder, err := s.GetNextOrder()
	if err != nil {
		return nil, err
	}

	// Create topic with hardcoded values
	topic := &model.Topic{
		Name:      topicRequest.Name,
		Order:     nextOrder,
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Create the topic
	if err := s.CreateTopic(topic); err != nil {
		return nil, err
	}

	return topic, nil
}

func (s *topicService) GetAllTopics() ([]model.Topic, error) {
	return s.topicRepo.FindAll()
}

func (s *topicService) GetTopicByID(id string) (*model.Topic, error) {
	// Convert string to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.topicRepo.FindByID(uint(idUint))
}

func (s *topicService) ValidateTopicName(name string, excludeID uint) error {
	existingTopic, err := s.topicRepo.FindByName(name)
	if err != nil {
		// If error is "record not found", then the name is available
		if err.Error() == "record not found" {
			return nil
		}
		return err
	}

	// If we found a topic with the same name, check if it's the same topic (for updates)
	if existingTopic.ID == excludeID {
		return nil
	}

	return errors.New("topic name already exists")
}

// MoveTopicToPosition moves a specific topic to a new position and reorders all topics accordingly
func (s *topicService) MoveTopicToPosition(topicID uint, newOrder int) error {
	topics, err := s.topicRepo.FindAll()
	if err != nil {
		return err
	}

	// Use the utility function to reorder items with target
	reorderedTopics := utils.ReorderItemsWithTarget(topics,
		func(t model.Topic) int { return t.Order },
		func(t *model.Topic, order int) { t.Order = order },
		func(t model.Topic) interface{} { return t.ID },
		topicID, newOrder)

	// Update all topics in the database
	for _, topic := range reorderedTopics {
		if err := s.topicRepo.Update(&topic); err != nil {
			return err
		}
	}

	return nil
}

func (s *topicService) UpdateTopic(topic *model.Topic) error {
	err := s.topicRepo.Update(topic)
	if err != nil {
		// Check for duplicate name error first
		if s.handleDuplicateNameError(err).Error() == "topic name already exists" {
			return errors.New("topic name already exists")
		}
		return s.handleDuplicateOrderError(err)
	}
	return nil
}

// UpdateTopicWithValidation handles all business logic for updating a topic
func (s *topicService) UpdateTopicWithValidation(id string, topicRequest *model.UpdateTopicRequest) (*model.Topic, error) {
	// Get existing topic to preserve fields
	existingTopic, err := s.GetTopicByID(id)
	if err != nil {
		return nil, errors.New("topic not found")
	}

	// Update only provided fields
	if topicRequest.Name != nil {
		// Validate topic name uniqueness (excluding current topic)
		if err := s.ValidateTopicName(*topicRequest.Name, existingTopic.ID); err != nil {
			return nil, err
		}
		existingTopic.Name = *topicRequest.Name
	}

	if topicRequest.Order != nil {
		// Move topic to the new position and reorder all topics accordingly
		if err := s.MoveTopicToPosition(existingTopic.ID, *topicRequest.Order); err != nil {
			return nil, err
		}

		// Get the updated topic
		updatedTopic, err := s.GetTopicByID(id)
		if err != nil {
			return nil, err
		}

		return updatedTopic, nil
	} else {
		// No order update, just update other fields
		if err := s.UpdateTopic(existingTopic); err != nil {
			return nil, err
		}
	}

	existingTopic.UpdatedBy = "admin"
	return existingTopic, nil
}

func (s *topicService) DeleteTopic(id string) error {
	// Convert string to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	return s.topicRepo.Delete(uint(idUint))
}
