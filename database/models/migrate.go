package models

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB, logger *log.Logger){
	i := 0
	err := db.AutoMigrate(&Resident{})
	if err != nil {
		i++
		logger.Println("Error Migrate() model: Resident")
		return
	}
	err = db.AutoMigrate(&ParsedResident{})
	if err != nil {
		i++
		logger.Println("Error Migrate() model: ParsedResident")
		return
	}
	if i == 0 {
		logger.Println("Auto Migration has done successfully")
	}
}

