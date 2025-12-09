package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;size:50"`
	Password string `gorm:"not null;size:255"`
}
