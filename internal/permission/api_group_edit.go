package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type groupEditRequest struct {
	ctx *boot.Context

	Id   uint32 `validate:"required,numeric,min=1" json:"id"`
	Name string `validate:"required,max=30" json:"name"`
	Ord  uint32 `validate:"required,numeric,min=1" json:"ord"`
}

type groupEditResponse struct {
	GroupNames map[uint32]string `json:"group_names"`
}

func NewApiGroupEdit(ctx *boot.Context) boot.Logic {
	return &groupEditRequest{ctx: ctx}
}

func (request *groupEditRequest) Run() *api.Response {
	p := models.PermissionGroup{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的分组")
	}

	p.Name = request.Name
	p.Ord = request.Ord

	if db.Save(&p).Error != nil {
		return api.NewErrorResponse("编辑失败")
	}
	return api.NewMDResponse("编辑成功", &groupEditResponse{
		GroupNames: groupNames(),
	})
}
