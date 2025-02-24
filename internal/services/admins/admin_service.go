package admins

import (
	"atm/internal/db_driver"
	"context"
)

type AdminService struct {
	driver db_driver.IDBDriver
}

func NewAdminService(driver db_driver.IDBDriver) *AdminService {
	return &AdminService{driver: driver}
}

func (s *AdminService) FindAdmin(ctx context.Context, adminNumber string, adminPassword string) error {
	err := s.driver.FindAdmin(ctx, adminNumber, adminPassword)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) CreateAccount(ctx context.Context, accountNumber string, accountPin uint32) {
	s.driver.CreateAccount(ctx, accountNumber, accountPin)
}

func (s *AdminService) DeleteAccount(ctx context.Context, accountNumber string) {
	s.driver.DeleteAccount(ctx, accountNumber)
}
