package handler

import (
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopicDetailHandler struct {
	Service service.TopicDetailService
}

func NewTopicDetailHandler(service service.TopicDetailService) *TopicDetailHandler {
	return &TopicDetailHandler{Service: service}
}

// CreateTopicDetail godoc
// @Summary Create a new topic detail
// @Tags topic-details
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic ID"
// @Param detail body model.CreateTopicDetailRequest true "Topic Detail object"
// @Success 201 {object} model.TopicDetail
// @Failure 400 {object} model.BadRequestError
// @Failure 500 {object} model.InternalServerError
// @Router /topics/{id}/details [post]
func (h *TopicDetailHandler) CreateTopicDetail(c *gin.Context) {
	topicID := c.Param("id")
	if topicID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Topic ID is required"})
		return
	}

	var detailRequest model.CreateTopicDetailRequest
	if err := c.ShouldBindJSON(&detailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := h.Service.CreateTopicDetailWithValidation(topicID, &detailRequest)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, detail)
}

// GetAllDetailsByTopicID godoc
// @Summary Get all details for a topic
// @Tags topic-details
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic ID"
// @Success 200 {array} model.TopicDetail
// @Failure 400 {object} model.BadRequestError
// @Failure 500 {object} model.InternalServerError
// @Router /topics/{id}/details [get]
func (h *TopicDetailHandler) GetAllDetailsByTopicID(c *gin.Context) {
	topicID := c.Param("id")
	if topicID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Topic ID is required"})
		return
	}

	details, err := h.Service.GetAllDetailsByTopicID(topicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

// GetDetailByID godoc
// @Summary Get a topic detail by ID
// @Tags topic-details
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic Detail ID"
// @Success 200 {object} model.TopicDetail
// @Failure 400 {object} model.BadRequestError
// @Failure 404 {object} model.NotFoundError
// @Router /details/{id} [get]
func (h *TopicDetailHandler) GetDetailByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	detail, err := h.Service.GetDetailByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detail)
}

// UpdateTopicDetail godoc
// @Summary Update a topic detail
// @Tags topic-details
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic Detail ID"
// @Param detail body model.UpdateTopicDetailRequest true "Updated Topic Detail object"
// @Success 200 {object} model.TopicDetail
// @Failure 400 {object} model.BadRequestError
// @Failure 404 {object} model.NotFoundError
// @Failure 500 {object} model.InternalServerError
// @Router /details/{id} [put]
func (h *TopicDetailHandler) UpdateTopicDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var detailRequest model.UpdateTopicDetailRequest
	if err := c.ShouldBindJSON(&detailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := h.Service.UpdateTopicDetailWithValidation(id, &detailRequest)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, detail)
}

// DeleteTopicDetail godoc
// @Summary Delete a topic detail
// @Tags topic-details
// @Security BearerAuth
// @Param id path string true "Topic Detail ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.BadRequestError
// @Failure 500 {object} model.InternalServerError
// @Router /details/{id} [delete]
func (h *TopicDetailHandler) DeleteTopicDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	if err := h.Service.DeleteTopicDetail(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
