package infrastructure

import (
	"net/http"

	"github.com/Ateto1204/swep-msg-serv/internal/delivery"
	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(msgUseCase usecase.MsgUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	})

	handler := delivery.NewMsgHandler(msgUseCase)
	router.POST("/api/msg", handler.SaveMsg)
	router.POST("/api/msg/id", handler.GetMsg)
	router.GET("/", handler.Handle)

	return router
}
