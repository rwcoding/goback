package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type groupListRequest struct {
	ctx *boot.Context

	Page     int `validate:"required,numeric,min=1" json:"page"`
	PageSize int `validate:"required,numeric,max=20" json:"page_size"`
}

type groupItemResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Ord  uint32 `json:"ord"`
}

type groupListResponse struct {
	Datas    []groupItemResponse `json:"datas"`
	Count    int64               `json:"count"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

func NewApiGroupList(ctx *boot.Context) boot.Logic {
	return &groupListRequest{ctx: ctx}
}

func (request *groupListRequest) Run() *api.Response {
	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var us []models.PermissionGroup
	var c int64

	db.Model(&models.PermissionGroup{}).
		Order("ord").Offset(offset).Limit(pageSize).Find(&us)
	db.Model(&models.PermissionGroup{}).Count(&c)

	list := []groupItemResponse{}
	for _, v := range us {
		list = append(list, groupItemResponse{
			Id:   v.Id,
			Name: v.Name,
			Ord:  v.Ord,
		})
	}

	return api.NewDataResponse(&groupListResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
