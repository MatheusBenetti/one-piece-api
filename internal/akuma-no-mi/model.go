package akumanomi

import "gorm.io/gorm"

type AkumaNoMi struct {
	gorm.Model
	Name        string `json:"name"`
	Style       string `json:"style"`
	Meaning     string `json:"meaning"`
	Description string `json:"description"`
	Sort        string `json:"sort"`
}
