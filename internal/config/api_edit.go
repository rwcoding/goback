package config

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id   uint32 `validate:"required,numeric,min=1" json:"id"`
	Name string `validate:"required,max=100" json:"name"`
	K    string `validate:"required,max=200" json:"k"`
	V    string `validate:"required,max=3000" json:"v"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.Config{}
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}

	p.Name = strings.TrimSpace(request.Name)
	p.K = strings.TrimSpace(request.K)
	p.V = strings.TrimSpace(request.V)

	if db().Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}
	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
