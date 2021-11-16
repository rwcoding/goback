package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id         uint32 `validate:"required,numeric,min=1" json:"id"`
	Name       string `validate:"required,max=30" json:"name"`
	Permission string `validate:"required,min=1,max=100" json:"permission"`
	Type       uint8  `validate:"required,min=1,max=2" json:"type"`
	Gid        uint32 `validate:"required,numeric,min=1" json:"gid"`
}

type editResponse struct {
	Id uint32 `json:"id"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.Permission{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的权限")
	}

	if !verifyGid(request.Gid) {
		return api.NewErrorResponse("无效的分组")
	}

	if !VerifyType(request.Type) {
		return api.NewErrorResponse("无效的类型")
	}

	//验证是否在api中
	if request.Type == models.PermissionTypeApi {
		isIn := false
		for k, _ := range boot.GetAuthorities() {
			if k == request.Permission {
				isIn = true
			}
		}
		if !isIn {
			return api.NewErrorResponse("无效的api权限")
		}
	}

	p.Name = strings.TrimSpace(request.Name)
	p.Permission = strings.TrimSpace(request.Permission)
	p.Gid = request.Gid
	p.Type = request.Type

	if db.Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}
	return api.NewMDResponse("修改成功", &editResponse{
		Id: p.Id,
	})
}
