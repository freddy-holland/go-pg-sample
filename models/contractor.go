package models

import (
	"time"
)

type Contractor struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   bool      `json:"enabled"`
	ContactID uint      `json:"contact_id"`
	Rate      float32   `json:"rate"`
	Contact   Contact   `gorm:"foreignKey:ContactID;references:ID"`
}
