package config

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	var u models.Config
	if db().Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}

	if db().Delete(&u).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	return api.NewSuccessResponse("删除成功")
}
