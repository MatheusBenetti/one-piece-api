package akumanomi

type CreateAkumaNoMiRequest struct {
	Name        string
	Model       string
	Meaning     string
	Description string
	Sort        string
}

type AkumaNoMiResponse struct {
	ID          uint
	Name        string
	Model       string
	Meaning     string
	Description string
	Sort        string
}
