package repository

import (
	"FlightBookingAPI/grpc/flightService/models"
	"errors"

	"github.com/google/uuid"
)

type FlightRepository interface {
	CreateFlight(model *models.Flight) error
	UpdateFlight(model *models.Flight) (*models.Flight, error)
	ListFlights() ([]*models.Flight, error)
	FindFlightsByPlaneId(id string) ([]*models.Flight, error)
	FindFlights(param *models.FindFlightParams) ([]*models.Flight, error)
	GetFlight(id string) (*models.Flight, error)
	UpdateFlightAvailableSeat(model *models.Flight) (*models.Flight, error)
	DeleteFlight(id string) error
}

func (db *dbmanager) CreateFlight(model *models.Flight) error {
	if model == nil {
		return errors.New("invalid parameter (model is nil)")
	}
	if model.Id == "" {
		model.Id = uuid.NewString()
	}

	return db.Create(model).Error
}

func (db *dbmanager) UpdateFlight(model *models.Flight) (*models.Flight, error) {
	updatedFlight := &models.Flight{}
	result := db.
		Where(&models.Flight{Id: model.Id}).
		Updates(models.Flight{
			PlaneId:       model.PlaneId,
			Departure:     model.Departure,
			Arrival:       model.Arrival,
			TimeDepart:    model.TimeDepart,
			TimeArrive:    model.TimeArrive,
			AvailableSeat: model.AvailableSeat,
		}).
		Find(&updatedFlight)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}
	return updatedFlight, nil
}

func (db *dbmanager) UpdateFlightAvailableSeat(model *models.Flight) (*models.Flight, error) {
	updatedFlight := &models.Flight{}
	result := db.
		Model(&model).
		Update("available_seat", model.AvailableSeat). // Allow update value 0
		Find(&updatedFlight)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}
	return updatedFlight, nil
}

func (db *dbmanager) ListFlights() ([]*models.Flight, error) {
	flights := []*models.Flight{}
	err := db.Find(&flights).Error
	return flights, err
}

func (db *dbmanager) FindFlightsByPlaneId(id string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	err := db.Where(&models.Flight{PlaneId: id}).Find(&flights).Error
	return flights, err
}

func (db *dbmanager) FindFlights(param *models.FindFlightParams) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	err := db.Where(&models.Flight{Arrival: param.Arrival, Departure: param.Departure}).
		Where("time_depart BETWEEN ? AND ?", param.StartTimeRange, param.EndTimeRange).Find(&flights).Error
	return flights, err
}

func (db *dbmanager) GetFlight(id string) (*models.Flight, error) {
	flight := &models.Flight{}
	err := db.Where(&models.Flight{Id: id}).Find(&flight).Error
	return flight, err
}

func (db *dbmanager) DeleteFlight(id string) error {
	result := db.Where(&models.Flight{Id: id}).Delete(&models.Flight{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("RECORD NOT FOUND")
	}
	return nil
}
