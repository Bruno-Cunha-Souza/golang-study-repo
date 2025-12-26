package services

import (
	"fmt"
	"time"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
)

func saveLogsSite(siteID uint, statusCode int, logDes string) {
	log := models.LogSite{
		SiteID:     siteID,
		StatusCode: statusCode,
		LogDes:     logDes,
		Hora:       time.Now(),
	}

	result := database.DB.Create(&log)
	if result.Error != nil {
		fmt.Println("Erro ao salvar log no banco de dados:", result.Error)
	}
}
