package pirate

type CreatePirateRequest struct {
	Name        string   `json:"name" binding:"required"`
	AkumaNoMiID *uint    `json:"akuma_no_mi_id"`
	Age         int16    `json:"age"`
	Weapon      []string `json:"weapon"`
	Bounty      float64  `json:"bounty"`
	Rank        string   `json:"rank"`
}

type PirateResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	AkumaNoMiID *uint    `json:"akuma_no_mi_id"`
	Age         int16    `json:"age"`
	Weapon      []string `json:"weapon"`
	Bounty      float64  `json:"bounty"`
	Rank        string   `json:"rank"`
}
