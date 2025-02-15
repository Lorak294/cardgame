package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize the database instance: %s",err)
	}
}