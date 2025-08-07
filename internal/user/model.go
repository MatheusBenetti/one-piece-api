package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `gorm:"not null"`
	// Role     string `json:"role" gorm:"default:'player'"`
}
