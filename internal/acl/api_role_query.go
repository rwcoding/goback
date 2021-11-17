package acl

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type roleQueryRequest struct {
	ctx *boot.Context

	RoleId uint32 `validate:"required,numeric,min=1" json:"role_id"`
}

type permission struct {
	Id         uint32 `json:"id"`
	Gid        uint32 `json:"gid"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

type group struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type havePermission struct {
	Permission string `json:"permission"`
}

type roleQueryResponse struct {
	RoleId          uint32           `json:"role_id"`
	RoleName        string           `json:"role_name"`
	Groups          []group          `json:"groups"`
	Permissions     []permission     `json:"permissions"`
	PermissionsHave []havePermission `json:"permissions_have"`
}

func NewApiRoleQuery(ctx *boot.Context) boot.Logic {
	return &roleQueryRequest{ctx: ctx}
}

func (request *roleQueryRequest) Run() *api.Response {
	var role models.Role
	if db().Take(&role, request.RoleId).Error != nil {
		return api.NewErrorResponse("无效的角色")
	}

	var groups []group
	db().Model(&models.PermissionGroup{}).Order("ord").Find(&groups)
	groups = append(groups, group{
		Id:   0,
		Name: "未定义",
	})

	var permissions []permission
	db().Model(&models.Permission{}).Find(&permissions)

	var hp []havePermission
	db().Model(&models.RolePermission{}).Select("permission").Where("role_id=?", request.RoleId).Find(&hp)

	return api.NewDataResponse(&roleQueryResponse{
		RoleName:        role.Name,
		RoleId:          role.Id,
		Groups:          groups,
		Permissions:     permissions,
		PermissionsHave: hp,
	})
}
