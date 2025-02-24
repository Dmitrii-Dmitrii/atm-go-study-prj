package accounts

import (
	"atm/internal/db_driver"
	"atm/internal/models"
	"context"
	"errors"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	manager *CurrentAccountManager
	driver  db_driver.IDBDriver
}

func NewAccountService(driver db_driver.IDBDriver, manager *CurrentAccountManager) *AccountService {
	return &AccountService{driver: driver, manager: manager}
}

func (s *AccountService) FindAccount(ctx context.Context, accountNumber string, accountPin uint32) (*models.Account, error) {
	account, err := s.driver.FindAccount(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	if accountPin != account.GetPin() {
		return nil, errors.New("invalid pin")
	}
	return account, nil
}

func (s *AccountService) GetBalance(ctx context.Context) (decimal.Decimal, error) {
	accountBalance, err := s.driver.GetBalance(ctx, s.manager.account.GetNumber())
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return accountBalance, nil
}

func (s *AccountService) Refill(ctx context.Context, amount decimal.Decimal) {
	accountBalance := s.manager.account.GetBalance()
	newBalance := accountBalance.Add(amount)
	s.driver.UpdateBalance(ctx, s.manager.account.GetNumber(), newBalance, amount, models.REFILL)
	s.manager.account.UpdateBalance(newBalance)
}

func (s *AccountService) Withdraw(ctx context.Context, amount decimal.Decimal) error {
	accountBalance := s.manager.account.GetBalance()
	if accountBalance.LessThan(amount) {
		return errors.New("Account balance is lower than the amount to withdraw!")
	}
	newBalance := accountBalance.Sub(amount)
	s.driver.UpdateBalance(ctx, s.manager.account.GetNumber(), newBalance, amount, models.WITHDRAW)
	s.manager.account.UpdateBalance(newBalance)
	return nil
}

func (s *AccountService) Transfer(ctx context.Context, receiverAccountNumber string, amount decimal.Decimal) error {
	senderAccountBalance := s.manager.account.GetBalance()
	if senderAccountBalance.LessThan(amount) {
		return errors.New("Sender account balance is lower than the amount to transfer!")
	}
	newSenderBalance := senderAccountBalance.Sub(amount)
	receiverBalance, err := s.driver.GetBalance(ctx, receiverAccountNumber)
	if err != nil {
		return err
	}
	newReceiverBalance := receiverBalance.Add(amount)
	err = s.driver.Transfer(ctx, s.manager.account.GetNumber(), receiverAccountNumber, newSenderBalance, newReceiverBalance, amount)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) GetTransactionHistory(ctx context.Context) ([]models.Transaction, error) {
	accountNumber := s.manager.account.GetNumber()
	transactionHistory, err := s.driver.GetTransactionHistory(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	return transactionHistory, nil
}
