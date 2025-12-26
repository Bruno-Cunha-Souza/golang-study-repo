package database

import (
	"log"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDB() {
	stringConect := "host=localhost user=postgres password=my_password dbname=polls port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConect))
	if err != nil {
		log.Panic("Erro ao conectar ao Banco", err.Error())
	}
	DB.AutoMigrate(models.Site{}, models.LogSite{})
}
