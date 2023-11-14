package main

import (
	"hw8/config/crm_core"
	"hw8/internal/crm_core/app"
)

func main() {
	// Configuration
	cfg := crm_core.NewConfig()

	// Run
	app.Run(cfg)
}
