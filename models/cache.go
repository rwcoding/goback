package models

type Cache struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Sign      string
	Data      string
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *Cache) TableName() string {
	return "gobui_cache"
}
