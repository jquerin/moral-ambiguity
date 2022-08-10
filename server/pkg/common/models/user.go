package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique_index;not null"`
	Email    string `json:"email" gorm:"unique_index;not null"`
	Password string `json:"password" gorm:"not null"`
	Names    string `json:"names"`
}
