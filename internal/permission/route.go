package permission

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.permission.list", NewApiList, "权限列表")
	goback.Route("goback.permission.add", NewApiAdd, "权限增加")
	goback.Route("goback.permission.edit", NewApiEdit, "权限编辑")
	goback.Route("goback.permission.info", NewApiInfo, "权限详情")
	goback.Route("goback.permission.delete", NewApiDelete, "权限删除")
	goback.Route("goback.permission.shift", NewApiShift, "权限转移")
	goback.Route("goback.permission.init", NewApiInit, "权限初始化")

	goback.Route("goback.permission.group.list", NewApiGroupList, "权限分组列表")
	goback.Route("goback.permission.group.add", NewApiGroupAdd, "权限分组增加")
	goback.Route("goback.permission.group.edit", NewApiGroupEdit, "权限分组编辑")
	goback.Route("goback.permission.group.info", NewApiGroupInfo, "权限分组详情")
	goback.Route("goback.permission.group.delete", NewApiGroupDelete, "权限分组删除")
}
