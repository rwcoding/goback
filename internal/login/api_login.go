package login

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/dblog"
	"github.com/rwcoding/goback/pkg/logger"
	"github.com/rwcoding/goback/pkg/session"
	"github.com/rwcoding/goback/pkg/util"
)

type loginRequest struct {
	ctx *boot.Context

	Type     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Code     string `validate:"required"`
	ImgId    string `validate:"required" json:"img_id"`
}

type loginResponse struct {
	SessionId  string `json:"session_id"`
	SessionKey string `json:"session_key"`
}

func NewApiLogin(ctx *boot.Context) boot.Logic {
	return &loginRequest{ctx: ctx}
}

// Run 当前只处理账号密码登录
func (request *loginRequest) Run() *api.Response {
	ip := request.ctx.GetRemote()
	if !session.VerifySession(request.ImgId, request.Code) {
		return api.NewErrorResponse("验证码错误或已经过期")
	}

	adminer := &models.Adminer{}
	if db.Where("username=?", request.Username).Take(adminer).Error != nil {
		return api.NewErrorResponse("无效的账号或密码")
	}

	if util.Password(request.Password, adminer.Salt, true) != adminer.Password {
		return api.NewErrorResponse("无效的账号或密码")
	}

	if !adminer.IsOK() {
		return api.NewErrorResponse("账号已经锁定,请联系管理员")
	}

	db.Where("adminer_id", adminer.Id).Delete(models.Session{})

	sess := session.NewAuthSession(adminer.Id)
	if sess == nil {
		logger.Error("添加session失败")
		return api.NewErrorResponse("登录失败")
	}
	dblog.Add(&models.Log{
		AdminerId: adminer.Id,
		Type:      models.LogTypeLogin,
		Msg:       adminer.Name + " 登录",
		Ip:        ip,
	})

	return api.NewDataResponse(&loginResponse{
		SessionKey: sess.SessionValue,
		SessionId:  sess.SessionId,
	})
}
