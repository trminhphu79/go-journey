package network

import "github.com/gin-gonic/gin"

type Response interface {
	GetResCode() ResCode
	GetStatus() int
	GetMessage() string
	GetData() any
}

type ApiError interface {
	GetCode() int
	GetMessage() string
	Error() string
	Unwrap() error
}

type SendResponse interface {
	SuccessMsgRes(message string)
	SuccessDataRes(message string, data any)
	BadRequestErr(message string, err error)
	ForbiddenErr(message string, err error)
	UnauthorizedEr(message string, err error)
	NotFoundErr(message string, err error)
	InternalServerErr(message string, err error)
}

type ResponseSender interface {
	Debug() bool
	Send(ctx *gin.Context) SendResponse
}

type BaseController interface {
	ResponseSender
	Path() string
}

type Controller interface {
	BaseController
	AddRouters(group *gin.RouterGroup)
}

type BaseRouter interface {
	GetEngine() *gin.Engine
	Start(ip string, port uint16)
}

type Router interface {
	BaseRouter
	InitControllers(controllers []Controller)
}

type BaseModule[T any] interface {
	GetInstance() *T
}

type Module[T any] interface {
	BaseModule[T]
	Controllers() []Controller
}
