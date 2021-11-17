package acl

import (
	"github.com/rwcoding/goback/models"
	"strings"
)

var db = models.GetDb

func IsDefault(api string) bool {
	return api[:12] == "goback.login"
}

func Verify(adminer *models.Adminer, api string) bool {
	if adminer.Super() {
		return true
	}
	if adminer.Roles == "" {
		return false
	}
	key := strings.ReplaceAll(adminer.Roles, ",", "-")
	cache := models.Cache{}
	db().Where("sign=?", "acl-"+key).Take(cache)
	if cache.Id == 0 {
		cache.Data = GenerateByRoles(adminer.Roles)
	}
	if strings.Contains(","+cache.Data+",", ","+api+",") {
		return true
	}
	return false
}
