package config

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type listRequest struct {
	ctx *boot.Context

	Page     int `validate:"required,numeric,min=1" json:"page"`
	PageSize int `validate:"omitempty,numeric,max=20" json:"page_size"`
}

type itemResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	K    string `json:"k"`
	V    string `json:"v"`
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
	var us []models.Config
	var c int64

	db().Order("id desc").Offset(offset).Limit(pageSize).Find(&us)
	db().Model(&models.Config{}).Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id:   v.Id,
			Name: v.Name,
			K:    v.K,
			V:    v.V,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
