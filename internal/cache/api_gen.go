package cache

import (
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type generateRequest struct {
	ctx *boot.Context
}

func NewApiGenerate(ctx *boot.Context) boot.Logic {
	return &generateRequest{ctx: ctx}
}

func (request *generateRequest) Run() *api.Response {
	acl.PrefabInit()
	return api.NewSuccessResponse("已重新生成权限初始化缓存")
}
