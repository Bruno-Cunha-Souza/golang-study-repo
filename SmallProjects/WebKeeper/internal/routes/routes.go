package routes

import (
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/sites", controllers.ShowSites)
	r.GET("/SearchSite", controllers.SearchSite)
	r.POST("/CreateSite", controllers.CreateSite)
	r.DELETE("/DeleteSite", controllers.DeleteSite)
	r.PATCH("/EditSite", controllers.EditSite)
	r.Run(":3000")
}
