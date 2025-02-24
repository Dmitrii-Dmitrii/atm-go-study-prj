package db_driver

import (
	"atm/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
	"os"
	"time"
)

type DBDriver struct {
	rwdb *pgxpool.Pool
}

func NewAccountDriver(rwdb *pgxpool.Pool) *DBDriver {
	return &DBDriver{rwdb: rwdb}
}

func (d *DBDriver) FindAccount(ctx context.Context, accountNumber string) (*models.Account, error) {
	var accountId pgtype.UUID
	var accountBalance decimal.Decimal
	var accountPin uint32
	err := d.rwdb.QueryRow(ctx, queryFindAccount, accountNumber).Scan(&accountId, &accountPin, &accountBalance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	account := models.NewAccount(accountId, accountNumber, accountPin, accountBalance)
	return account, nil
}

func (d *DBDriver) GetBalance(ctx context.Context, accountNumber string) (decimal.Decimal, error) {
	var accountBalance decimal.Decimal
	err := d.rwdb.QueryRow(ctx, queryGetBalance, accountNumber).Scan(&accountBalance)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return decimal.Zero, fmt.Errorf("Account %d not found: %v\n", accountNumber, err)
		}
		return decimal.Zero, err
	}

	return accountBalance, nil
}

func (d *DBDriver) UpdateBalance(ctx context.Context, accountNumber string, newBalance, amount decimal.Decimal, transactionType models.TransactionType) {
	d.rwdb.QueryRow(ctx, queryUpdateBalance, accountNumber, newBalance)
	transactionId := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	transactionDate := time.Now()
	d.rwdb.QueryRow(ctx, queryAddTransaction, accountNumber, transactionId, transactionType, amount, transactionDate)
}

func (d *DBDriver) Transfer(ctx context.Context, senderAccountNumber, receiverAccountNumber string, newSenderBalance, newReceiverBalance, amount decimal.Decimal) error {
	tx, err := d.rwdb.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	tx.QueryRow(ctx, queryUpdateBalance, senderAccountNumber, newSenderBalance)
	transactionId := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	transactionType := models.OUTGOING_TRANSFER
	transactionDate := time.Now()
	d.rwdb.QueryRow(ctx, queryAddTransaction, senderAccountNumber, transactionId, transactionType, amount, transactionDate)

	tx.QueryRow(ctx, queryUpdateBalance, receiverAccountNumber, newReceiverBalance)
	transactionId = pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	transactionType = models.INCOMING_TRANSFER
	d.rwdb.QueryRow(ctx, queryAddTransaction, receiverAccountNumber, transactionId, transactionType, amount, transactionDate)
	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("ошибка при коммите: %w", err)
	}
	return nil
}

func (d *DBDriver) GetTransactionHistory(ctx context.Context, accountNumber string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	rows, err := d.rwdb.Query(ctx, queryGetTransactions, accountNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.TransactionId, &transaction.TransactionType, &transaction.Amount, &transaction.Date); err != nil {
			return nil, err
		}
		transaction.AccountNumber = accountNumber
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (d *DBDriver) CreateAccount(ctx context.Context, accountNumber string, accountPin uint32) {
	accountId := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	d.rwdb.QueryRow(ctx, queryCreateAccount, accountId, accountNumber, accountPin, decimal.Zero)
}

func (d *DBDriver) DeleteAccount(ctx context.Context, accountNumber string) {
	d.rwdb.QueryRow(ctx, queryDeleteAccount, accountNumber)
}

func (d *DBDriver) FindAdmin(ctx context.Context, adminNumber, adminPassword string) error {
	var adminId pgtype.UUID
	err := d.rwdb.QueryRow(ctx, queryFindAdmin, adminNumber, adminPassword).Scan(&adminId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}
