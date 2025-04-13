package middleware

import (
	"app/api/auth"
	"app/arch/network"
	"app/utils"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type authenticateHandler struct {
	network.ResponseSender
	network.ContextPayload
	authService auth.AuthService
}

func NewAuthenticateHandler(authService auth.AuthService) network.AuthenticationProvider {
	return &authenticateHandler{
		ResponseSender: network.NewResponseSender(),
		authService:    authService,
	}
}

func (m *authenticateHandler) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(network.Authorization)
		logger.Info("Middleware is running...", authHeader)
		if len(authHeader) == 0 {
			m.Send(ctx).UnauthorizedErr("permission denied: missing Authorization", nil)
			return
		}

		token := utils.ExtractBearerToken(authHeader)
		if token == "" {
			m.Send(ctx).UnauthorizedErr("permission denied: invalid Authorization", nil)
			return
		}

		claims, err := m.authService.ValidateAccessToken(token)
		if err != nil {
			m.Send(ctx).UnauthorizedErr(err.Error(), err)
			return
		}

		sub, ok := claims["sub"].(string)

		if !ok {
			m.Send(ctx).UnauthorizedErr("permission denied: claims subject does not exists", err)
			return
		}

		user, err := m.authService.FindUserById(sub)
		if err != nil {
			m.Send(ctx).UnauthorizedErr("permission denied: claims subject does not exists", err)
			return
		}

		m.SetUser(ctx, user)
		ctx.Next()
	}
}
