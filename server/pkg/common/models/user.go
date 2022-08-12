package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;default:null"`
	Email    string `json:"email" gorm:"unique;not null;default:null"`
	Password string `json:"password,omitempty" gorm:"not null;default:null"`
	Names    string `json:"names"`
}
