package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type groupAddRequest struct {
	ctx *boot.Context

	Name string `validate:"required,max=30" json:"name"`
	Ord  uint32 `validate:"required,numeric,min=1" json:"ord"`
}

type groupAddResponse struct {
	Id         uint32            `json:"id"`
	GroupNames map[uint32]string `json:"group_names"`
}

func NewApiGroupAdd(ctx *boot.Context) boot.Logic {
	return &groupAddRequest{ctx: ctx}
}

func (request *groupAddRequest) Run() *api.Response {
	p := models.PermissionGroup{
		Name: strings.TrimSpace(request.Name),
		Ord:  request.Ord,
	}
	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}
	return api.NewMDResponse("添加成功", &groupAddResponse{
		Id:         p.Id,
		GroupNames: groupNames(),
	})
}
