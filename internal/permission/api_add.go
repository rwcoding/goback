package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name       string `validate:"required,max=30" json:"name"`
	Permission string `validate:"required,min=1,max=100" json:"permission"`
	Type       uint8  `validate:"required,min=1,max=2" json:"type"`
	Gid        uint32 `validate:"required,numeric,min=1" json:"gid"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
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

	p := models.Permission{
		Name:       strings.TrimSpace(request.Name),
		Permission: strings.TrimSpace(request.Permission),
		Type:       request.Type,
		Gid:        request.Gid,
	}

	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}

func verifyGid(gid uint32) bool {
	return db.Take(&models.PermissionGroup{}, gid).RowsAffected > 0
}

func VerifyType(typ uint8) bool {
	return typ == models.PermissionTypeApi || typ == models.PermissionTypeDef
}
