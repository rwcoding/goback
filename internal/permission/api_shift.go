package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strconv"
)

type shiftRequest struct {
	ctx *boot.Context

	Kw  string `validate:"required,max=100,min=3" json:"kw"`
	Gid uint32 `validate:"omitempty,numeric,min=0" json:"gid"`
}

func NewApiShift(ctx *boot.Context) boot.Logic {
	return &shiftRequest{ctx: ctx}
}

func (request *shiftRequest) Run() *api.Response {
	if request.Gid > 0 && db().Take(&models.PermissionGroup{}, request.Gid).Error != nil {
		return api.NewErrorResponse("无效的分组")
	}

	var ps []models.Permission

	if db().Model(&models.Permission{}).Where("permission LIKE ?", "%"+request.Kw+"%").Find(&ps).RowsAffected == 0 {
		return api.NewErrorResponse("没有对应的权限")
	}

	ids := make([]uint32, len(ps))
	for _, v := range ps {
		ids = append(ids, v.Id)
	}
	rows := db().Model(&models.Permission{}).Where("id IN ?", ids).Update("gid", request.Gid).RowsAffected
	return api.NewSuccessResponse("转移了 " + strconv.Itoa(int(rows)) + " 条数据")
}
