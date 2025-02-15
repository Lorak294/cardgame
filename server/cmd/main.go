package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize the database instance: %s",err)
	}

	userRep := user.NewRepository(dbConn.GetDb())
	userSrv := user.NewService(userRep)
	userHandler := user.NewHandler(userSrv)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler,wsHandler)
	router.StartRouter("0.0.0.0:8080")

}