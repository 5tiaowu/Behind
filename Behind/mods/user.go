package mods

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(100;not null;unique"`
	Password  string `gorm:"size:255;"`
}
