package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	FindCountryId(Id int) (models.Country, error)
	DeleteCountry(Id int, Country models.Country) (models.Country, error)
	CreateCountry(Country models.Country) (models.Country, error)
	UpdateCountry(Id int, Country models.Country) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindCountry() ([]models.Country, error) {
	var Countries []models.Country
	err := r.db.Find(&Countries).Error

	return Countries, err
}

func (r *repositories) FindCountryId(Id int) (models.Country, error) {
	var Countries models.Country
	err := r.db.First(&Countries, Id).Error

	return Countries, err
}

func (r *repositories) DeleteCountry(Id int, Country models.Country) (models.Country, error) {
	err := r.db.Delete(&Country).Error

	return Country, err
}

func (r *repositories) CreateCountry(Country models.Country) (models.Country, error) {
	err := r.db.Create(&Country).Error

	return Country, err
}

func (r *repositories) UpdateCountry(Id int, Country models.Country) (models.Country, error) {
	err := r.db.Save(&Country).Error

	return Country, err
}
