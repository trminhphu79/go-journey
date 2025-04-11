package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sender struct{}

func NewResponseSender() ResponseSender {
	return &sender{}
}

func (m *sender) Debug() bool {
	return gin.Mode() != gin.ReleaseMode
}

func (m *sender) Send(ctx *gin.Context) SendResponse {
	return &send{
		debug:   m.Debug(),
		context: ctx,
	}
}

type send struct {
	debug   bool
	context *gin.Context
}

func (s *send) SuccessMsgRes(msg string) {
	s.sendResponse(NewSuccessMsgRes((msg)))
}

func (s *send) SuccessDataRes(msg string, data any) {
	s.sendResponse(NewSuccessDataRes(msg, data))
}

func (s *send) BadRequestErr(msg string, err error) {
	s.sendError(NewBadRequestErr(msg, err))
}

func (s *send) UnauthorizedEr(msg string, err error) {
	s.sendError(NewUnauthorizedErr(msg, err))
}

func (s *send) NotFoundErr(msg string, err error) {
	s.sendError(NewNotFoundErr(msg, err))
}

func (s *send) InternalServerErr(msg string, err error) {
	s.sendError(NewInternalServerErr(msg, err))
}

func (s *send) ForbiddenErr(msg string, err error) {
	s.sendError(NewForbiddenErr(msg, err))
}

func (s *send) sendResponse(response Response) {
	s.context.JSON(int(response.GetStatus()), response)
	s.context.Abort()
}

func (s *send) sendError(err ApiError) {
	var res Response

	switch err.GetCode() {
	case http.StatusBadRequest:
		res = NewBadRequestRes(err.GetMessage())
	case http.StatusForbidden:
		res = NewForbiddenRes(err.GetMessage())
	case http.StatusUnauthorized:
		res = NewUnauthorizedRes(err.GetMessage())
	case http.StatusNotFound:
		res = NewNotFoundRes(err.GetMessage())
	case http.StatusInternalServerError:
		if s.debug {
			res = NewInternalServerErrorRes(err.Unwrap().Error())
		}
	default:
		if s.debug {
			res = NewInternalServerErrorRes(err.Unwrap().Error())
		}
	}

	if res == nil {
		res = NewInternalServerErrorRes("An unexpected error occurred. Please try again later.")
	}

	s.sendResponse(res)
}
