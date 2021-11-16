package config

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb()

func init() {
	goback.Route("goback.config.list", NewApiList, "配置列表")
	goback.Route("goback.config.add", NewApiAdd, "配置增加")
	goback.Route("goback.config.edit", NewApiEdit, "配置编辑")
	goback.Route("goback.config.info", NewApiInfo, "配置详情")
	goback.Route("goback.config.delete", NewApiDelete, "配置删除")
}
