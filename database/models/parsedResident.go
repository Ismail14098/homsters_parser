package models

import (
	"gorm.io/gorm"
	"time"
)

type ParsedResident struct {
	ResidentID uint			 `gorm:"primaryKey"`
	Resident Resident
	Parsed bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
