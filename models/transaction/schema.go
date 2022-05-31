package transaction

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TransactionRequest struct {
	OriginId      string  `json:"originId"  validate:"required,necsfield=DestinationId,max=50"`
	DestinationId string  `json:"destinationId"  validate:"required,max=50"`
	Amount        float64 `json:"amount"  validate:"required,gte=1,lte=1000000"`
}
type Transaction struct {
	TransactionId string       `json:"transactionId"  validate:"required,max=50"`
	OriginId      string       `json:"originId"  validate:"required,necsfield=DestinationId,max=50"`
	DestinationId string       `json:"destinationId"  validate:"required,max=50"`
	Status        string       `json:"status" validate:"required,max=50"`
	Amount        float64      `json:"amount"  validate:"required,gte=1,lte=1000000"`
	Created       sql.NullTime `json:"created"`
}

// NewOrder : new order instance
func NewTransaction() *Transaction {
	var now sql.NullTime
	now.Time = time.Now()
	now.Valid = true
	return &Transaction{
		TransactionId: uuid.NewString(),
		Created:       now,
		Status:        "Created",
	}
}

// ValidateSchema Validates structures
func (e *TransactionRequest) ValidateSchema() error {
	return validator.New().Struct(e)
}

func (e *Transaction) ValidateSchema() error {
	return validator.New().Struct(e)
}
