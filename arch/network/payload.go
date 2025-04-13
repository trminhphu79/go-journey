package network

import (
	"app/api/auth/model"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	payloadUser string = "user"
)

type ContextPayload interface {
	SetUser(ctx *gin.Context, value *model.User)
	MustGetUser(ctx *gin.Context) *model.User
}

type payload struct{}

func CreateContextPayload() ContextPayload {
	return &payload{}
}

func (u *payload) SetUser(ctx *gin.Context, value *model.User) {
	ctx.Set(payloadUser, value)
}

func (u *payload) MustGetUser(ctx *gin.Context) *model.User {
	value, ok := ctx.MustGet(payloadUser).(*model.User)
	if !ok {
		panic(errors.New(payloadUser + " missing for context"))
	}
	return value
}
