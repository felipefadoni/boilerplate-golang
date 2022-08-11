package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string `gorm:"type:varchar(200);not null"`
	Email    string `gorm:"type:varchar(200);not null"`
	Login    string `gorm:"type:varchar(200);not null;unique"`
	Password string `gorm:"type:varchar(200);not null"`
	Accepted bool   `gorm:"type:boolean;not null"`
}
