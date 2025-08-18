package haki

type CreateHakiRequest struct {
	Name string `json:"name" binding:"required"`
}

type HakiResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
