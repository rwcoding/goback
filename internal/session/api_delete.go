package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	session := models.Session{}
	if db.Take(&session, request.Id).Error != nil {
		return api.NewErrorResponse("无效的会话")
	}
	db.Delete(session)
	return api.NewSuccessResponse("删除成功")
}
