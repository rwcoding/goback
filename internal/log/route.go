package cache

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb()

func init() {
	goback.Route("goback.log.list", NewApiList, "日志列表")
	goback.Route("goback.log.info", NewApiInfo, "日志信息")
}
