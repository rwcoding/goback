package acl

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type batchQueryRequest struct {
	ctx *boot.Context
}

type role struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type batchQueryResponse struct {
	Roles       []role       `json:"roles"`
	Groups      []group      `json:"groups"`
	Permissions []permission `json:"permissions"`
}

func NewApiBatchQuery(ctx *boot.Context) boot.Logic {
	return &batchQueryRequest{ctx: ctx}
}

func (request *batchQueryRequest) Run() *api.Response {
	var groups []group
	db().Model(&models.PermissionGroup{}).Order("ord").Find(&groups)
	groups = append(groups, group{
		Id:   0,
		Name: "未定义",
	})

	var permissions []permission
	db().Model(&models.Permission{}).Find(&permissions)

	var roles []role
	db().Model(&models.Role{}).Find(&roles)

	return api.NewDataResponse(&batchQueryResponse{
		Roles:       roles,
		Groups:      groups,
		Permissions: permissions,
	})
}
