package main

import (
	"event-manager/handlers"
	"event-manager/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Inicializar la base de datos y otras configuraciones
	utils.InitDatabase()

	// Generar eventos de ejemplo
	handlers.GenerateSampleEvents(utils.DB)

	// Configurar las rutas y el servidor Gin
	router := gin.Default()
	handlers.SetupEventHandlers(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor HTTP
	router.Run(":8080")
}
