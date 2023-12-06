package utils

import (
	"event-manager/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=Cambiar123 sslmode=disable")
	if err != nil {
		panic("Error al conectar con la base de datos: " + err.Error())
	}

	// AutoMigrate creará las tablas automáticamente basándose en tus modelos
	DB.AutoMigrate(&models.Event{})
}
