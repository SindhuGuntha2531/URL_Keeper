package routes

import (
	"github.com/gin-gonic/gin"
	"kyc_urls/controllers"
)

func GetUrlRoutes(router *gin.Engine) {
	router.GET("/links", controllers.GetLinks)
}

func GetByNameRoutes(router *gin.Engine) {
	router.GET("/:id", controllers.GetLinkByName)
}

func GetByCatRoutes(router *gin.Engine) {
	router.GET("/cat/:id", controllers.GetLinksByCat)
}

func GetByEnvRoutes(router *gin.Engine) {
	router.GET("/env/:id", controllers.GetLinksByEnv)
}

func CreateRoutes(router *gin.Engine) {
	router.POST("/links", controllers.CreateLink)
}

func DeleteRoutes(router *gin.Engine) {
	router.DELETE("/:id", controllers.DeleteLink)
}
