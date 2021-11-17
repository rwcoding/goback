package cache

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.session.list", NewApiList, "会话列表")
	goback.Route("goback.session.delete", NewApiDelete, "会话删除")
	goback.Route("goback.session.clean", NewApiClean, "会话清空")
}
