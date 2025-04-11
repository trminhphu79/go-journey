package network

type baseController struct {
	ResponseSender
	basePath string
}

func NewBaseController(basePath string) BaseController {
	return &baseController{
		ResponseSender: NewResponseSender(),
		basePath:       basePath,
	}
}

func (c *baseController) Path() string {
	return c.basePath
}
