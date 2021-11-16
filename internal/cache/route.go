package cache

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb()

func init() {
	goback.Route("goback.cache.list", NewApiList, "缓存列表")
	goback.Route("goback.cache.info", NewApiInfo, "缓存信息")
	goback.Route("goback.cache.delete", NewApiDelete, "缓存删除")
	goback.Route("goback.cache.generate", NewApiGenerate, "重生权限")
}
