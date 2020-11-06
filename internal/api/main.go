package main

import (
	"log"

	"github.com/errntry/errntry/internal/api/route"
	"github.com/errntry/errntry/internal/config"
)

func main() {
	_, err := config.Init()
	if err != nil {
		log.Fatalf("Failed to load config: %q", err)
	}
	if err := route.Run(); err != nil {
		log.Fatalln(err)
	}
}
