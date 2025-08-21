package piratehaki

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PirateHakiHandler struct {
	service *PirateHakiService
}

func NewPirateHakiHandler(service *PirateHakiService) *PirateHakiHandler {
	return &PirateHakiHandler{service: service}
}

func (h *PirateHakiHandler) AssignHakiToPirate(c *gin.Context) {
	var req AssignHakiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignHakiToPirate(req.PirateID, req.HakiID, req.Level); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Haki assigned to pirate"})
}

func (h *PirateHakiHandler) UpdateHakiLevel(c *gin.Context) {
	var req UpdateHakiLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateHakiLevel(req.PirateID, req.HakiID, req.Level); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Haki level updated"})
}

func (h *PirateHakiHandler) GetHakisByPirateId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	hakis, err := h.service.GetHakisByPirateID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hakis)
}
