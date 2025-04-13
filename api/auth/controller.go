package auth

import (
	"app/api/auth/dto"
	"app/arch/network"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type authController struct {
	network.BaseController
	service AuthService
}

func CreateController(
	service AuthService,
) network.Controller {
	return &authController{
		BaseController: network.NewBaseController("api/v1/auth"),
		service:        service,
	}
}

func (c *authController) AddRouters(group *gin.RouterGroup) {
	group.POST("/registration", c.Registration)
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

// Authenticate godoc
// @Summary Authenticate account
// @Description Authenticate account using username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginDTO true "Registration details"
// @Success 200 {object} dto.LoginRessponseDTO
// @Failure 400 {object} network.apiError
// @Failure 404 {object} network.apiError
// @Router /api/v1/auth/authenticate [post]
func (c *authController) Authenticate(ctx *gin.Context) {
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
