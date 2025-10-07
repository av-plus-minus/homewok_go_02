package repository

import "ledger/internal/domain"

type TransactionRepository struct {
	transactions []domain.Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		transactions: make([]domain.Transaction, 0),
	}
}

func (r *TransactionRepository) Save(transaction domain.Transaction) error {
	transaction.ID = len(r.transactions) + 1
	r.transactions = append(r.transactions, transaction)
	return nil
}

func (r *TransactionRepository) FindAll() []domain.Transaction {
	result := make([]domain.Transaction, len(r.transactions))
	copy(result, r.transactions)
	return result
}
