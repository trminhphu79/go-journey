package main

import (
	"app/internal/database"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Application init...")

	db, err := database.Connect()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connect successfully: ", db)

}
