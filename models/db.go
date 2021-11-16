package models

import "gorm.io/gorm"

var db *gorm.DB

func SetDb(g *gorm.DB) {
	db = g
}

func GetDb() *gorm.DB {
	return db
}
