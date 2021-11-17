package role

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strconv"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	var u models.Role
	if db().Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的角色")
	}

	var c int64
	if db().Where("FIND_IN_SET(" + strconv.Itoa(int(request.Id)) + ", roles)").
		Count(&c); c > 0 {
		return api.NewErrorResponse("尚有关联用户，不能删除，请先取消用户关联该角色")
	}

	if db().Delete(&u).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	//删除关联权限
	db().Where("role_id", request.Id).Delete(&models.RolePermission{})

	return api.NewSuccessResponse("删除成功")
}
