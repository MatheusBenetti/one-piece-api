package haki

import "gorm.io/gorm"

type Haki struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
