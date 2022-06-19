package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//returns database connection
func GetDatabase() *gorm.DB {

	connection, err := gorm.Open(sqlite.Open("pinhead.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the database")
	}

	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalln("Failed to connect to the database")
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful.")
	return connection
}

//create user table in userdb
func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(User{})
}

//closes database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalln("Failed to connect to the database")
	}
	sqldb.Close()
}
