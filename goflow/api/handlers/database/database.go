package internal_database

import (
	"fmt"
	"log"

	config "github.com/gotoflow/core/handlers/config"
	database "github.com/gotoflow/core/handlers/database"
)

func GetDatabase() database.Database {
	config := config.LoadConfig()
	fmt.Println(config.Database.Host)
	db:= database.GetDatabase(config.Database.Driver, config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Database)
	err := db.GetConnection()
	if (err != nil) {
		log.Fatalf("Error on connect DB %s", err)
		panic(err)
	}
	return db
}