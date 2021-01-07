package database

import (
	"github.com/Ismail14098/homsters_parser/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func Initialize(log *log.Logger) *gorm.DB{
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open( postgres.Open(dbConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	models.Migrate(db, log)
	return db
}
