package main

import (
	"log"
	"os"

	"github.com/DearCaprga/IdeaForge/internal/app"
)

func main() {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = ":8080"
	}

	if err := app.NewRouter().Run(address); err != nil {
		log.Fatal(err)
	}
}
