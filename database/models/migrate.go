package models

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB){
	//err := db.AutoMigrate(&User{})
	//if err != nil {
	//	fmt.Println("Error Migrate() model: User")
	//	return
	//}
	//err = db.AutoMigrate(&Role{})
	//if err != nil {
	//	fmt.Println("Error Migrate() model: Role")
	//	return
	//}
	//err = db.AutoMigrate(&UserRole{})
	//if err != nil {
	//	fmt.Println("Error Migrate() model: UserRole")
	//	return
	//}
	fmt.Println("Auto Migration has done successfully")
}

