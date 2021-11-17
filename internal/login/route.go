package login

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.login", NewApiLogin, "登录")
	goback.Route("goback.login.captcha", NewApiCaptcha, "验证码")
	goback.Route("goback.login.logout", NewApiLogout, "退出")
}
