package api

import (
	akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	"github.com/MatheusBenetti/one-piece-api/internal/pirate"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(akumaNoMiHandler *akumanomi.AkumaNoMiHandler, pirateHandler *pirate.PirateHandler) *gin.Engine {
	router := gin.Default()

	akumaNoMiGroup := router.Group("/akumanomi")
	{
		akumaNoMiGroup.POST("/", akumaNoMiHandler.AddAkumaNoMi)
		akumaNoMiGroup.GET("/:name", akumaNoMiHandler.FindByName)
		akumaNoMiGroup.GET("/id/:id", akumaNoMiHandler.FindById)
		akumaNoMiGroup.GET("/", akumaNoMiHandler.GetAllAkumaNoMi)
	}

	piratesGroup := router.Group("/pirate")
	{
		piratesGroup.POST("/", pirateHandler.AddPirate)
		piratesGroup.GET("/:name", pirateHandler.FindByName)
		piratesGroup.GET("/id/:id", pirateHandler.FindById)
		piratesGroup.GET("/", pirateHandler.GetAllPirates)
	}

	return router
}
