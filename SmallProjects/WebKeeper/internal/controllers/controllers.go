package controllers

import (
	"net/http"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowSites(c *gin.Context) {
	var site []models.Site
	database.DB.Find(&site)
	c.JSON(200, site)
}
func CreateSite(c *gin.Context) {
	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error()})
		return
	}
	database.DB.Create(&site)
	c.JSON(http.StatusOK, site)
}
func DeleteSite(c *gin.Context) {
	var site models.Site
	id := c.Params.ByName("id")

	database.DB.Delete(&site, id)
	c.JSON(http.StatusOK, gin.H{"data": "Website successfully deleted"})
}
func EditSite(c *gin.Context) {
	var site models.Site
	id := c.Params.ByName("id")
	database.DB.First(&site, id)

	// Atualizar os campos do site com os dados da solicitação JSON
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&site).UpdateColumns(site)
	c.JSON(http.StatusOK, gin.H{"data": "Website successfully Updated"})
}
func SearchSite(c *gin.Context) {
	var site models.Site
	id := c.Params.ByName("id")
	database.DB.First(&site, id)

	if site.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Site not found"})
		return
	}

	c.JSON(http.StatusOK, site)
}
