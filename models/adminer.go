package models

import "gorm.io/plugin/soft_delete"

const AdminerStatusOk = 1
const AdminerStatusLock = 2

func AdminerStatusNames() map[uint8]string {
	return map[uint8]string{
		AdminerStatusOk:   "正常",
		AdminerStatusLock: "锁定",
	}
}

type Adminer struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Username  string
	Password  string
	Salt      string
	Name      string
	Phone     string
	Roles     string
	Status    uint8
	IsSuper   uint8
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (u *Adminer) TableName() string {
	return "gobui_adminer"
}

func (u *Adminer) IsOK() bool {
	return u.Id > 0 && u.Status == AdminerStatusOk
}

func (u *Adminer) Super() bool {
	return u.Id > 0 && (u.IsSuper == 1 || u.Id == 1)
}

func HasAdminer(username string) bool {
	var c int64
	if db.Model(&Adminer{}).Where("username=?", username).
		Unscoped().
		Count(&c); c > 0 {
		return true
	}
	return false
}
