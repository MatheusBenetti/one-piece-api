package api

import (
	akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(akumaNoMiHandler *akumanomi.AkumaNoMiHandler) *gin.Engine {
	router := gin.Default()

	akumaNoMiGroup := router.Group("/akumanomi")
	{
		akumaNoMiGroup.POST("/", akumaNoMiHandler.AddAkumaNoMi)
		akumaNoMiGroup.GET("/:name", akumaNoMiHandler.FindByName)
		akumaNoMiGroup.GET("/id/:id", akumaNoMiHandler.FindById)
		akumaNoMiGroup.GET("/", akumaNoMiHandler.GetAllAkumaNoMi)
	}

	return router
}
