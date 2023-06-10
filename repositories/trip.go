package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	FindTripId(Id int) (models.Trip, error)
	GetUpdateId(Id int) (models.Trip, error)
	GetCountryId(Id int) (models.CountryResponse, error)
	DeleteTrip(Id int, Trip models.Trip) (models.Trip, error)
	CreateTrip(Trip models.Trip) (models.Trip, error)
	UpdateTrip(Id int, Trip models.Trip) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindTrip() ([]models.Trip, error) {
	var Countries []models.Trip
	err := r.db.Preload("Country").Find(&Countries).Error

	return Countries, err
}

func (r *repositories) FindTripId(Id int) (models.Trip, error) {
	var Countries models.Trip
	err := r.db.Preload("Country").First(&Countries, Id).Error

	return Countries, err
}
func (r *repositories) GetUpdateId(Id int) (models.Trip, error) {
	var Countries models.Trip
	err := r.db.Preload("Country").First(&Countries, Id).Error

	return Countries, err
}
func (r *repositories) GetCountryId(Id int) (models.CountryResponse, error) {
	var Countries models.CountryResponse
	err := r.db.First(&Countries, Id).Error

	return Countries, err
}

func (r *repositories) DeleteTrip(Id int, Trip models.Trip) (models.Trip, error) {
	err := r.db.Delete(&Trip).Error

	return Trip, err
}

func (r *repositories) CreateTrip(Trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&Trip).Error

	return Trip, err
}

func (r *repositories) UpdateTrip(Id int, Trip models.Trip) (models.Trip, error) {
	err := r.db.Save(&Trip).Error

	return Trip, err
}
