package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1"`
}

type infoResponse struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
	UpdatedAt uint32 `json:"updated_at"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	cache := models.Cache{}
	if db.Take(&cache, request.Id).Error != nil {
		return api.NewErrorResponse("无效的缓存")
	}

	return api.NewDataResponse(&infoResponse{
		Id:        cache.Id,
		Name:      cache.Name,
		Sign:      cache.Sign,
		Data:      cache.Data,
		UpdatedAt: cache.UpdatedAt,
	})
}
