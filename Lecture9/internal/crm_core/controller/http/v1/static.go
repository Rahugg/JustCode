package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/logger"
)

type staticRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newStaticRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	//r := &staticRoutes{s, l}

	staticHandler := handler.Group("/static")
	{
		//middleware for users
		//staticHandler.GET("/:file_name_extension", r.getFile)
		//staticHandler.Static("/", "./Lecture9/internal/crm_core/files")
		staticHandler.Static("/", "./internal/crm_core/files")
	}
}

//func (sh *staticRoutes) getFile(ctx *gin.Context) {
//
//}
