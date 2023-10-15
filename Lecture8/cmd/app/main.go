package main

import (
	"hw8/config"
	"hw8/internal/app"
)

func main() {
	// Configuration
	cfg := config.NewConfig()

	// Run
	app.Run(cfg)
}
