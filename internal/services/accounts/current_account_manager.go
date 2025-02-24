package accounts

import (
	"atm/internal/models"
)

type CurrentAccountManager struct {
	account *models.Account
}

func NewCurrentAccountManager() *CurrentAccountManager {
	return &CurrentAccountManager{}
}

func (m *CurrentAccountManager) SetAccount(account *models.Account) {
	m.account = account
}
