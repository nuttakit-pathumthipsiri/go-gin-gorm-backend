package handler

import (
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	Service service.TopicService
}

func NewTopicHandler(service service.TopicService) *TopicHandler {
	return &TopicHandler{Service: service}
}

// CreateTopic godoc
// @Summary Create a new topic
// @Tags topics
// @Accept json
// @Produce json
// @Param topic body model.CreateTopicRequest true "Topic request object"
// @Success 201 {object} model.Topic
// @Failure 400 {object} model.BadRequestError
// @Failure 500 {object} model.InternalServerError
// @Router /topics [post]
func (h *TopicHandler) CreateTopic(c *gin.Context) {
	var topicRequest model.CreateTopicRequest
	if err := c.ShouldBindJSON(&topicRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topic, err := h.Service.CreateTopicWithValidation(&topicRequest)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, topic)
}

// GetAllTopics godoc
// @Summary Get all topics
// @Tags topics
// @Produce json
// @Success 200 {array} model.Topic
// @Failure 500 {object} model.InternalServerError
// @Router /topics [get]
func (h *TopicHandler) GetAllTopics(c *gin.Context) {
	topics, err := h.Service.GetAllTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

// GetTopicByID godoc
// @Summary Get a topic by ID
// @Tags topics
// @Produce json
// @Param id path string true "Topic ID"
// @Success 200 {object} model.Topic
// @Failure 400 {object} model.BadRequestError
// @Failure 404 {object} model.NotFoundError
// @Router /topics/{id} [get]
func (h *TopicHandler) GetTopicByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	topic, err := h.Service.GetTopicByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topic)
}

// UpdateTopic godoc
// @Summary Update a topic
// @Tags topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Param topic body model.UpdateTopicRequest true "Updated Topic request object"
// @Success 200 {object} model.Topic
// @Failure 400 {object} model.BadRequestError
// @Failure 404 {object} model.NotFoundError
// @Failure 500 {object} model.InternalServerError
// @Router /topics/{id} [put]
func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var topicRequest model.UpdateTopicRequest
	if err := c.ShouldBindJSON(&topicRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topic, err := h.Service.UpdateTopicWithValidation(id, &topicRequest)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, topic)
}

// DeleteTopic godoc
// @Summary Delete a topic
// @Tags topics
// @Param id path string true "Topic ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.BadRequestError
// @Failure 500 {object} model.InternalServerError
// @Router /topics/{id} [delete]
func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	if err := h.Service.DeleteTopic(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
