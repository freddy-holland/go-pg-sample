package models

import (
	"time"
)

type Client struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	Enabled       bool      `json:"enabled"`
	ContactID     uint      `json:"contact_id"`
	POReference   string    `json:"po_reference"`
	BusinessOpen  time.Time `json:"business_open"`
	BusinessClose time.Time `json:"business_close"`
	Rate          float32   `json:"rate"`
	Contact       Contact   `gorm:"foreignKey:ContactID;references:ID"`
}
