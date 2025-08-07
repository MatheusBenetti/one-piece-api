package api

import (
	akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	"github.com/MatheusBenetti/one-piece-api/internal/pirate"
	"github.com/MatheusBenetti/one-piece-api/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(akumaNoMiHandler *akumanomi.AkumaNoMiHandler, pirateHandler *pirate.PirateHandler, userHandler *user.UserHandler) *gin.Engine {
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

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.Login)
	}

	return router
}
