package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getMonitorInterval() time.Duration {
	if interval := os.Getenv("MONITOR_INTERVAL"); interval != "" {
		if seconds, err := strconv.Atoi(interval); err == nil && seconds > 0 {
			return time.Duration(seconds) * time.Second
		}
	}
	return 5 * time.Second
}

func StartMonit() {
	interval := getMonitorInterval()
	fmt.Printf("Iniciando monitoramento (intervalo: %v)\n", interval)

	for {
		var sites []models.Site
		result := database.DB.Find(&sites)

		if result.Error != nil {
			fmt.Println("Erro ao buscar sites:", result.Error)
			time.Sleep(interval)
			continue
		}

		for _, site := range sites {
			go testSite(site)
		}
		time.Sleep(interval)
	}
}

func testSite(site models.Site) {
	resp, err := httpClient.Get(site.URL)
	if err != nil {
		saveLogsSite(site.ID, 0, "Connection Error: "+err.Error())
		return
	}
	defer resp.Body.Close()

	logDes := http.StatusText(resp.StatusCode)
	if logDes == "" {
		logDes = "Unknown Status"
	}
	saveLogsSite(site.ID, resp.StatusCode, logDes)
}
