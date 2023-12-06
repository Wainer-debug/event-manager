package handlers

import (
	"time"
)

// RespuestaEvento es una estructura específica para respuestas JSON de eventos.
type RespuestaEvento struct {
	ID                 uint      `json:"id"`
	Nombre             string    `json:"nombre"`
	TypeEventID        uint      `json:"type_event_id"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	Status             string    `json:"status"`
	ManagementRequired bool      `json:"management_required"`
	ManagementStatus   string    `json:"management_status"`
}

// RespuestaGestion es una estructura específica para respuestas JSON de eventos clasificados por gestión.
type RespuestaGestion struct {
	ManagementRequired    []RespuestaEvento `json:"management_required"`
	NonManagementRequired []RespuestaEvento `json:"non_management_required"`
}

// Respuesta es una estructura general para respuestas JSON.
type Respuesta struct {
	Mensaje string `json:"mensaje"`
}
