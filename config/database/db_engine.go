package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"smallcase/utils"
)

var (
	dbConnection *gorm.DB
)

func Initialize(dsn string) {
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open database connection. Err: ", utils.Marshal(err))
		panic("failed to connect database")
	} else {
		dbConnection = conn
	}
}

func GetConnection() *gorm.DB {
	return dbConnection
}

func CloseDatabaseConnection() {
	err := dbConnection.Close()
	log.Fatal("Failed to close the database connection. Err: ", utils.Marshal(err))
}
