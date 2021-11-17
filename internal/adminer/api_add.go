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

type addRequest struct {
	ctx *boot.Context

	Username string `validate:"required,min=5,max=30"`
	Name     string `validate:"required,min=1,max=30"`
	Password string `validate:"required,min=5,max=30"`
	Phone    string `validate:"required,numeric,max=11"`
	Roles    string `validate:"omitempty,min=0,max=200"`
	Status   uint8  `validate:"required,min=0,max=9"`
	IsSuper  uint8  `validate:"omitempty,min=0,max=1" json:"is_super"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	adminer := request.ctx.GetAdminer()
	ip := request.ctx.GetRemote()

	if !adminer.Super() && request.IsSuper == 1 {
		return api.NewErrorResponse("您不能添加超级管理员")
	}

	if models.HasAdminer(request.Username) {
		return api.NewErrorResponse("账号已经存在")
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

	salt := util.RandString(32)
	pwd := util.Password(request.Password, salt, false)
	newUser := models.Adminer{
		Username: strings.TrimSpace(request.Username),
		Password: pwd,
		Salt:     salt,
		Name:     strings.TrimSpace(request.Name),
		Phone:    strings.TrimSpace(request.Phone),
		Roles:    roles,
		Status:   request.Status,
		IsSuper:  request.IsSuper,
	}
	ret := db().Create(&newUser)
	if ret.Error == nil {
		dblog.Add(&models.Log{
			AdminerId: adminer.Id,
			Type:      models.LogTypeAddAdminer,
			Msg:       adminer.Name + " 添加管理员 " + newUser.Name,
			Target:    newUser.Id,
			Ip:        ip,
		})
		return api.NewMDResponse("添加成功", &addResponse{
			Id: newUser.Id,
		})
	} else {
		logger.Error(ret.Error.Error())
	}
	return api.NewErrorResponse("添加失败")
}
