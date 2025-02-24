package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Transaction struct {
	TransactionId   pgtype.UUID
	AccountNumber   string
	TransactionType TransactionType
	Amount          float64
	Date            time.Time
}

type TransactionType int

const (
	REFILL            TransactionType = 0
	WITHDRAW          TransactionType = 1
	INCOMING_TRANSFER TransactionType = 2
	OUTGOING_TRANSFER TransactionType = 3
)
