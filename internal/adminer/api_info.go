package adminer

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1"`
}

type infoResponse struct {
	Id          uint32            `json:"id"`
	Username    string            `json:"username"`
	Name        string            `json:"name"`
	Phone       string            `json:"phone"`
	Roles       string            `json:"roles"`
	Status      uint8             `json:"status"`
	IsSuper     uint8             `json:"is_super"`
	StatusNames map[uint8]string  `json:"status_names"`
	RoleNames   map[uint32]string `json:"role_names"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var user models.Adminer
	if db().Take(&user, request.Id).Error != nil {
		return api.NewErrorResponse("无效的用户")
	}
	return api.NewDataResponse(&infoResponse{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Phone:       user.Phone,
		Roles:       user.Roles,
		Status:      user.Status,
		IsSuper:     user.IsSuper,
		StatusNames: models.AdminerStatusNames(),
		RoleNames:   acl.GetRoleMap(),
	})
}
