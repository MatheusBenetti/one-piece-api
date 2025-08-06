package pirate

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PirateHandler struct {
	service *PirateService
}

func NewHandler(service *PirateService) *PirateHandler {
	return &PirateHandler{service: service}
}

func (h *PirateHandler) AddPirate(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"endpoint": "AddPirate",
		"method":   "POST",
	}).Info("Starting to create a pirate")

	var req CreatePirateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"error":   err.Error(),
			"request": req,
		}).Error("Error validating data")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pirate, err := h.service.AddPirate(req.Name, req.AkumaNoMiID, req.Age, req.Weapon, req.Bounty, req.Rank)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := PirateResponse{
		ID:          pirate.ID,
		Name:        pirate.Name,
		AkumaNoMiID: pirate.AkumaNoMiID,
		Age:         pirate.Age,
		Weapon:      pirate.Weapon,
		Bounty:      pirate.Bounty,
		Rank:        pirate.Rank,
	}

	logrus.WithFields(logrus.Fields{
		"pirate_name": req.Name,
	}).Info("Pirata created")

	c.JSON(http.StatusCreated, response)
}

func (h *PirateHandler) FindByName(c *gin.Context) {
	name := c.Param("name")

	pirate, err := h.service.FindByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid pirate name"})
		return
	}

	response := PirateResponse{
		ID:          pirate.ID,
		Name:        pirate.Name,
		AkumaNoMiID: pirate.AkumaNoMiID,
		Age:         pirate.Age,
		Weapon:      pirate.Weapon,
		Bounty:      pirate.Bounty,
		Rank:        pirate.Rank,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PirateHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	pirate, err := h.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid pirate"})
		return
	}

	response := PirateResponse{
		ID:          pirate.ID,
		Name:        pirate.Name,
		AkumaNoMiID: pirate.AkumaNoMiID,
		Age:         pirate.Age,
		Weapon:      pirate.Weapon,
		Bounty:      pirate.Bounty,
		Rank:        pirate.Rank,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PirateHandler) GetAllPirates(c *gin.Context) {
	pirates, err := h.service.GetAllPirates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch Pirates"})
		return
	}

	var responses []PirateResponse
	for _, pirate := range pirates {
		responses = append(responses, PirateResponse{
			ID:          pirate.ID,
			Name:        pirate.Name,
			AkumaNoMiID: pirate.AkumaNoMiID,
			Age:         pirate.Age,
			Weapon:      pirate.Weapon,
			Bounty:      pirate.Bounty,
			Rank:        pirate.Rank,
		})
	}

	c.JSON(http.StatusOK, responses)
}
