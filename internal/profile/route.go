package profile

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.profile.edit", NewApiEdit, "个人编辑")
	goback.Route("goback.profile.info", NewApiInfo, "个人信息")
	goback.Route("goback.profile.password", NewApiPassword, "修改密码")
}
