package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hw8/config/crm_core"
	middleware2 "hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/controller/http/v1"
	debugRoute "hw8/internal/crm_core/controller/http/v1/debug"
	repoPkg "hw8/internal/crm_core/repository"
	servicePkg "hw8/internal/crm_core/service"
	"hw8/internal/crm_core/transport"
	"hw8/pkg/crm_core/cache"
	"hw8/pkg/crm_core/httpserver/debug"
	"hw8/pkg/crm_core/httpserver/public"
	"hw8/pkg/crm_core/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *crm_core.Configuration) {
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)
	// migrate the tables with gorm.Migrator
	Migrate(repo, l)

	//REDIS implementation
	redisClient, err := cache.NewRedisClient()
	if err != nil {
		return
	}

	contactCache := cache.NewContactCache(redisClient, 10*time.Minute)

	validateTransport := transport.NewTransport(cfg)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg, validateTransport)
	handler := gin.Default()
	handlerDebug := gin.Default()

	v1.NewRouter(handler, service, l, middleware, contactCache)
	debugRoute.NewDebugRouter(handlerDebug, l)
	httpServer := public.New(handler, cfg, public.Port(cfg.HTTP.Port))
	debugServer := debug.New(handlerDebug, cfg, debug.Port(cfg.HTTP.DebugPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("crm_system - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Notify: %w", err))
	case err = <-debugServer.Notify():
		l.Error(fmt.Errorf("crm_system - Run - debugServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Shutdown: %w", err))
	}

	err = debugServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("crm_system - Run - debugServer.Shutdown: %w", err))
	}

}
