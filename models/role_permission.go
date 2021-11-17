package models

type RolePermission struct {
	Id         uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	RoleId     uint32
	Permission string
	CreatedAt  uint32
	UpdatedAt  uint32
}

func (m *RolePermission) TableName() string {
	return "goback_role_permission"
}

//func (m *RolePermission) AfterSave(tx *gorm.DB) (err error) {
//	var ps []map[string]string
//	result := db().Model(m).Where("role_id", m.Id).Select("permission").Find(&ps)
//
//	var rp strings.Builder
//	if result.RowsAffected > 0 {
//		rp.WriteString(",")
//		for _, v := range ps {
//			rp.WriteString(v["permission"])
//			rp.WriteString(",")
//		}
//	}
//
//	permissions := rp.String()
//	sign := "role-" + strconv.Itoa(int(m.RoleId))
//	cache := &Cache{}
//	db().Where("sign=?", sign).First(cache)
//	if cache.Id == 0 {
//		cache.Sign = sign
//		cache.Data = permissions
//		db().Create(cache)
//	} else if cache.Data != permissions {
//		cache.Data = permissions
//		db().Save(cache)
//	}
//	return
//}
