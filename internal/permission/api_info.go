package permission

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
	Id         uint32            `json:"id"`
	Name       string            `json:"name"`
	Permission string            `json:"permission"`
	Gid        uint32            `json:"gid"`
	Type       uint8             `json:"type"`
	GroupNames map[uint32]string `json:"group_names"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.Permission
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的权限")
	}
	return api.NewDataResponse(&infoResponse{
		Id:         p.Id,
		Name:       p.Name,
		Permission: p.Permission,
		Gid:        p.Gid,
		Type:       p.Type,
		GroupNames: groupNames(),
	})
}
