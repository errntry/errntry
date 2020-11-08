package main

import (
	"log"

	"github.com/errntry/errntry/internal/api/route"
	"github.com/profclems/go-dotenv"
)

func main() {
	dotenv.SetConfigFile(".env")
	err := dotenv.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %q", err)
	}
	if err := route.Run(); err != nil {
		log.Fatalln(err)
	}
}
