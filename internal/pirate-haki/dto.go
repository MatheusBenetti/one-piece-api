package piratehaki

type AssignHakiRequest struct {
	PirateID uint   `json:"pirate_id" binding:"required"`
	HakiID   uint   `json:"haki_id" binding:"required"`
	Level    string `json:"level" binding:"required"`
}

type AssignHakiResponse struct {
	PirateID uint   `json:"pirate_id"`
	HakiID   uint   `json:"haki_id"`
	Level    string `json:"level"`
}

type UpdateHakiLevelRequest struct {
	PirateID uint   `json:"pirate_id" binding:"required"`
	HakiID   uint   `json:"haki_id" binding:"required"`
	Level    string `json:"level" binding:"required"`
}

type UpdateHakiLevelResponse struct {
	PirateID uint   `json:"pirate_id"`
	HakiID   uint   `json:"haki_id"`
	Level    string `json:"level"`
}

type PirateHakiResponse struct {
	HakiID   uint   `json:"haki_id"`
	HakiName string `json:"haki_name"`
	Level    string `json:"level"`
}
