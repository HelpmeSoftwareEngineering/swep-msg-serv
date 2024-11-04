package infrastructure

import (
	"net/http"

	"github.com/Ateto1204/swep-msg-serv/internal/delivery"
	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(msgUseCase usecase.MsgUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	})

	handler := delivery.NewMsgHandler(msgUseCase)
	router.POST("/api/msg", handler.SaveMsg)
	router.POST("/api/msg/id", handler.GetMsg)
	router.PATCH("api/msg/read", handler.ReadMsg)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
