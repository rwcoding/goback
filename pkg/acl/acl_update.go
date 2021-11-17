package acl

import (
	"github.com/rwcoding/goback/models"
	"strconv"
	"strings"
)

// Generate 生成所有关联某个角色的权限
// 修改角色关联权限后使用该方法更新
func Generate(roleId uint32) {
	var users []models.Adminer
	if db().Where("FIND_IN_SET("+strconv.Itoa(int(roleId))+", roles)").
		Find(&users).RowsAffected == 0 {
		return
	}

	exist := map[string]bool{}
	for _, v := range users {
		if v.Roles == "" {
			continue
		}
		if exist[v.Roles] {
			continue
		}

		GenerateByRoles(v.Roles)

		exist[v.Roles] = true
	}
}

// GenerateByRoles 生成角色集合的权限列表
// 修改了用户的角色信息后，需要调用该方法更新权限
func GenerateByRoles(roles string) string {
	if roles == "" {
		return ""
	}
	arr := strings.Split(roles, ",")
	var permissions []struct{ Permission string }
	db().Model(&models.RolePermission{}).
		Where("role_id IN ?", arr).
		Select("permission").Find(&permissions)

	sb := strings.Builder{}
	et := map[string]bool{}
	size := len(permissions)
	for k, v := range permissions {
		if et[v.Permission] {
			continue
		}
		et[v.Permission] = true
		sb.WriteString(v.Permission)
		if k < size-1 {
			sb.WriteString(",")
		}
	}
	key := "acl-" + strings.ReplaceAll(roles, ",", "-")

	//存入数据库
	cache := &models.Cache{}
	db().Where("sign=?", key).First(cache)
	if cache.Id > 0 {
		cache.Data = sb.String()
		db().Save(cache)
	} else {
		cache.Sign = key
		cache.Data = sb.String()
		db().Create(cache)
	}
	return sb.String()
}

func PrefabInit() {
	var us []struct{ Roles string }
	db().Model(&models.Adminer{}).Select("roles").Unscoped().Group("roles").Find(&us)
	if len(us) > 0 {
		for _, v := range us {
			GenerateByRoles(v.Roles)
		}
	}
}
