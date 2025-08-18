package api

import (
	akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	haki "github.com/MatheusBenetti/one-piece-api/internal/haki"
	pirate "github.com/MatheusBenetti/one-piece-api/internal/pirate"
	pirateHaki "github.com/MatheusBenetti/one-piece-api/internal/pirate-haki"
	user "github.com/MatheusBenetti/one-piece-api/internal/user"
	auth "github.com/MatheusBenetti/one-piece-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	akumaNoMiHandler *akumanomi.AkumaNoMiHandler,
	hakiHandler *haki.HakiHandler,
	pirateHandler *pirate.PirateHandler,
	pirateHakiHandler *pirateHaki.PirateHakiHandler,
	userHandler *user.UserHandler) *gin.Engine {
	router := gin.Default()

	akumaNoMiGroup := router.Group("/akumanomi")
	{
		akumaNoMiGroup.POST("/", akumaNoMiHandler.AddAkumaNoMi, auth.AuthMiddleware())
		akumaNoMiGroup.GET("/:name", akumaNoMiHandler.FindByName)
		akumaNoMiGroup.GET("/id/:id", akumaNoMiHandler.FindById)
		akumaNoMiGroup.GET("/", akumaNoMiHandler.GetAllAkumaNoMi)
	}

	hakiGroup := router.Group("/haki")
	{
		hakiGroup.POST("/", hakiHandler.CreateHaki, auth.AuthMiddleware())
		hakiGroup.GET("/", hakiHandler.GetAllHakis)
	}

	piratesGroup := router.Group("/pirate")
	{
		piratesGroup.POST("/", pirateHandler.AddPirate, auth.AuthMiddleware())
		piratesGroup.GET("/:name", pirateHandler.FindByName)
		piratesGroup.GET("/id/:id", pirateHandler.FindById)
		piratesGroup.GET("/", pirateHandler.GetAllPirates)
		piratesGroup.GET("/:id/hakis", pirateHakiHandler.GetHakisByPirateId)
	}

	pirateHakiGroup := router.Group("/pirate-hakis")
	{
		pirateHakiGroup.POST("/assign", pirateHakiHandler.AssignHakiToPirate, auth.AuthMiddleware())
		pirateHakiGroup.PUT("/update-level", pirateHakiHandler.UpdateHakiLevel, auth.AuthMiddleware())
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.Login)
	}

	return router
}
