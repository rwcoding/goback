package profile

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/dblog"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Name  string `validate:"required,max=5,min=2"`
	Phone string `validate:"required,numeric,len=11"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {
	adminer := request.ctx.GetAdminer()
	ip := request.ctx.GetRemote()

	ond := dblog.NewONData().OldStruct(map[string]interface{}{
		"name":  adminer.Name,
		"phone": adminer.Phone,
	})

	if db.Model(adminer).Updates(models.Adminer{
		Name:  strings.TrimSpace(request.Name),
		Phone: strings.TrimSpace(request.Phone),
	}).Error != nil {
		return api.NewErrorResponse("更新失败")
	}

	ond.NewStruct(map[string]interface{}{
		"name":  adminer.Name,
		"phone": adminer.Phone,
	})

	dblog.Add(&models.Log{
		AdminerId: adminer.Id,
		Type:      models.LogTypeEditAdminer,
		Msg:       adminer.Name + " 修改自己的信息",
		Details:   ond.Json(),
		Target:    adminer.Id,
		Ip:        ip,
	})
	return api.NewSuccessResponse("更新成功")
}
