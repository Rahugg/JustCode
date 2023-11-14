package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hw8/config/auth"
	middleware2 "hw8/internal/auth/controller/http/middleware"
	"hw8/internal/auth/controller/http/v1"
	repoPkg "hw8/internal/auth/repository"
	servicePkg "hw8/internal/auth/service"
	httpserver2 "hw8/pkg/auth/httpserver"
	"hw8/pkg/auth/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *auth.Configuration) {
	l := logger.New(cfg.GinMode)
	repo := repoPkg.New(cfg, l)
	// Migrate the tables with gorm.Migrator
	Migrate(repo, l)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg)
	handler := gin.Default()

	v1.NewRouter(handler, service, l, middleware)
	httpServer := httpserver2.New(handler, cfg, httpserver2.Port(cfg.HttpPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("auth - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("auth - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("auth - Run - httpServer.Shutdown: %w", err))
	}

}
