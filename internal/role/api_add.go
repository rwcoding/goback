package role

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name string `validate:"required,max=30" json:"name"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.Role{
		Name: strings.TrimSpace(request.Name),
	}

	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
