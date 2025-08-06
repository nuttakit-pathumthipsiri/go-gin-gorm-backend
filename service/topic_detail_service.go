package service

import (
	"errors"
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/repository"
	"go-gin-gorm-backend/utils"
	"strconv"
	"strings"
)

type TopicDetailService interface {
	CreateTopicDetail(detail *model.TopicDetail) error
	CreateTopicDetailWithValidation(topicID string, detailRequest *model.CreateTopicDetailRequest) (*model.TopicDetail, error)
	GetAllDetailsByTopicID(topicID string) ([]model.TopicDetail, error)
	GetDetailByID(id string) (*model.TopicDetail, error)
	UpdateTopicDetail(detail *model.TopicDetail) error
	UpdateTopicDetailWithValidation(id string, detailRequest *model.UpdateTopicDetailRequest) (*model.TopicDetail, error)
	DeleteTopicDetail(id string) error
	GetNextDetailOrder(topicID string) (int, error)
	MoveTopicDetailToPosition(detailID uint, newOrder int) error
	ValidateTopicDetailName(name string, excludeID uint) error
}

type topicDetailService struct {
	topicDetailRepo repository.TopicDetailRepository
}

func NewTopicDetailService(topicDetailRepo repository.TopicDetailRepository) TopicDetailService {
	return &topicDetailService{topicDetailRepo}
}

func (s *topicDetailService) handleDuplicateOrderError(err error) error {
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

func (s *topicDetailService) handleDuplicateNameError(err error) error {
	if err == nil {
		return nil
	}

	// Check for SQL Server unique constraint violation on name field
	if strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "UNIQUE constraint") ||
		strings.Contains(err.Error(), "Cannot insert duplicate key") {
		return errors.New("topic detail name already exists")
	}

	return err
}

func (s *topicDetailService) GetNextDetailOrder(topicID string) (int, error) {
	details, err := s.GetAllDetailsByTopicID(topicID)
	if err != nil {
		return 0, err
	}

	return utils.GetNextOrder(details, func(d model.TopicDetail) int { return d.Order }), nil
}

func (s *topicDetailService) ValidateTopicDetailName(name string, excludeID uint) error {
	existingDetail, err := s.topicDetailRepo.FindByName(name)
	if err != nil {
		// If error is "record not found", then the name is available
		if err.Error() == "record not found" {
			return nil
		}
		return err
	}

	// If we found a detail with the same name, check if it's the same detail (for updates)
	if existingDetail.ID == excludeID {
		return nil
	}

	return errors.New("topic detail name already exists")
}

func (s *topicDetailService) CreateTopicDetail(detail *model.TopicDetail) error {
	err := s.topicDetailRepo.Create(detail)
	if err != nil {
		// Check for duplicate name error first
		if s.handleDuplicateNameError(err).Error() == "topic detail name already exists" {
			return errors.New("topic detail name already exists")
		}
		return s.handleDuplicateOrderError(err)
	}
	return nil
}

// CreateTopicDetailWithValidation handles all business logic for creating a topic detail
func (s *topicDetailService) CreateTopicDetailWithValidation(topicID string, detailRequest *model.CreateTopicDetailRequest) (*model.TopicDetail, error) {
	// Convert string to uint
	topicIDUint, err := strconv.ParseUint(topicID, 10, 32)
	if err != nil {
		return nil, errors.New("invalid topic ID format")
	}

	// Validate topic detail name uniqueness
	if err := s.ValidateTopicDetailName(detailRequest.Name, 0); err != nil {
		return nil, err
	}

	// Get the next order number for this topic
	nextOrder, err := s.GetNextDetailOrder(topicID)
	if err != nil {
		return nil, err
	}

	// Create detail with hardcoded values
	detail := &model.TopicDetail{
		TopicID:   uint(topicIDUint),
		Name:      detailRequest.Name,
		Order:     nextOrder,
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Create the detail
	if err := s.CreateTopicDetail(detail); err != nil {
		return nil, err
	}

	return detail, nil
}

func (s *topicDetailService) GetAllDetailsByTopicID(topicID string) ([]model.TopicDetail, error) {
	// Convert string to uint
	topicIDUint, err := strconv.ParseUint(topicID, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.topicDetailRepo.FindAllByTopicID(uint(topicIDUint))
}

func (s *topicDetailService) GetDetailByID(id string) (*model.TopicDetail, error) {
	// Convert string to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.topicDetailRepo.FindByID(uint(idUint))
}

func (s *topicDetailService) DeleteTopicDetail(id string) error {
	// Convert string to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	return s.topicDetailRepo.Delete(uint(idUint))
}

// MoveTopicDetailToPosition moves a specific topic detail to a new position and reorders all details accordingly
func (s *topicDetailService) MoveTopicDetailToPosition(detailID uint, newOrder int) error {
	// First, get the topic detail to find its topic ID
	detail, err := s.topicDetailRepo.FindByID(detailID)
	if err != nil {
		return err
	}

	// Get all details for this topic
	details, err := s.GetAllDetailsByTopicID(strconv.FormatUint(uint64(detail.TopicID), 10))
	if err != nil {
		return err
	}

	// Use the utility function to reorder items with target
	reorderedDetails := utils.ReorderItemsWithTarget(details,
		func(d model.TopicDetail) int { return d.Order },
		func(d *model.TopicDetail, order int) { d.Order = order },
		func(d model.TopicDetail) interface{} { return d.ID },
		detailID, newOrder)

	// Update all details in the database
	for _, detail := range reorderedDetails {
		if err := s.topicDetailRepo.Update(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (s *topicDetailService) UpdateTopicDetail(detail *model.TopicDetail) error {
	err := s.topicDetailRepo.Update(detail)
	if err != nil {
		// Check for duplicate name error first
		if s.handleDuplicateNameError(err).Error() == "topic detail name already exists" {
			return errors.New("topic detail name already exists")
		}
		return s.handleDuplicateOrderError(err)
	}
	return nil
}

// UpdateTopicDetailWithValidation handles all business logic for updating a topic detail
func (s *topicDetailService) UpdateTopicDetailWithValidation(id string, detailRequest *model.UpdateTopicDetailRequest) (*model.TopicDetail, error) {
	// Get existing detail to preserve fields
	existingDetail, err := s.GetDetailByID(id)
	if err != nil {
		return nil, errors.New("topic detail not found")
	}

	// Update only provided fields
	if detailRequest.Name != nil {
		// Validate topic detail name uniqueness (excluding current detail)
		if err := s.ValidateTopicDetailName(*detailRequest.Name, existingDetail.ID); err != nil {
			return nil, err
		}
		existingDetail.Name = *detailRequest.Name
	}

	if detailRequest.Order != nil {
		// Move topic detail to the new position and reorder all details accordingly
		if err := s.MoveTopicDetailToPosition(existingDetail.ID, *detailRequest.Order); err != nil {
			return nil, err
		}

		// Get the updated detail
		updatedDetail, err := s.GetDetailByID(id)
		if err != nil {
			return nil, err
		}

		return updatedDetail, nil
	} else {
		// No order update, just update other fields
		if err := s.UpdateTopicDetail(existingDetail); err != nil {
			return nil, err
		}
	}

	existingDetail.UpdatedBy = "admin"
	return existingDetail, nil
}
