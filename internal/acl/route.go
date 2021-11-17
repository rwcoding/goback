package acl

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.acl.role.query", NewApiRoleQuery, "角色权限查询")
	goback.Route("goback.acl.role.set", NewApiRoleSet, "角色权限设置")
	goback.Route("goback.acl.batch.query", NewApiBatchQuery, "权限查询")
	goback.Route("goback.acl.batch.set", NewApiBatchSet, "权限批量设置")
}
