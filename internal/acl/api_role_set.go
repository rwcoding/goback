package acl

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strconv"
	"strings"
)

type roleSetRequest struct {
	ctx *boot.Context

	RoleId      uint32 `validate:"required,numeric,min=1" json:"role_id"`
	Permissions string `validate:"required,max=10000" json:"permissions"`
}

func NewApiRoleSet(ctx *boot.Context) boot.Logic {
	return &roleSetRequest{ctx: ctx}
}

func (request *roleSetRequest) Run() *api.Response {
	if db.Take(&models.Role{}, request.RoleId).Error != nil {
		return api.NewErrorResponse("无效的角色")
	}

	var permissions []models.RolePermission
	if request.Permissions != "" {
		var ids []uint32
		ps := strings.Split(request.Permissions, ",")
		for _, v := range ps {
			id, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			ids = append(ids, uint32(id))
		}
		var pList []models.Permission
		if db.Where("id IN ?", ids).Find(&pList).RowsAffected > 0 {
			for _, v := range pList {
				permissions = append(permissions, models.RolePermission{
					Permission: v.Permission,
					RoleId:     request.RoleId,
				})
			}
		}
	}

	db.Where("role_id=?", request.RoleId).Delete(&models.RolePermission{})
	if len(permissions) > 0 {
		db.Create(&permissions)
	}
	acl.Generate(request.RoleId)

	return api.NewSuccessResponse("设置成功")
}
