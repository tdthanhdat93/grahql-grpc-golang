package repository

import (
	"FlightBookingAPI/database"
	"FlightBookingAPI/grpc/flightService/models"
	"log"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.OpenDBConnection("flights")

	if err != nil {
		log.Printf("[Error]: Failed to connect database: %v\n", err)
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		log.Printf("[Error]: Migrate database: %v\n", err)
		return nil, err
	}

	return &dbmanager{db}, nil
}
