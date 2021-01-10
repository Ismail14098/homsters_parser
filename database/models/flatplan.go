package models

import "gorm.io/gorm"

type Flatplan struct {
	gorm.Model
	ResidentID uint
	Resident Resident
	RoomCount uint
	SqM float64
	MinLevel uint
	MaxLevel uint
	Image string `gorm:"uniqueIndex"`
}
