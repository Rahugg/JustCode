package main

import (
	"hw8/config/auth"
	"hw8/internal/auth/app"
)

func main() {
	// Configuration
	cfg := auth.NewConfig()

	// Run
	app.Run(cfg)
}
