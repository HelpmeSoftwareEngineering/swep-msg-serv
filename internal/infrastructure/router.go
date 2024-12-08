package infrastructure

import (
	"net/http"

	"github.com/Ateto1204/swep-msg-serv/internal/delivery"
	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(msgUseCase usecase.MsgUseCase, notifUseCase usecase.NotifUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	})

	msgHandler := delivery.NewMsgHandler(msgUseCase)
	router.POST("/api/msg-create", msgHandler.SaveMsg)
	router.POST("/api/msg-get", msgHandler.GetMsg)
	router.PATCH("/api/msg-read", msgHandler.ReadMsg)
	router.DELETE("/api/msg-del", msgHandler.DeleteMsg)

	notifHandler := delivery.NewNotifHandler(notifUseCase)
	router.POST("/api/notif-create", notifHandler.SaveNotif)
	router.POST("/api/notif-get", notifHandler.GetNotif)
	router.DELETE("/api/notif-del", notifHandler.DeleteNotif)

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
