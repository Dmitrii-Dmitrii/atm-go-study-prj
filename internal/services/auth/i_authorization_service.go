package auth

import "context"

type IAuthorizationService interface {
	AuthorizeAccount(ctx context.Context) error
	AuthorizeAdmin(ctx context.Context) error
}
