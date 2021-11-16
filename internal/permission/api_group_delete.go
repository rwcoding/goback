package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type groupDeleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

type groupDeleteResponse struct {
	GroupNames map[uint32]string `json:"group_names" json:"group_names"`
}

func NewApiGroupDelete(ctx *boot.Context) boot.Logic {
	return &groupDeleteRequest{ctx: ctx}
}

func (request *groupDeleteRequest) Run() *api.Response {
	var u models.PermissionGroup
	if db.Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的分组")
	}
	if db.Delete(&u).Error != nil {
		return api.NewErrorResponse("删除失败")
	}

	//更新原有权限分组为0（未定义分组）
	db.Model(&models.PermissionGroup{}).Where("gid=?", request.Id).Update("gid", 0)

	return api.NewMDResponse("删除成功", &groupDeleteResponse{
		GroupNames: groupNames(),
	})
}
