package models

const (
	PermissionTypeApi uint8 = 1
	PermissionTypeDef uint8 = 2
)

func PermissionTypeNames() map[uint8]string {
	return map[uint8]string{
		PermissionTypeApi: "API",
		PermissionTypeDef: "自定义",
	}
}

type Permission struct {
	Id         uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string
	Permission string
	Gid        uint32
	Type       uint8
	CreatedAt  uint32
	UpdatedAt  uint32
}

func (m *Permission) TableName() string {
	return "goback_permission"
}
