package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type cleanRequest struct {
	ctx *boot.Context

	Type uint32 `validate:"required,numeric,min=1"`
}

func NewApiClean(ctx *boot.Context) boot.Logic {
	return &cleanRequest{ctx: ctx}
}

func (request *cleanRequest) Run() *api.Response {
	if request.Type == 0 {
		db().Delete(models.Session{})
	} else if request.Type == models.SessionTypeAuth {
		db().Where("type=?", models.SessionTypeAuth).Delete(models.Session{})
	} else if request.Type == models.SessionTypeCaptcha {
		db().Where("type=?", models.SessionTypeCaptcha).Delete(models.Session{})
	}

	return api.NewSuccessResponse("清理成功")
}
