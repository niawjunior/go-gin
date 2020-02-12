package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/niawjunior/gin-app/types"
)

var db *gorm.DB

func Init() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	fmt.Println(dbUri)

	connect, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Println(err)
	}

	db = connect
	db.Debug().AutoMigrate(&types.Users{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
