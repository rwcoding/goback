package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type listRequest struct {
	ctx *boot.Context

	Page      int    `validate:"required,numeric,min=1" json:"page"`
	PageSize  int    `validate:"required,numeric,max=20" json:"page_size"`
	AdminerId uint32 `validate:"omitempty,numeric,min=0" json:"adminer_id"`
}

type itemResponse struct {
	Id        uint64 `json:"id"`
	AdminerId uint32 `json:"adminer_id"`
	SessionId string `json:"session_id"`
	Type      uint8  `json:"type"`
	Expire    uint32 `json:"expire"`
	CreatedAt uint32 `json:"created_at"`
	UpdatedAt uint32 `json:"updated_at"`
}

type listResponse struct {
	Datas    []itemResponse `json:"datas"`
	Count    int64          `json:"count"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

func NewApiList(ctx *boot.Context) boot.Logic {
	return &listRequest{ctx: ctx}
}

func (request *listRequest) Run() *api.Response {
	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var data []models.Session
	var c int64

	tx1 := db().Model(&models.Session{}).Order("id").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Session{})
	if request.AdminerId > 0 {
		tx1.Where("adminer_id=?", request.AdminerId)
		tx2.Where("adminer_id=?", request.AdminerId)
	}
	tx1.Find(&data)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range data {
		list = append(list, itemResponse{
			Id:        v.Id,
			AdminerId: v.AdminerId,
			SessionId: v.SessionId[:8] + "......",
			Type:      v.Type,
			Expire:    v.Expire,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.CreatedAt,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
