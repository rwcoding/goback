package adminer

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=20" json:"page_size"`
	Username string `validate:"omitempty,max=30" json:"username"`
}

type itemResponse struct {
	Id        uint32 `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Roles     string `json:"roles"`
	Status    uint8  `json:"status"`
	IsSuper   uint8  `json:"is_super" json:"is_super"`
	CreatedAt uint32 `json:"created_at"`
}

type listResponse struct {
	Datas       []itemResponse    `json:"datas"`
	Count       int64             `json:"count"`
	Page        int               `json:"page"`
	PageSize    int               `json:"page_size"`
	StatusNames map[uint8]string  `json:"status_names"`
	RoleNames   map[uint32]string `json:"role_names"`
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
	var us []models.Adminer
	var c int64

	tx1 := db.Model(&models.Adminer{}).Order("id").Offset(offset).Limit(pageSize)
	tx2 := db.Model(&models.Adminer{})
	if request.Username != "" {
		tx1.Where("username=?", request.Username)
		tx2.Where("username=?", request.Username)
	}
	tx1.Find(&us)
	tx2.Count(&c)

	list := []itemResponse{}
	for _, v := range us {
		list = append(list, itemResponse{
			Id:        v.Id,
			Username:  v.Username,
			Name:      v.Name,
			Phone:     v.Phone,
			Roles:     v.Roles,
			Status:    v.Status,
			IsSuper:   v.IsSuper,
			CreatedAt: v.CreatedAt,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:       list,
		Count:       c,
		Page:        request.Page,
		PageSize:    pageSize,
		StatusNames: models.AdminerStatusNames(),
		RoleNames:   acl.GetRoleMap(),
	})
}
