package main

import (
	accounts2 "atm/internal/db_driver"
	"atm/internal/services/accounts"
	"atm/internal/services/auth"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func main() {
	connString := "postgres://dmitrii:1111@localhost:5432/postgres"
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	fmt.Println("Connected to database")
	defer dbpool.Close()

	driver := accounts2.NewAccountDriver(dbpool)
	manager := accounts.NewCurrentAccountManager()
	accountService := accounts.NewAccountService(driver, manager)
	authorizationService := auth.NewAuthorizationService(accountService, manager)

	err = authorizationService.AuthorizeAccount(ctx)
	if err != nil {
		log.Fatalf("Unable to authorize account: %v", err)
	}
	fmt.Println("Authorized account")

	for {
		var selectedInput int
		fmt.Println("Input account pin: ")
		_, _ = fmt.Fscan(os.Stdin, &selectedInput)
		switch selectedInput {
		case 0:
			break
		}
	}
}
