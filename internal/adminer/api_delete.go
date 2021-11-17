package adminer

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/dblog"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	adminer := request.ctx.GetAdminer()
	ip := request.ctx.GetRemote()

	var u models.Adminer
	if db().Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的用户")
	}
	if u.Id == 1 {
		return api.NewErrorResponse("该用户无法删除")
	}
	if u.IsSuper == 1 && adminer.Id != 1 {
		return api.NewErrorResponse("您无权操作该用户")
	}

	if db().Delete(&u).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}
	dblog.Add(&models.Log{
		AdminerId: adminer.Id,
		Type:      models.LogTypeDeleteAdminer,
		Msg:       adminer.Name + " 删除 " + u.Name,
		Target:    u.Id,
		Ip:        ip,
	})
	return api.NewSuccessResponse("删除成功")
}
