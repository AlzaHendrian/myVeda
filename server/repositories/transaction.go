package repositories

import (
	"myveda/models"

	"gorm.io/gorm"
)

type RepoTransaction interface {
	FindTransactionByLogin(userId int) ([]models.Transaction, error)
	FindTransactions() ([]models.Transaction, error)
}

func TransactionRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactionByLogin(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id=?", userId).Preload("Product").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Product").Preload("User").Find(&transactions).Error

	return transactions, err
}
