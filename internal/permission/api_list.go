package permission

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type listRequest struct {
	ctx *boot.Context

	Page       int    `validate:"required,numeric,min=1" json:"page"`
	PageSize   int    `validate:"required,numeric,max=20" json:"page_size"`
	Permission string `validate:"omitempty,max=100" json:"permission"`
}

type itemResponse struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	Gid        uint32 `json:"gid"`
	Type       uint8  `json:"type"`
}

type listResponse struct {
	Datas      []itemResponse    `json:"datas"`
	Count      int64             `json:"count"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	GroupNames map[uint32]string `json:"group_names"`
	TypeNames  map[uint8]string  `json:"type_names"`
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
	var us []models.Permission
	var c int64

	tx1 := db().Model(&models.Permission{}).Order("id").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Permission{})
	if request.Permission != "" {
		tx1.Where("permission LIKE ?", "%"+request.Permission+"%")
		tx2.Where("permission LIKE ?", "%"+request.Permission+"%")
	}
	tx1.Find(&us)
	tx2.Count(&c)

	list := []itemResponse{}
	for _, v := range us {
		list = append(list, itemResponse{
			Id:         v.Id,
			Name:       v.Name,
			Gid:        v.Gid,
			Permission: v.Permission,
			Type:       v.Type,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:      list,
		Count:      c,
		Page:       request.Page,
		PageSize:   pageSize,
		GroupNames: groupNames(),
		TypeNames:  models.PermissionTypeNames(),
	})
}

func groupNames() map[uint32]string {
	ret := map[uint32]string{}
	ret[0] = "未定义"
	var data []models.PermissionGroup
	if db().Find(&data).RowsAffected == 0 {
		return nil
	}
	for _, v := range data {
		ret[v.Id] = v.Name
	}
	return ret
}
