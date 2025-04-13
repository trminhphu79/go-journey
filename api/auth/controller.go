package auth

import (
	"app/api/auth/dto"
	"app/arch/network"
	"strings"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type authController struct {
	network.BaseController
	service AuthService
}

func CreateController(
	service AuthService,
	authProvider network.AuthenticationProvider,
) network.Controller {
	return &authController{
		BaseController: network.NewBaseController("api/v1/auth", authProvider),
		service:        service,
	}
}

func (c *authController) AddRouters(group *gin.RouterGroup) {
	group.POST("/registration", c.Registration)
	group.POST("/login", c.Login)
	group.POST("/authenticate", c.Authenticate)
}

// Registration godoc
// @Summary Create new account
// @Description Create new account using username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegistrationDTO true "Registration details"
// @Success 200 {object} model.User
// @Failure 400 {object} network.apiError
// @Failure 404 {object} network.apiError
// @Router /api/v1/auth/registration [post]
func (c *authController) Registration(ctx *gin.Context) {
	var dto dto.RegistrationDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		logger.WithError(err).Warn("Invalid input for registration")
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}

	result, err := c.service.RegisterUser(dto)
	if err != nil {
		c.Send(ctx).ComposeError((err))
		return
	}

	c.Send(ctx).SuccessDataRes("Register user success", result)
}

// Login godoc
// @Summary Login account
// @Description Login account using username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginDTO true "Registration details"
// @Success 200 {object} dto.LoginRessponseDTO
// @Failure 400 {object} network.apiError
// @Failure 404 {object} network.apiError
// @Router /api/v1/auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var dto dto.LoginDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		logger.WithError(err).Warn("Invalid input for registration")
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}
	loginData, err := c.service.Login(dto)

	if err != nil {
		c.Send(ctx).ComposeError((err))
		return
	}

	c.Send(ctx).SuccessDataRes("Register user success", loginData)
}

// Authenticate godoc
// @Summary Authenticate using accessToken
// @Description Authenticate using accessToken
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} network.apiError
// @Failure 404 {object} network.apiError
// @Router /api/v1/auth/authenticate [post]
func (c *authController) Authenticate(ctx *gin.Context) {
	c.service.Authenticate("2")
	var rawToken string
	rawToken = ctx.GetHeader("Authorization")
	if rawToken == "" {
		logger.Error("Missing token header")
		c.Send(ctx).UnauthorizedErr("Unauthorized", nil)
		return
	}

	logger.Info("Get rawToken from header: ", rawToken)

	parts := strings.Split(rawToken, " ")
	if len(parts) < 2 {
		logger.Error("Token format is invalid")
		c.Send(ctx).UnauthorizedErr("Unauthorized", nil)
		return
	}

	accessToken := strings.Split(rawToken, " ")[1]
	if accessToken == "" {
		logger.Error("Token value is empty")
		c.Send(ctx).UnauthorizedErr("Unauthorized", nil)
		return
	}

	logger.Info("accessToken after handle: ", accessToken)
	user, err := c.service.Authenticate(accessToken)
	if err != nil {
		c.Send(ctx).UnauthorizedErr("Unauthorized", err)
		return
	}
	c.Send(ctx).SuccessDataRes("Login success", user)
	ctx.Abort()
}
