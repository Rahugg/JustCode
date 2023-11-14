package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hw8/config/crm_core"
	middleware2 "hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/controller/http/v1"
	repoPkg "hw8/internal/crm_core/repository"
	servicePkg "hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/cache"
	httpserverPkg "hw8/pkg/crm_core/httpserver"
	"hw8/pkg/crm_core/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *crm_core.Configuration) {
	l := logger.New(cfg.GinMode)
	repo := repoPkg.New(cfg, l)
	// migrate the tables with gorm.Migrator
	Migrate(repo, l)

	//REDIS implementation
	redisClient, err := cache.NewRedisClient()
	if err != nil {
		return
	}

	contactCache := cache.NewContactCache(redisClient, 10*time.Minute)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg)
	handler := gin.Default()

	v1.NewRouter(handler, service, l, middleware, contactCache)
	httpServer := httpserverPkg.New(handler, cfg, httpserverPkg.Port(cfg.HttpPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("crm_system - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Shutdown: %w", err))
	}

}
