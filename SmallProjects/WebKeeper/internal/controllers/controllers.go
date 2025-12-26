package controllers

import (
	"net/http"
	"net/url"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowSites(c *gin.Context) {
	var sites []models.Site
	database.DB.Find(&sites)
	c.JSON(http.StatusOK, sites)
}

func CreateSite(c *gin.Context) {
	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if site.URL == "" || site.Nome == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nome e url são obrigatórios"})
		return
	}

	parsedURL, err := url.ParseRequestURI(site.URL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url inválida, deve começar com http:// ou https://"})
		return
	}

	if err := database.DB.Create(&site).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "site já cadastrado ou erro ao criar"})
		return
	}
	c.JSON(http.StatusCreated, site)
}

func DeleteSite(c *gin.Context) {
	id := c.Param("id")

	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "site não encontrado"})
		return
	}

	database.DB.Delete(&site)
	c.JSON(http.StatusOK, gin.H{"message": "site removido com sucesso"})
}

func EditSite(c *gin.Context) {
	id := c.Param("id")

	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "site não encontrado"})
		return
	}

	var input models.Site
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.URL != "" {
		parsedURL, err := url.ParseRequestURI(input.URL)
		if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url inválida"})
			return
		}
		site.URL = input.URL
	}
	if input.Nome != "" {
		site.Nome = input.Nome
	}

	database.DB.Save(&site)
	c.JSON(http.StatusOK, site)
}

func SearchSite(c *gin.Context) {
	id := c.Param("id")

	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "site não encontrado"})
		return
	}

	c.JSON(http.StatusOK, site)
}

func GetSiteLogs(c *gin.Context) {
	id := c.Param("id")

	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "site não encontrado"})
		return
	}

	var logs []models.LogSite
	database.DB.Where("site_id = ?", id).Order("hora desc").Limit(100).Find(&logs)
	c.JSON(http.StatusOK, logs)
}
