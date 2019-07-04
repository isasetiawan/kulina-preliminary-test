package database

import (
	"fmt"
	"hello/basic_coding/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func Init() {
	connection_string := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD"),
	)
	db, err = gorm.Open("postgres", connection_string)

	if err != nil {
		fmt.Println(err)
		panic("DB error")
	}
	db.AutoMigrate(&model.UserReview{})
}

func DbManager() *gorm.DB {
	return db
}
