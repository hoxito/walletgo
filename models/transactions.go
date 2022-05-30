package models

import "database/sql"

type MsgTransaction struct {
	Type    string      `json:"type"`
	Message Transaction `json:"message"`
}
type Transaction struct {
	TransactionId string       `json:"transactionId"`
	OriginId      string       `json:"originId"`
	DestinationId string       `json:"destinationId"`
	Status        string       `json:"status"`
	Amount        float64      `json:"amount"`
	Created       sql.NullTime `json:"created"`
}

type Wallet struct {
	WalletId string       `json:"walletId"`
	UserId   string       `json:"userId"`
	Name     string       `json:"name"`
	Ballance float64      `json:"ballance"`
	Currency string       `json:"currency"`
	Created  sql.NullTime `json:"created"`
	Updated  sql.NullTime `json:"updated"`
	Deleted  sql.NullTime `json:"deleted"`
}

type User struct {
	UserId   string       `json:"userId"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Deleted  sql.NullTime `json:"deleted"`
	Created  sql.NullTime `json:"created"`
}
