package models

const (
	LogTypeLogin         = 1
	LogTypePassword      = 2
	LogTypeAddAdminer    = 3
	LogTypeEditAdminer   = 4
	LogTypeDeleteAdminer = 5
)

type Log struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	AdminerId uint32
	Type      uint32
	Msg       string
	Details   string
	Ip        string
	Target    uint32
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *Log) TableName() string {
	return "goback_log"
}

func LogTypeNames() map[int]string {
	return map[int]string{
		LogTypeLogin:         "登录",
		LogTypePassword:      "修改密码",
		LogTypeAddAdminer:    "增加管理",
		LogTypeEditAdminer:   "编辑管理",
		LogTypeDeleteAdminer: "删除管理",
	}
}
