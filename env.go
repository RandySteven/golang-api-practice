package main

import (
	"log"
	"os"
	"test-api/app"
	"test-api/configs"
	"test-api/entities/models"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func AppPort() string {
	return os.Getenv("APP_PORT")
}

func InitHandlers() *app.Handlers {
	config := InitConfig()

	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = repository.Automigrate()
	if err != nil {
		return nil
	}

	handlers, err := app.NewHandlers(*repository)
	if err != nil {
		return nil
	}

	return handlers
}
