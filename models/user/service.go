package user

import (
	"fmt"

	"github.com/hoxito/walletgo/models/wallet"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" `
	Password string `json:"password" validate:"required"`
}

// newOS Creates a new user
func CreateUser(user *UserRequest) (*User, error) {

	if err := user.ValidateSchema(); err != nil {
		return nil, err
	}
	newUser := NewUser()
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Password = user.Password

	fmt.Println("User Created")
	newUser, err := insert(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Get wrapper
func Get(UserId string) (*User, error) {
	return findByID(UserId)
}

// Get wrapper
func Wallet(UserId, WalletId string) (*wallet.Wallet, error) {
	return getWallet(UserId, WalletId)
}

// Get wrapper
func Wallets(UserId string) ([]*wallet.Wallet, error) {
	return getWallets(UserId)
}

//  wrapper

func Users() ([]*User, error) {
	return findAll()
}
func FindByLogin(user *Login) (*User, error) {
	return findByLogin(user)
}
