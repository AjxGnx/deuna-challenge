package repository

import (
	"github.com/AjxGnx/deuna-challenge/internal/domain/entity"
	"github.com/AjxGnx/deuna-challenge/internal/domain/model"
	"gorm.io/gorm"
)

type Transaction interface {
	Create(transaction model.Transaction) (entity.Transaction, error)
	GetByID(id uint) (model.Transaction, error)
	Update(transaction model.Transaction) (entity.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) Transaction {
	return transactionRepo{
		db,
	}
}

func (repo transactionRepo) Create(transaction model.Transaction) (entity.Transaction, error) {
	result := repo.db.Create(&transaction).Scan(&transaction)
	if result.Error != nil {
		return entity.Transaction{}, result.Error
	}

	return transaction.ToEntity(), nil
}

func (repo transactionRepo) GetByID(id uint) (model.Transaction, error) {
	var transaction model.Transaction

	result := repo.db.First(&transaction, id)
	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}

func (repo transactionRepo) Update(transaction model.Transaction) (entity.Transaction, error) {
	result := repo.db.
		Model(&transaction).
		Where("id = ?", transaction.ID).
		Updates(transaction).
		Scan(&transaction)

	if result.Error != nil {
		return entity.Transaction{}, result.Error
	}

	return transaction.ToEntity(), nil
}
