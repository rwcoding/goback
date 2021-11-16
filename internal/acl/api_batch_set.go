package acl

import (
	"fmt"
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"strconv"
	"strings"
)

type batchSetRequest struct {
	ctx *boot.Context

	Type        int    `validate:"required,numeric,min=1,max=2"`
	Roles       string `validate:"required,min=1,max=10000"`
	Permissions string `validate:"required,min=1,max=10000"`
}

func NewApiBatchSet(ctx *boot.Context) boot.Logic {
	return &batchSetRequest{ctx: ctx}
}

func (request *batchSetRequest) Run() *api.Response {
	isAdd := request.Type == 1

	var ids []uint32
	ps := strings.Split(request.Permissions, ",")
	for _, v := range ps {
		id, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		ids = append(ids, uint32(id))
	}
	var permissions []string
	var pList []models.Permission
	if db.Where("id IN ?", ids).Find(&pList).RowsAffected > 0 {
		for _, v := range pList {
			permissions = append(permissions, v.Permission)
		}
	}

	roleIsList := strings.Split(request.Roles, ",")
	fmt.Println(permissions)
	fmt.Println(roleIsList)
	for _, v := range roleIsList {
		roleId, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		if db.Take(&models.Role{}, roleId).Error != nil {
			return api.NewErrorResponse("无效的角色" + v)
		}
		if isAdd {
			var addList []models.RolePermission
			var rps []models.RolePermission
			db.Where("role_id=?", roleId).Find(&rps)
			for _, vv := range pList {
				exists := false
				for _, vvv := range rps {
					if vvv.Permission == vv.Permission {
						exists = true
					}
				}
				if !exists {
					addList = append(addList, models.RolePermission{
						RoleId:     uint32(roleId),
						Permission: vv.Permission,
					})
				}
			}
			fmt.Println(addList)
			if len(addList) > 0 {
				db.Create(&addList)
			}
		} else {
			db.Where("role_id=?", roleId).Where("permission IN ?", permissions).
				Delete(&models.RolePermission{})
		}

		acl.Generate(uint32(roleId))
	}

	return api.NewSuccessResponse("设置成功")
}
