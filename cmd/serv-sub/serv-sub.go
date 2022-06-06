package main

import (
	"log"

	"github.com/perocha/dapr-starter/config"
	"github.com/perocha/dapr-starter/internal/app"
)

//
// Main entry point
//
func main() {
	// Read configuration file
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("main::Failed to read config: %v", err)
	}

	app.Run(cfg)
}
