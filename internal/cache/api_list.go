package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"required,numeric,max=20" json:"page_size"`
	Sign     string `validate:"omitempty,max=30" json:"sign"`
}

type itemResponse struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
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
	var data []models.Cache
	var c int64

	tx1 := db().Model(&models.Cache{}).Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Cache{})
	if request.Sign != "" {
		tx1.Where("sign=?", request.Sign)
		tx2.Where("sign=?", request.Sign)
	}
	tx1.Find(&data)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range data {
		list = append(list, itemResponse{
			Id:        v.Id,
			Sign:      v.Sign,
			Name:      v.Name,
			Data:      v.Data,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
