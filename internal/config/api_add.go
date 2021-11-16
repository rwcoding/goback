package config

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name string `validate:"required,max=100" json:"name"`
	K    string `validate:"required,max=200" json:"k"`
	V    string `validate:"required,max=3000" json:"v"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.Config{
		Name: strings.TrimSpace(request.Name),
		K:    strings.TrimSpace(request.K),
		V:    strings.TrimSpace(request.V),
	}

	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
