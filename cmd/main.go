package main

import (
	"log"

	"github.com/kunal70006/realtime-chat-go/db"
	"github.com/kunal70006/realtime-chat-go/internal/user"
	"github.com/kunal70006/realtime-chat-go/router"
)

func main() {

	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not init database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
