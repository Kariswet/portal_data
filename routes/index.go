package routes

import (
	"portal-data-kalsel/controller"
	"portal-data-kalsel/docs"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))

	artikelGuard := router.Group("/publikasi/artikel")
	{
		artikelGuard.POST("/create", controller.InsertArtikelController)
		artikelGuard.PUT("/update/:id", controller.UpdateArtikelController)
		artikelGuard.DELETE("/delete/:id", controller.DeleteArtikelController)
	}

	artikel := router.Group("/publikasi/artikel")
	{
		artikel.GET("/get-one/:id", controller.GetArtikelController)
		artikel.POST("get-all", controller.GetAllArtikelController)
	}

	infografisGuard := router.Group("/publikasi/infografis")
	{
		infografisGuard.POST("/create", controller.InsertInfografisController)
		infografisGuard.PUT("/update/:id", controller.UpdateInfografisController)
		infografisGuard.DELETE("/delete/:id", controller.DeleteInfografisController)
	}

	infografis := router.Group("/publikasi/infografis")
	{
		infografis.GET("/get-one/:id", controller.GetInfografisController)
		infografis.POST("/get-all", controller.GetAllInfografisController)
	}

	router.Run()
}
