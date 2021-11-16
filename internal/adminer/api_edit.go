package adminer

import (
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/dblog"
	"github.com/rwcoding/goback/pkg/logger"
	"github.com/rwcoding/goback/pkg/util"
	"sort"
	"strconv"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,min=1"`
	//Username string `validate:"required,min=5,max=30"`
	Name     string `validate:"required,min=1,max=30"`
	Password string `validate:"omitempty,min=5,max=30"`
	Phone    string `validate:"required,numeric,max=11"`
	Roles    string `validate:"omitempty,min=0,max=200"`
	Status   uint8  `validate:"required,min=0,max=9"`
	IsSuper  uint8  `validate:"omitempty,min=0,max=1" json:"is_super"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {
	adminer := request.ctx.GetAdminer()
	ip := request.ctx.GetRemote()

	var u models.Adminer
	if db.Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的用户")
	}
	if adminer.Id != 1 && u.Id == 1 {
		return api.NewErrorResponse("您无权操作该用户")
	}
	if !adminer.Super() && (u.Super() || request.IsSuper == 1) {
		return api.NewErrorResponse("您无权操作该用户")
	}

	//验证角色
	rids := strings.Split(strings.TrimSpace(request.Roles), ",")
	var ridsSort []int
	repeat := map[string]bool{}
	for _, v := range rids {
		if repeat[v] || v == "" {
			continue
		}
		vv, err := strconv.Atoi(v)
		if err != nil {
			return api.NewErrorResponse("无效的角色")
		}
		repeat[v] = true
		ridsSort = append(ridsSort, vv)
		if _, ok := acl.GetSyncRoleMap().Load(uint32(vv)); !ok {
			return api.NewErrorResponse("未知的角色")
		}
	}
	sort.Ints(ridsSort)
	roles := ""
	for k, v := range ridsSort {
		roles += strconv.Itoa(v)
		if k < len(ridsSort)-1 {
			roles += ","
		}
	}

	pwd := strings.TrimSpace(request.Password)
	u.Name = strings.TrimSpace(request.Name)
	u.Phone = strings.TrimSpace(request.Phone)
	u.Roles = roles
	u.Status = request.Status
	u.IsSuper = request.IsSuper

	changePassword := false
	if pwd != "" && util.Password(pwd, u.Salt, false) != u.Password {
		u.Salt = util.RandString(32)
		u.Password = util.Password(request.Password, u.Salt, false)
		changePassword = true
	}

	err := db.Save(u).Error

	if err != nil {
		logger.Error(err.Error())
		return api.NewErrorResponse("修改失败")
	} else {
		//修改角色权限
		acl.GenerateByRoles(u.Roles)
		if changePassword {
			dblog.Add(&models.Log{
				AdminerId: adminer.Id,
				Type:      models.LogTypePassword,
				Msg:       adminer.Name + " 修改了 " + u.Name + " 的密码",
				Target:    u.Id,
				Ip:        ip,
			})
		}
	}

	return api.NewSuccessResponse("修改成功")
}
