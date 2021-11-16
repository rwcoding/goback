package role

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

type infoResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.Role
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的角色")
	}
	return api.NewDataResponse(&infoResponse{
		Id:   p.Id,
		Name: p.Name,
	})
}
