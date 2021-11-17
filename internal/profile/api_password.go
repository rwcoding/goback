package profile

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/dblog"
	"github.com/rwcoding/goback/pkg/logger"
	"github.com/rwcoding/goback/pkg/util"
	"strings"
)

type passwordRequest struct {
	ctx *boot.Context

	OldPassword string `validate:"required,max=30,min=5" json:"old_password"`
	NewPassword string `validate:"required,max=30,min=5" json:"new_password"`
}

func NewApiPassword(ctx *boot.Context) boot.Logic {
	return &passwordRequest{ctx: ctx}
}

func (request *passwordRequest) Run() *api.Response {
	adminer := request.ctx.GetAdminer()
	ip := request.ctx.GetRemote()

	old := strings.TrimSpace(request.OldPassword)
	newPwd := strings.TrimSpace(request.NewPassword)

	if old == newPwd {
		return api.NewErrorResponse("新旧密码相同")
	}

	if util.Password(old, adminer.Salt, false) != adminer.Password {
		return api.NewErrorResponse("原密码错误")
	}

	salt := util.RandString(32)
	pwd := util.Password(newPwd, salt, false)

	ret := db().Model(adminer).Updates(&models.Adminer{
		Salt:     salt,
		Password: pwd,
	})

	if ret.Error == nil {
		dblog.Add(&models.Log{
			AdminerId: adminer.Id,
			Type:      models.LogTypePassword,
			Msg:       adminer.Name + " 修改 " + adminer.Name + " 的密码",
			Target:    adminer.Id,
			Ip:        ip,
		})
		return api.NewResponse(api.CodeNeedLogin, "密码修改成功，请重新登录", nil)
	} else {
		logger.Error(ret.Error.Error())
	}

	return api.NewErrorResponse("更新失败")
}
