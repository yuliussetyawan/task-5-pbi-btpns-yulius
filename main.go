package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/db"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil{
	 log.Fatalf("Error loading .env file: %s", err)
	}
	db.StartDB()
	r := router.StartRoute()
	r.Run(":8080")
}