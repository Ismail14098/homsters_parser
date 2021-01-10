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
	err = db.AutoMigrate(&Developer{})
	if err != nil {
		i++
		logger.Println("Error Migrate() model: Developer")
		return
	}
	err = db.AutoMigrate(&Flatplan{})
	if err != nil {
		i++
		logger.Println("Error Migrate() model: Flatplan")
		return
	}
	if i == 0 {
		logger.Println("Auto Migration has done successfully")
	} else {
		log.Fatal(err)
	}
}

