package models

import "time"

type Event struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Name               string    `json:"name" gorm:"column:name"`
	TypeEventID        uint      `json:"type_event_id" gorm:"column:type_event_id"`
	Description        string    `json:"description" gorm:"column:description"`
	Date               time.Time `json:"date" gorm:"column:date;type:timestamp with time zone;default:current_timestamp"`
	Status             string    `json:"status" gorm:"column:status"`
	ManagementRequired bool      `json:"management_required" gorm:"column:management_required"`
	ManagementStatus   string    `json:"management_status" gorm:"column:management_status"`
}
