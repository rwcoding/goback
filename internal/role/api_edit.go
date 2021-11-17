package role

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id   uint32 `validate:"required,numeric,min=1" json:"id"`
	Name string `validate:"required,max=30" json:"name"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.Role{}
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的角色")
	}

	p.Name = strings.TrimSpace(request.Name)

	if db().Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}
	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
