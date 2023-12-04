package db

import (
	"fmt"
	"log"
	"os"

	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     =  os.Getenv("USER")
	pass 		 = 	os.Getenv("PASSWORD")
	dbPort   =  os.Getenv("DB_PORT")
	dbname   =  os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, dbPort)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error, please check your database connection ", err)
	}

	fmt.Println("Connection success!")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}