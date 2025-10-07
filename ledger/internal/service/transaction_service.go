package service

import (
	"fmt"
	"time"

	"ledger/internal/domain"
	"ledger/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repo: repository.NewTransactionRepository(),
	}
}

func (s *TransactionService) AddTransaction(tx domain.Transaction) error {
	if tx.Amount == 0 {
		return fmt.Errorf("transaction amount cannot be zero")
	}

	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	return s.repo.Save(tx)
}

func (s *TransactionService) ListTransactions() []domain.Transaction {
	return s.repo.FindAll()
}
