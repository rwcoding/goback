package login

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type logoutRequest struct {
	ctx *boot.Context
}

func NewApiLogout(ctx *boot.Context) boot.Logic {
	return &logoutRequest{ctx: ctx}
}

func (request logoutRequest) Run() *api.Response {
	db().Where("sess=?", request.ctx.GetSession()).Delete(&models.Session{})
	// 是否要主要所有该用户相关的session
	return api.NewSuccessResponse("注销成功，请重新登录")
}
