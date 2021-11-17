package models

const (
	SessionTypeAuth    = 1
	SessionTypeCaptcha = 2
)

type Session struct {
	Id           uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	AdminerId    uint32
	Type         uint8
	SessionId    string
	SessionValue string
	Expire       uint32
	CreatedAt    uint32
	UpdatedAt    uint32
}

func (u *Session) TableName() string {
	return "goback_session"
}
