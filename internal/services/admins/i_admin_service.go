package admins

import "context"

type IAdminService interface {
	FindAdmin(ctx context.Context, adminNumber string, adminPassword string) error
	CreateAccount(ctx context.Context, accountNumber string, accountPin uint32)
	DeleteAccount(ctx context.Context, accountNumber string)
}
