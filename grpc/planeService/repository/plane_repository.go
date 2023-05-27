package repository

import (
	"FlightBookingAPI/grpc/planeService/models"
	"errors"

	"github.com/google/uuid"
)

type PlaneRepository interface {
	CreatePlane(model *models.Plane) error
	UpdatePlane(model *models.Plane) (*models.Plane, error)
	ListPlanes() ([]*models.Plane, error)
	FindPlane(id string) (*models.Plane, error)
	DeletePlane(id string) error
}

func (db *dbmanager) CreatePlane(model *models.Plane) error {
	if model == nil {
		return errors.New("invalid parameter (model is nil)")
	}
	if model.Id == "" {
		model.Id = uuid.NewString()
	}

	return db.Create(model).Error
}

func (db *dbmanager) UpdatePlane(req *models.Plane) (*models.Plane, error) {
	updatedPlane := &models.Plane{}
	result := db.
		Where(&models.Plane{Id: req.Id}).
		Updates(models.Plane{Status: req.Status, Type: req.Type, TotalSeat: req.TotalSeat}).
		Find(&updatedPlane)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("RECORD NOT FOUND")
	}
	return updatedPlane, nil
}

func (db *dbmanager) ListPlanes() ([]*models.Plane, error) {
	planes := []*models.Plane{}
	err := db.Find(&planes).Error
	return planes, err
}

func (db *dbmanager) FindPlane(id string) (*models.Plane, error) {
	plane := &models.Plane{}
	result := db.Where(&models.Plane{Id: id}).First(&plane)
	return plane, result.Error
}

func (db *dbmanager) DeletePlane(id string) error {
	result := db.Where(&models.Plane{Id: id}).Delete(&models.Plane{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("RECORD NOT FOUND")
	}
	return nil
}
