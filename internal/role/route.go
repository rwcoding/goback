package role

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb()

func init() {
	goback.Route("goback.role.list", NewApiList, "角色列表")
	goback.Route("goback.role.add", NewApiAdd, "角色增加")
	goback.Route("goback.role.edit", NewApiEdit, "角色编辑")
	goback.Route("goback.role.info", NewApiInfo, "角色详情")
	goback.Route("goback.role.delete", NewApiDelete, "角色删除")
}
