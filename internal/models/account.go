package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Account struct {
	id      pgtype.UUID
	number  string
	pin     uint32
	balance decimal.Decimal
}

func NewAccount(id pgtype.UUID, number string, pin uint32, balance decimal.Decimal) *Account {
	return &Account{id: id, number: number, pin: pin, balance: balance}
}

func (account *Account) GetNumber() string {
	return account.number
}

func (account *Account) GetPin() uint32 {
	return account.pin
}

func (account *Account) GetBalance() decimal.Decimal {
	return account.balance
}

func (account *Account) UpdateBalance(newBalance decimal.Decimal) {
	account.balance = newBalance
}
