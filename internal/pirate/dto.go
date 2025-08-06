package pirate

type CreatePirateRequest struct {
	Name        string   `json:"name" binding:"required,min=2,max=100"`
	AkumaNoMiID *uint    `json:"akuma_no_mi_id" binding:"omitempty,min=0"`
	Age         int16    `json:"age" binding:"required,min=0,max=1000"`
	Weapon      []string `json:"weapon" binding:"omitempty,dive,min=0"`
	Bounty      float64  `json:"bounty" binding:"required,min=0"`
	Rank        string   `json:"rank" binding:"required,oneof=Captain Commander Officer Crew"`
}

type PirateResponse struct {
	ID          uint
	Name        string
	AkumaNoMiID *uint
	Age         int16
	Weapon      []string
	Bounty      float64
	Rank        string
}
