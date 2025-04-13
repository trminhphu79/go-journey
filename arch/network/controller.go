package network

type baseController struct {
	ResponseSender
	authProvider AuthenticationProvider
	basePath     string
}

func NewBaseController(basePath string, authProvider AuthenticationProvider) BaseController {
	return &baseController{
		ResponseSender: NewResponseSender(),
		basePath:       basePath,
		authProvider:   authProvider,
	}
}

func (c *baseController) Path() string {
	return c.basePath
}
