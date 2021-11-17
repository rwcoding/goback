package models

type PermissionGroup struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Pid       uint32
	Name      string
	Ord       uint32
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *PermissionGroup) TableName() string {
	return "goback_permission_group"
}
