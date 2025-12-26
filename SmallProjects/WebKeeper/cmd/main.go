package main

import (
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/routes"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/services"
)

func main() {
	database.ConectDB()
	go services.StartMonit()
	routes.HandleRequests()
}
