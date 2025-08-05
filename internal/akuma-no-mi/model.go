package akumanomi

type AkumaNoMi struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Model       string `json:"model"`
	Meaning     string `json:"meaning"`
	Description string `json:"description"`
	Sort        string `json:"sort"`
}
