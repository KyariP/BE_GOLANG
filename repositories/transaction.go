package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	FindTransactionId(Id int) (models.Transaction, error)
	GetTripId(Id int) (models.TripResponse, error)
	GetUserId(Id int) (models.UserResponse, error)
	GetTransByUser(Id int) ([]models.Transaction, error)
	DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	// UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(Status string, Id int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindTransaction() ([]models.Transaction, error) {
	var Tansactions []models.Transaction
	err := r.db.Preload("User").Preload("Trip.Country").Find(&Tansactions).Error

	return Tansactions, err
}
func (r *repositories) GetTransByUser(Id int) ([]models.Transaction, error) {

	var Transactions []models.Transaction
	err := r.db.Where("id_user = ?", Id).Preload("Trip.Country").Find(&Transactions).Error

	return Transactions, err
}

func (r *repositories) FindTransactionId(Id int) (models.Transaction, error) {
	var Tansactions models.Transaction
	err := r.db.Preload("Trip.Country").First(&Tansactions, Id).Error

	return Tansactions, err
}
func (r *repositories) GetTripId(Id int) (models.TripResponse, error) {
	var Trip models.TripResponse
	err := r.db.Preload("Country").First(&Trip, Id).Error

	return Trip, err
}

func (r *repositories) GetUserId(Id int) (models.UserResponse, error) {
	var Users models.UserResponse
	err := r.db.First(&Users, Id).Error

	return Users, err
}

func (r *repositories) DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&Transaction).Error

	return Transaction, err
}

func (r *repositories) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&Transaction).Error

	return Transaction, err
}

// func (r *repositories) UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
// 	err := r.db.Save(&Transaction).Error

// 	return Transaction, err
// }

func (r *repositories) UpdateTransaction(status string, Id int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Trip").Preload("User").First(&transaction, Id)

	if status != transaction.Status && status == "success" {
		var Trip models.Trip
		r.db.First(&Trip, transaction.Trip.Id)
		Trip.Current_Quota = Trip.Current_Quota + transaction.CounterQty
		r.db.Save(&Trip)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
