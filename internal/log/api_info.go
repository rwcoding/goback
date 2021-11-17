package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

type infoResponse struct {
	Id        uint64 `json:"id"`
	AdminerId uint32 `json:"adminer_id"`
	Type      uint32 `json:"type"`
	Msg       string `json:"msg"`
	Details   string `json:"details"`
	Target    uint32 `json:"target"`
	CreatedAt uint32 `json:"created_at"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	log := models.Log{}
	if db().Take(&log, request.Id).Error != nil {
		return api.NewErrorResponse("无效的日志")
	}

	return api.NewDataResponse(&infoResponse{
		Id:        log.Id,
		AdminerId: log.AdminerId,
		Type:      log.Type,
		Msg:       log.Msg,
		Details:   log.Details,
		Target:    log.Target,
		CreatedAt: log.CreatedAt,
	})
}
