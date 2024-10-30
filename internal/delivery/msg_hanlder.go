package delivery

import (
	"net/http"

	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MsgHandler struct {
	msgUseCase usecase.MsgUseCase
}

func NewMsgHandler(msgUseCase usecase.MsgUseCase) *MsgHandler {
	return &MsgHandler{msgUseCase}
}

func (h *MsgHandler) SaveMsg(c *gin.Context) {
	type Input struct {
		UserID  string `json:"user_id"`
		Content string `json:"content"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg, err := h.msgUseCase.SaveMsg(input.UserID, input.Content)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (h *MsgHandler) GetMsg(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg, err := h.msgUseCase.GetMsg(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (h *MsgHandler) ReadMsg(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.msgUseCase.ReadMsg(input.ID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Update sucessful")
}
