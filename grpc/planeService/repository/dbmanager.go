package repository

import (
	"FlightBookingAPI/database"
	"FlightBookingAPI/grpc/planeService/models"
	"log"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PlaneRepository, error) {
	db, err := database.OpenDBConnection("planes")

	if err != nil {
		log.Printf("[Error]: Failed to connect database: %v\n", err)
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Plane{},
	)

	if err != nil {
		log.Printf("[Error]: Migrate database: %v\n", err)
		return nil, err
	}

	return &dbmanager{db}, nil
}
