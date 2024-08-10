package models

import (
	"time"
)

type Booking struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      time.Time  `json:"end_time"`
	Timezone     string     `json:"timezone"`
	Method       string     `json:"method"`
	Location     *string    `json:"location"`
	Language     string     `json:"language"`
	ClientID     uint       `json:"client_id"`
	ContractorID uint       `json:"contractor_id"`
	Client       Client     `gorm:"foreignKey:ClientID;references:ID"`
	Contractor   Contractor `gorm:"foreignKey:ContractorID;references:ID"`
}
