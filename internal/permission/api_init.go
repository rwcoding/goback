package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type initRequest struct {
	ctx *boot.Context
}

func NewApiInit(ctx *boot.Context) boot.Logic {
	return &initRequest{ctx: ctx}
}

func (request *initRequest) Run() *api.Response {
	var p []models.Permission
	if db().Where("type=?", models.PermissionTypeApi).Find(&p).Error != nil {
		return api.NewErrorResponse("无效的权限")
	}

	var add []models.Permission
	for k, v := range boot.GetAuthorities() {
		exists := false
		for _, vv := range p {
			if vv.Permission == k {
				exists = true
			}
		}
		if !exists {
			add = append(add, models.Permission{
				Name:       v,
				Permission: k,
				Type:       models.PermissionTypeApi,
			})
		}
	}

	if len(add) > 0 {
		db().Create(&add)
	}

	var del []uint32
	for _, v := range p {
		exists := false
		for kk, _ := range boot.GetAuthorities() {
			if kk == v.Permission {
				exists = true
			}
		}
		if !exists {
			del = append(del, v.Id)
		}
	}

	if len(del) > 0 {
		db().Where("id IN ?", del).Delete(&models.Permission{})
	}

	return api.NewSuccessResponse("已初始化api权限列表")
}
