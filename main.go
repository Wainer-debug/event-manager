package main

import (
	_ "event-manager/docs"
	"event-manager/handlers"
	"event-manager/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Event Manager API
// @version 1.0
// @description Esta API permite gestionar eventos y organizarlos según el estado de revisión.
// @description En la API se puedan agregar, editar, leer y borrar la información de los eventos y obtener eventos revisados por Requiere Gestión / Sin Gestión
// @host localhost:8080
func main() {

	// Inicializar la base de datos y otras configuraciones
	utils.InitDatabase()

	// Generar eventos de ejemplo.
	handlers.GenerateSampleEvents(utils.DB)

	// Configurar las rutas y el servidor Gin
	router := gin.Default()
	handlers.SetupEventHandlers(router)

	// Ruta visual de Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor HTTP
	router.Run(":8080")
}
