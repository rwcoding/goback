package profile

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type infoRequest struct {
	ctx *boot.Context
}

type infoResponse struct {
	Username    string            `json:"username"`
	Name        string            `json:"name"`
	Phone       string            `json:"phone"`
	Roles       string            `json:"roles"`
	Status      uint8             `json:"status"`
	IsSuper     uint8             `json:"is_super"`
	CreatedAt   uint32            `json:"created_at"`
	StatusNames map[uint8]string  `json:"status_names"`
	RoleNames   map[uint32]string `json:"role_names"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (r *infoRequest) Run() *api.Response {
	adminer := r.ctx.GetAdminer()
	return api.NewDataResponse(&infoResponse{
		Username:    adminer.Username,
		Name:        adminer.Name,
		Phone:       adminer.Phone,
		Roles:       adminer.Roles,
		Status:      adminer.Status,
		IsSuper:     adminer.IsSuper,
		CreatedAt:   adminer.CreatedAt,
		StatusNames: models.AdminerStatusNames(),
		RoleNames:   acl.GetRoleMap(),
	})
}
