package adminer

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func init() {
	goback.Route("goback.adminer.list", NewApiList, "管理列表")
	goback.Route("goback.adminer.add", NewApiAdd, "管理增加")
	goback.Route("goback.adminer.edit", NewApiEdit, "管理编辑")
	goback.Route("goback.adminer.info", NewApiInfo, "管理信息")
	goback.Route("goback.adminer.delete", NewApiDelete, "管理删除")
}
