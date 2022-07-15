package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//returns database connection
func GetDatabase() (*gorm.DB, error) {

	connection, err := gorm.Open(sqlite.Open("pinhead.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqldb, err := connection.DB()
	if err != nil {
		return nil, err
	}

	err = sqldb.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connection successful.")
	return connection, nil
}

//create user table in userdb
func InitialMigration() error {
	allModels := []interface{}{&Score{}, &User{}, &Game{}}
	connection, err := GetDatabase()
	if err != nil {
		return err
	}
	defer CloseDatabase(connection)
	fmt.Println("Automigrating database.")
	err = connection.AutoMigrate(allModels...)
	if err != nil {
		return err
	}
	connection.Migrator().DropTable("GameScore", "GameUser")
	connection.Create(&Game{})
	return nil
}

//closes database connection
func CloseDatabase(connection *gorm.DB) error {
	sqldb, err := connection.DB()
	if err != nil {
		return err
	}
	fmt.Println("Closing database connection.")
	sqldb.Close()
	return nil
}
