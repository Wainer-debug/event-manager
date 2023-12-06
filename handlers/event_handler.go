package handlers

import (
	"event-manager/models"
	"event-manager/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Event = models.Event

var validate = validator.New()

// Configuración de los manejadores de eventos en el enrutador de Gin.
func SetupEventHandlers(router *gin.Engine) {
	router.GET("/events", GetEvents)
	router.POST("/events", CreateEvent)
	router.PUT("/events/:id", UpdateEvent)
	router.DELETE("/events/:id", DeleteEvent)
	router.GET("/events/management_required", GetManagementStatus)
}

// @Summary Obtener eventos
// @Description Obtiene todos los eventos disponibles
// @ID get-events
// @Produce json
// @Success 200 {array} RespuestaEvento
// @Router /events [get]
func GetEvents(c *gin.Context) {
	var events []models.Event
	if err := utils.DB.Find(&events).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener eventos"})
		return
	}

	c.JSON(200, events)
}

// @Summary Crear un nuevo evento
// @Description Crea un nuevo evento con la información proporcionada en el cuerpo JSON de la solicitud
// @ID create-event
// @Accept json
// @Produce json
// @Param event body models.Event true "Detalles del evento a crear"
// @Success 201 {object} RespuestaEvento
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	var newEvent models.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, Respuesta{Mensaje: err.Error()})
		return
	}

	if err := validate.Struct(newEvent); err != nil {
		c.JSON(http.StatusBadRequest, Respuesta{Mensaje: err.Error()})
		return
	}

	// Establecer el ManagementRequired y ManagementStatus basándome en el tipo de evento y el estado
	if newEvent.Status == "REVISADO" {
		if newEvent.TypeEventID == 1 || newEvent.TypeEventID == 2 {
			newEvent.ManagementRequired = true
			newEvent.ManagementStatus = "Requiere gestión"
		} else {
			newEvent.ManagementRequired = false
			newEvent.ManagementStatus = "Sin gestión"
		}
	} else {
		newEvent.ManagementRequired = false
		newEvent.ManagementStatus = ""
	}

	if err := utils.DB.Create(&newEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Respuesta{Mensaje: "Error al crear el evento"})
		return
	}

	c.JSON(http.StatusCreated, RespuestaEvento{
		ID:                 newEvent.ID,
		Nombre:             newEvent.Name,
		TypeEventID:        newEvent.TypeEventID,
		Description:        newEvent.Description,
		Date:               newEvent.Date,
		Status:             newEvent.Status,
		ManagementRequired: newEvent.ManagementRequired,
		ManagementStatus:   newEvent.ManagementStatus,
	})
}

// @Summary Actualizar un evento existente
// @Description Actualiza un evento existente según el ID proporcionado en la URL con la información del cuerpo JSON de la solicitud
// @ID update-event
// @Accept json
// @Produce json
// @Param id path int true "ID del evento a actualizar"
// @Param event body RespuestaEvento true "Detalles actualizados del evento"
// @Success 200 {object} RespuestaEvento
// @Router /events/{id} [put]
func UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")
	var updatedEvent Event

	if err := c.ShouldBindWith(&updatedEvent, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingEvent Event
	if err := utils.DB.First(&existingEvent, eventID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el evento"})
		}
		return
	}

	// Actualizar campos según sea necesario
	existingEvent.Name = updatedEvent.Name
	existingEvent.Description = updatedEvent.Description

	if err := utils.DB.Save(&existingEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el evento"})
		return
	}

	c.JSON(http.StatusOK, existingEvent)
}

// @Summary Eliminar un evento existente
// @Description Elimina un evento según el ID proporcionado en la URL
// @ID delete-event
// @Produce json
// @Param id path int true "ID del evento a eliminar"
// @Success 200 {object} Respuesta
// @Router /events/{id} [delete]
func DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")

	var existingEvent Event
	if err := utils.DB.First(&existingEvent, eventID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el evento"})
		}
		return
	}

	if err := utils.DB.Delete(&existingEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento eliminado exitosamente"})
}

// Crear eventos nuevos de manera estática.
func GenerateSampleEvents(db *gorm.DB) {

	if err := db.Create(&models.Event{
		ID:                 1,
		Name:               "Ejemplo1",
		TypeEventID:        3,
		Description:        "Descripción del evento Ejemplo1",
		Date:               time.Now(),
		Status:             "PENDIENTE POR REVISAR",
		ManagementRequired: false,
		ManagementStatus:   "",
	}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Error al crear evento 1: el ID ya existe")
		} else {
			fmt.Println("Error al crear evento 1:", err)
		}
	}

	if err := db.Create(&models.Event{
		ID:                 2,
		Name:               "Ejemplo2",
		TypeEventID:        2,
		Description:        "Descripción del evento Ejemplo2",
		Date:               time.Now(),
		Status:             "REVISADO",
		ManagementRequired: true,
		ManagementStatus:   "Sin gestión",
	}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Error al crear evento 2: el ID ya existe")
		} else {
			fmt.Println("Error al crear evento 2:", err)
		}
	}

	if err := db.Create(&models.Event{
		ID:                 3,
		Name:               "Ejemplo3",
		TypeEventID:        1,
		Description:        "Descripción del evento Ejemplo2",
		Date:               time.Now(),
		Status:             "REVISADO",
		ManagementRequired: true,
		ManagementStatus:   "Requiere gestión",
	}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Error al crear evento 3: el ID ya existe")
		} else {
			fmt.Println("Error al crear evento 3:", err)
		}
	}
}

// EventToRespuestaEvento convierte un models.Event a RespuestaEvento
func EventToRespuestaEvento(event models.Event) RespuestaEvento {
	return RespuestaEvento{
		ID:                 event.ID,
		Nombre:             event.Name,
		TypeEventID:        event.TypeEventID,
		Description:        event.Description,
		Date:               event.Date,
		Status:             event.Status,
		ManagementRequired: event.ManagementRequired,
		ManagementStatus:   event.ManagementStatus,
	}
}

// @Summary Obtener eventos revisados por Requiere Gestión / Sin Gestión
// @Description Obtiene una lista de eventos revisados clasificados por Requiere Gestión o Sin Gestión.
// @ID get-management-status
// @Produce json
// @Success 200 {object} RespuestaGestion
// @Router /events/management_required [get]
func GetManagementStatus(c *gin.Context) {
	var managementRequiredEvents []models.Event
	var nonManagementRequiredEvents []models.Event

	if err := utils.DB.Where("management_status = ?", "Requiere gestión").Find(&managementRequiredEvents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Respuesta{Mensaje: "Error al obtener eventos que requieren gestión"})
		return
	}

	if err := utils.DB.Where("management_status = ?", "Sin gestión").Find(&nonManagementRequiredEvents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Respuesta{Mensaje: "Error al obtener eventos sin gestión"})
		return
	}

	// Convertir models.Event a RespuestaEvento
	var managementRequiredResponse []RespuestaEvento
	for _, event := range managementRequiredEvents {
		managementRequiredResponse = append(managementRequiredResponse, EventToRespuestaEvento(event))
	}

	// Convertir models.Event a RespuestaEvento
	var nonManagementRequiredResponse []RespuestaEvento
	for _, event := range nonManagementRequiredEvents {
		nonManagementRequiredResponse = append(nonManagementRequiredResponse, EventToRespuestaEvento(event))
	}

	response := RespuestaGestion{
		ManagementRequired:    managementRequiredResponse,
		NonManagementRequired: nonManagementRequiredResponse,
	}

	c.JSON(http.StatusOK, response)
}
