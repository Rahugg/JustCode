package middleware

import (
	"github.com/gin-gonic/gin"
	"hw8/config/crm_core"
	"hw8/internal/crm_core/entity"
	"hw8/internal/crm_core/repository"
	"hw8/internal/crm_core/transport"
	"log"
	"net/http"
	"strings"
)

type Middleware struct {
	Repo              *repository.CRMSystemRepo
	Config            *crm_core.Configuration
	validateTransport *transport.ValidateTransport
}

func New(repo *repository.CRMSystemRepo, config *crm_core.Configuration, validateTransport *transport.ValidateTransport) *Middleware {
	return &Middleware{
		Repo:              repo,
		Config:            config,
		validateTransport: validateTransport,
	}
}

func (m *Middleware) CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before")

		c.Next()

		log.Println("after")
	}
}

func (m *Middleware) DeserializeUser(roles ...interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
				Status:  -1,
				Message: "You are not logged in",
			})
			return
		}

		resp, err := m.validateTransport.Validate(ctx, accessToken, roles)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &entity.CustomResponse{
				Status:  -1,
				Message: err.Error(),
			})
			return
		}
		ctx.Set("currentUser", resp.CurrentUser)
		ctx.Set("currentRole", resp.CurrentRole)
		ctx.Next()
	}
}
