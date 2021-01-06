package database

import (
	"fmt"
	"github.com/Ismail14098/homsters_parser/database/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Initialize() (*gorm.DB, error){
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open( postgres.Open(dbConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}

func Inject(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Set("db",db)
		ctx.Next()
	}
}
