package models

import "gorm.io/plugin/soft_delete"

type Role struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (m *Role) TableName() string {
	return "gobui_role"
}
