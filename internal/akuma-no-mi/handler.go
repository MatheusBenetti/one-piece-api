package akumanomi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AkumaNoMiHandler struct {
	service *AkumaNoMiService
}

func NewHandler(service *AkumaNoMiService) *AkumaNoMiHandler {
	return &AkumaNoMiHandler{service: service}
}

func (h *AkumaNoMiHandler) AddAkumaNoMi(c *gin.Context) {
	var req CreateAkumaNoMiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	akumaNoMi, err := h.service.AddAkumaNoMi(req.Name, req.Style, req.Meaning, req.Description, req.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := AkumaNoMiResponse{
		ID:          uint(akumaNoMi.ID),
		Name:        akumaNoMi.Name,
		Style:       akumaNoMi.Style,
		Meaning:     akumaNoMi.Meaning,
		Description: akumaNoMi.Description,
		Sort:        akumaNoMi.Sort,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AkumaNoMiHandler) FindByName(c *gin.Context) {
	name := c.Param("name")

	akumaNoMi, err := h.service.FindByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid fruit"})
		return
	}

	response := AkumaNoMiResponse{
		ID:          akumaNoMi.ID,
		Name:        akumaNoMi.Name,
		Style:       akumaNoMi.Style,
		Meaning:     akumaNoMi.Meaning,
		Description: akumaNoMi.Description,
		Sort:        akumaNoMi.Sort,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AkumaNoMiHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	akumaNoMi, err := h.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid fruit"})
		return
	}

	response := AkumaNoMiResponse{
		ID:          akumaNoMi.ID,
		Name:        akumaNoMi.Name,
		Style:       akumaNoMi.Style,
		Meaning:     akumaNoMi.Meaning,
		Description: akumaNoMi.Description,
		Sort:        akumaNoMi.Sort,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AkumaNoMiHandler) GetAllAkumaNoMi(c *gin.Context) {
	fruits, err := h.service.GetAllAkumaNoMi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch Akuma no Mi"})
		return
	}

	var responses []AkumaNoMiResponse
	for _, akumaNoMi := range fruits {
		responses = append(responses, AkumaNoMiResponse{
			ID:          akumaNoMi.ID,
			Name:        akumaNoMi.Name,
			Style:       akumaNoMi.Style,
			Meaning:     akumaNoMi.Meaning,
			Description: akumaNoMi.Description,
			Sort:        akumaNoMi.Sort,
		})
	}

	c.JSON(http.StatusOK, responses)
}
