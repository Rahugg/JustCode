package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hw8/config"
	middleware2 "hw8/internal/controller/http/middleware"
	"hw8/internal/controller/http/v1"
	repoPkg "hw8/internal/repository"
	servicePkg "hw8/internal/service"
	"hw8/pkg/httpserver"
	"hw8/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Configuration) {
	l := logger.New(cfg.GinMode)
	repo := repoPkg.New(cfg, l)
	// Migrate the tables with gorm.Migrator
	Migrate(repo, l)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg)
	handler := gin.Default()

	v1.NewRouter(handler, service, l, middleware)
	httpServer := httpserver.New(handler, cfg, httpserver.Port(cfg.HttpPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
