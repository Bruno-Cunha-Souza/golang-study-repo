package routes

import (
	"os"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/sites", controllers.ShowSites)
	r.GET("/sites/:id", controllers.SearchSite)
	r.GET("/sites/:id/logs", controllers.GetSiteLogs)
	r.POST("/sites", controllers.CreateSite)
	r.DELETE("/sites/:id", controllers.DeleteSite)
	r.PATCH("/sites/:id", controllers.EditSite)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
