package models

type Config struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	K         string
	V         string
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *Config) TableName() string {
	return "gobui_config"
}
