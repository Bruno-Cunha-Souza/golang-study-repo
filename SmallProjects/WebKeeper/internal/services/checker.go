package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
)

const delay = 5

func StartMonit() {
	for {
		var sites []models.Site
		result := database.DB.Find(&sites)

		if result.Error != nil {
			fmt.Println("Erro ao buscar sites:", result.Error)
			return
		}
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site models.Site) {
	resp, err := http.Get(site.URL)
	if err != nil {
		fmt.Println("Ocorreu um erro ao conectar ao URL:", err)
		return
	}

	defer resp.Body.Close()
	statusCode := resp.StatusCode

	switch statusCode {
	case 200:
		logDes := "Sucess Connecting"
		saveLogsSite(site.ID, statusCode, logDes)
	case 400:
		logDes := "Bad Request"
		saveLogsSite(site.ID, statusCode, logDes)
	case 401:
		logDes := "Unauthorized"
		saveLogsSite(site.ID, statusCode, logDes)
	case 404:
		logDes := "Not Found"
		saveLogsSite(site.ID, statusCode, logDes)
	case 500:
		logDes := "Internal Server Error"
		saveLogsSite(site.ID, statusCode, logDes)
	case 502:
		logDes := "Bad Gateway"
		saveLogsSite(site.ID, statusCode, logDes)
	case 504:
		logDes := "Gateway Timeout"
		saveLogsSite(site.ID, statusCode, logDes)
	default:
		logDes := "Error conecting"
		saveLogsSite(site.ID, statusCode, logDes)
	}
}
