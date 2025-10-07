package main

import (
	"fmt"
	"time"

	"ledger/internal/domain"
	"ledger/internal/service"
)

func main() {
	fmt.Println("Ledger service started")

	transactionService := service.NewTransactionService()

	testTransactions := []domain.Transaction{
		{
			Amount:      1500.50,
			Category:    "Salary",
			Description: "Monthly salary",
			Date:        time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Amount:      250.75,
			Category:    "Food",
			Description: "Groceries",
			Date:        time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			Amount:      100.00,
			Category:    "Entertainment",
			Description: "Cinema tickets",
			Date:        time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tx := range testTransactions {
		err := transactionService.AddTransaction(tx)
		if err != nil {
			fmt.Printf("Error adding transaction: %v\n", err)
		} else {
			fmt.Printf("Transaction added successfully: %s (%.2f)\n", tx.Category, tx.Amount)
		}
	}

	invalidTx := domain.Transaction{
		Amount:      0,
		Category:    "Test",
		Description: "Invalid transaction",
	}
	if err := transactionService.AddTransaction(invalidTx); err != nil {
		fmt.Printf("Expected error for zero amount: %v\n", err)
	}

	fmt.Println("\n=== All Transactions ===")
	allTransactions := transactionService.ListTransactions()
	for _, tx := range allTransactions {
		fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format("2006-01-02"))
	}
}
