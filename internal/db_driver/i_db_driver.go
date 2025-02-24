package db_driver

import (
	"atm/internal/models"
	"context"
	"github.com/shopspring/decimal"
)

type IDBDriver interface {
	FindAccount(ctx context.Context, accountNumber string) (*models.Account, error)
	GetBalance(ctx context.Context, accountNumber string) (decimal.Decimal, error)
	UpdateBalance(ctx context.Context, accountNumber string, newBalance, amount decimal.Decimal, transactionType models.TransactionType)
	Transfer(ctx context.Context, senderAccountNumber, receiverAccountNumber string, newSenderBalance, newReceiverBalance, amount decimal.Decimal) error
	GetTransactionHistory(ctx context.Context, accountNumber string) ([]models.Transaction, error)
	FindAdmin(ctx context.Context, adminNumber, adminPassword string) error
	CreateAccount(ctx context.Context, accountNumber string, accountPin uint32)
	DeleteAccount(ctx context.Context, accountNumber string)
}
