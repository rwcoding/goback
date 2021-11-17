package acl

import (
	"github.com/rwcoding/goback/models"
	"sync"
)

var rolesMap *sync.Map = &sync.Map{}

func GetSyncRoleMap() *sync.Map {
	return rolesMap
}

func GetRoleMap() map[uint32]string {
	return GenerateRoleMap()
}

func GenerateRoleMap() map[uint32]string {
	ret := map[uint32]string{}
	var roles []models.Role
	db().Find(&roles)
	if len(roles) == 0 {
		return nil
	}

	for _, v := range roles {
		ret[v.Id] = v.Name
	}
	return ret
}
