package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbUrl := os.Getenv("database_url")

	conn, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Gender{}, &Place{}, &SexActType{})
}

func GetDB() *gorm.DB {
	return db
}
