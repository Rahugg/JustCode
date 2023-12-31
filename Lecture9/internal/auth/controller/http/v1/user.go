package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/auth/controller/http/middleware"
	"hw8/internal/auth/entity"
	"hw8/internal/auth/service"
	"hw8/pkg/auth/logger"
	"log"
	"net/http"
)

type userRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newUserRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &userRoutes{s, l}

	adminHandler := handler.Group("/admin/user")
	{
		adminHandler.Use(MW.CustomLogger())
		adminHandler.Use(MW.DeserializeUser("admin"))

		//adminHandler.GET("/all", r.GetUsers)
		//adminHandler.PUT("/:id", r.UpdateUser)
		//adminHandler.DELETE("/:id", r.DeleteUser)
		//adminHandler.POST("/", r.CreateUser)

		adminHandler.GET("/test", func(ctx *gin.Context) {
			log.Println("hello from controller")
			ctx.JSON(http.StatusOK, "test")
		})
	}

	userHandler := handler.Group("/user")
	{
		userHandler.POST("/register", r.signUpManager)
		userHandler.POST("/login", r.signIn)
		userHandler.GET("/logout", r.logout)
	}
}

func (ur *userRoutes) signUpAdmin(ctx *gin.Context) {
	ur.signUp(ctx, 1, true, "admin")
}

func (ur *userRoutes) signUpManager(ctx *gin.Context) {
	verified := ur.s.Config.Gin.Mode == "debug"
	ur.signUp(ctx, 2, verified, "manager")
}

func (ur *userRoutes) signUp(ctx *gin.Context, roleId uint, verified bool, provider string) {
	var payload entity.SignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	data, err := ur.s.SignUp(ctx, &payload, roleId, verified, provider)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    data,
	})
}
func (ur *userRoutes) signIn(ctx *gin.Context) {
	var payload entity.SignInInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	data, err := ur.s.SignIn(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    data,
	})
}

func (ur *userRoutes) logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
