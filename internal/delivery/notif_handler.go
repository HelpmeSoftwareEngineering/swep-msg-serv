package delivery

import (
	"net/http"

	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

type NotifHandler struct {
	notifUseCase usecase.NotifUseCase
}

func NewNotifHandler(notifUseCase usecase.NotifUseCase) *NotifHandler {
	return &NotifHandler{notifUseCase}
}

func (h *NotifHandler) SaveNotif(c *gin.Context) {
	type Input struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	notif, err := h.notifUseCase.SaveNotif(input.UserID, input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notif)
}

func (h *NotifHandler) GetNotif(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	notif, err := h.notifUseCase.GetNotif(input.ID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notif)
}

func (h *NotifHandler) DeleteNotif(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.notifUseCase.DeleteNotif(input.ID); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
