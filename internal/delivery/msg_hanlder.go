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
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.msgUseCase.SaveMsg(input.ID, input.Name); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
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

func (h *MsgHandler) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, "hello zeabur")
}
