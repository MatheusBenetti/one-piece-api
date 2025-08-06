package pirate

import akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"

type CreatePirateRequest struct {
	Name        string
	AkumaNoMiID *uint
	AkumaNoMi   *akumanomi.AkumaNoMi
	Age         int16
	Weapon      []string
	Bounty      float64
	Rank        string
}

type PirateResponse struct {
	ID          uint
	Name        string
	AkumaNoMiID *uint
	AkumaNoMi   *akumanomi.AkumaNoMi
	Age         int16
	Weapon      []string
	Bounty      float64
	Rank        string
}
