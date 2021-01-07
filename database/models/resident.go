package models

import "gorm.io/gorm"

type Resident struct {
	gorm.Model
	Name string
	DeveloperName string
	CityName string
	DistrictName string
	SubDistrictName string
	PricePerSqM float32
	Currency string
	MinSize float32
	MaxSize float32
	MinRoomCount float32
	MaxRoomCount float32
	FloorCount float32
	CommissioningYear int
	CommissioningQuarter int
	ConstructionStatusLocalizedDescription string
	PhoneNumber string
	DirectPhone string
	ComplexUrl string
}
