package wallet

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Wallet data structure
type Wallet struct {
	WalletId string       `json:"walletId"  validate:"required"`
	Name     string       `json:"name" validate:"required"`
	UserId   string       `json:"userId" `
	Ballance string       `json:"ballance" validate:"required"`
	Currency string       `json:"currency" validate:"required"`
	Created  sql.NullTime `json:"created" validate:"required"`
	Deleted  sql.NullTime `json:"deleted" validate:"required"`
	Updated  sql.NullTime `json:"updated" validate:"required"`
}

// NewWallet Nueva instancia de usuario
func NewWallet() *Wallet {
	var now sql.NullTime
	now.Time = time.Now()
	now.Valid = true
	return &Wallet{
		WalletId: uuid.NewString(),
		Created:  now,
	}
}

// ValidateSchema validates wallet structure
func (e *Wallet) ValidateSchema() error {
	return validator.New().Struct(e)
}

func (e *WalletRequest) ValidateSchema() error {
	return validator.New().Struct(e)
}
