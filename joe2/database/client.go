package database

import (
	"log"
	"test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func DbConnectAndMigrate() (*gorm.DB, error) {
	Instance, dbError = gorm.Open(mysql.Open("root:Sethjoe12345@tcp(localhost:3306)/test_demo?parseTime=true"), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")

	if dbError = Instance.AutoMigrate(&models.User{}); dbError != nil {
		log.Println(dbError)
	}

	return Instance, dbError
}

// func Migrate() {
// 	Instance.AutoMigrate(&models.User{})
// 	log.Println("Database Migration Completed!")
// }
