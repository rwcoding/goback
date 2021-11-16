package cache

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/util"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"required,numeric,max=20" json:"page_size"`
	Type     int    `validate:"omitempty,numeric,min=0" json:"type"`
	Start    string `validate:"omitempty,len=10" json:"start"`
	End      string `validate:"omitempty,len=10" json:"end"`
}

type itemResponse struct {
	Id        uint64 `json:"id"`
	AdminerId uint32 `json:"adminer_id"`
	Type      uint32 `json:"type"`
	Msg       string `json:"msg"`
	Details   string `json:"details"`
	Ip        string `json:"ip"`
	Target    uint32 `json:"target"`
	CreatedAt uint32 `json:"created_at"`
}

type listResponse struct {
	Datas     []itemResponse `json:"datas"`
	Count     int64          `json:"count"`
	Page      int            `json:"page"`
	PageSize  int            `json:"page_size"`
	TypeNames map[int]string `json:"type_names"`
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
	var data []models.Log
	var c int64

	tx1 := db.Model(&models.Log{}).Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db.Model(&models.Log{})
	if request.Type != 0 {
		tx1.Where("type=?", request.Type)
		tx2.Where("type=?", request.Type)
	}
	if request.Start != "" {
		start, err := util.StringToUnix(request.Start)
		if err != nil {
			return api.NewErrorResponse(err.Error())
		}
		tx1.Where("created_at>=?", start)
		tx2.Where("created_at>=?", start)
	}
	if request.End != "" {
		end, err := util.StringToUnix(request.End)
		if err != nil {
			return api.NewErrorResponse(err.Error())
		}
		tx1.Where("created_at<?", end+86400)
		tx2.Where("created_at<?", end+86400)
	}
	tx1.Find(&data)
	tx2.Count(&c)

	list := []itemResponse{}
	for _, v := range data {
		list = append(list, itemResponse{
			Id:        v.Id,
			AdminerId: v.AdminerId,
			Type:      v.Type,
			Msg:       v.Msg,
			Details:   v.Details,
			Ip:        v.Ip,
			Target:    v.Target,
			CreatedAt: v.CreatedAt,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:     list,
		Count:     c,
		Page:      request.Page,
		PageSize:  pageSize,
		TypeNames: models.LogTypeNames(),
	})
}
