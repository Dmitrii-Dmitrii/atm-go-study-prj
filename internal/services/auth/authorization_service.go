package auth

import (
	"atm/internal/services/accounts"
	"atm/internal/services/admins"
	"context"
	"fmt"
	"os"
)

type AuthorizationService struct {
	accountService accounts.IAccountService
	manager        *accounts.CurrentAccountManager
	adminService   admins.IAdminService
}

func NewAuthorizationService(service accounts.IAccountService, manager *accounts.CurrentAccountManager) *AuthorizationService {
	return &AuthorizationService{accountService: service, manager: manager}
}

func (s *AuthorizationService) AuthorizeAccount(ctx context.Context) error {
	var accountNumber string
	var accountPin uint32
	fmt.Println("Input account number: ")
	_, err := fmt.Fscan(os.Stdin, &accountNumber)
	if err != nil {
		return err
	}
	fmt.Println("Input account pin: ")
	_, err = fmt.Fscan(os.Stdin, &accountPin)
	if err != nil {
		return err
	}
	account, err := s.accountService.FindAccount(ctx, accountNumber, accountPin)
	if err != nil {
		fmt.Println("Пошел нахуй!", err)
		return err
	}
	s.manager.SetAccount(account)
	return nil
}

func (s *AuthorizationService) AuthorizeAdmin(ctx context.Context) error {
	var adminNumber string
	var adminPassword string
	fmt.Println("Input admin number: ")
	_, err := fmt.Fscan(os.Stdin, &adminNumber)
	if err != nil {
		return err
	}
	fmt.Println("Input admin password: ")
	_, err = fmt.Fscan(os.Stdin, &adminPassword)
	if err != nil {
		return err
	}

	err = s.adminService.FindAdmin(ctx, adminNumber, adminPassword)
	if err != nil {
		fmt.Println("Пошел нахуй!", err)
		return err
	}
	return nil
}
