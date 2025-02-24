package accounts

import (
	"atm/internal/models"
	"context"
	"github.com/shopspring/decimal"
)

type IAccountService interface {
	FindAccount(ctx context.Context, accountNumber string, accountPin uint32) (*models.Account, error)
	GetBalance(ctx context.Context) (decimal.Decimal, error)
	Refill(ctx context.Context, amount decimal.Decimal)
	Withdraw(ctx context.Context, amount decimal.Decimal) error
	Transfer(ctx context.Context, receiverAccountNumber string, amount decimal.Decimal) error
	GetTransactionHistory(ctx context.Context) ([]models.Transaction, error)
}
