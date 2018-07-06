package main

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	)

func CreateConnection() (*gorm.DB,error){
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	//log.Printf("connection string : host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, DBName, password)

	return gorm.Open(
			"postgres",
			fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
						user, password, host, DBName,
			),
	)
}
