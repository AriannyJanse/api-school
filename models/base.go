package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	dbURI := ""

	if dbType == "mysql" {
		dbURI = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName) //Build connection string
	} else {
		dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	}

	fmt.Println(dbURI)

	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Student{}, &Teacher{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
