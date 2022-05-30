package wallet

import (
	"fmt"
)

type WalletRequest struct {
	Name     string `json:"name" validate:"required"`
	UserId   string `json:"email" `
	Currency string `json:"currency" validate:"required"`
}

// newOS Creates a new wallet
func CreateWallet(wallet *WalletRequest) (*Wallet, error) {

	if err := wallet.ValidateSchema(); err != nil {
		return nil, err
	}
	newWallet := NewWallet()
	newWallet.Name = wallet.Name
	newWallet.UserId = wallet.UserId
	newWallet.Currency = wallet.Currency

	fmt.Println("Wallet Created")
	newWallet, err := insert(newWallet)
	if err != nil {
		return nil, err
	}

	return newWallet, nil
}

func (e *Wallet) SetState(state string) {
	//TODO
}

// Get wrapper
func Get(WalletId string) (*Wallet, error) {
	return findByID(WalletId)
}

//  wrapper

func Wallets() ([]*Wallet, error) {
	return findAll()
}
