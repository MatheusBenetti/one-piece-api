package akumanomi

type CreateAkumaNoMiRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Model       string `json:"model" binding:"required,min=2,max=100"`
	Meaning     string `json:"meaning" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"required,min=2,max=100"`
	Sort        string `json:"sort" binding:"required,oneof=Zoan Mythical-Zoan Paramecia Logia None"`
}

type AkumaNoMiResponse struct {
	ID          uint
	Name        string
	Model       string
	Meaning     string
	Description string
	Sort        string
}
