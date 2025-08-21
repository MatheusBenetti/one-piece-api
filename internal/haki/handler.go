package haki

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HakiHandler struct {
	service *HakiService
}

func NewHakiHandler(service *HakiService) *HakiHandler {
	return &HakiHandler{service: service}
}

func (h *HakiHandler) CreateHaki(c *gin.Context) {
	var req CreateHakiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	haki, err := h.service.CreateHaki(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := HakiResponse{
		ID:   int(haki.ID),
		Name: haki.Name,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *HakiHandler) GetAllHakis(c *gin.Context) {
	hakis, err := h.service.GetAllHakis()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch Akuma no Mi"})
		return
	}

	var responses []HakiResponse
	for _, haki := range hakis {
		responses = append(responses, HakiResponse{
			ID:   int(haki.ID),
			Name: haki.Name,
		})
	}

	c.JSON(http.StatusOK, responses)
}
