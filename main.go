package main

import (
	"log"

	"github.com/Ateto1204/swep-msg-serv/internal/infrastructure"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
	"github.com/Ateto1204/swep-msg-serv/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	msgRepo := repository.NewMsgRepository(db)
	msgUseCase := usecase.NewMsgUseCase(msgRepo)

	notifRepo := repository.NewNotifRepository(db)
	notifUseCase := usecase.NewNotifUseCase(notifRepo)

	router := infrastructure.NewRouter(msgUseCase, notifUseCase)
	log.Println("Server Start:")
	router.Run(":8080")
}
