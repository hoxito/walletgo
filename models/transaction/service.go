package transaction

import (
	"fmt"
)

// newOS Creates a new transaction
func CreateTransaction(transaction *TransactionRequest) (*Transaction, error) {

	if err := transaction.ValidateSchema(); err != nil {
		return nil, err
	}
	newTransaction := NewTransaction()
	newTransaction.DestinationId = transaction.DestinationId
	newTransaction.OriginId = transaction.OriginId
	newTransaction.Amount = transaction.Amount

	fmt.Println("Transaction Created")
	newTransaction, err := Send(newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (e *Transaction) SetState(state string) {
	//TODO
}

// Get wrapper
func Get(TransactionId string) (*Transaction, error) {
	return findByID(TransactionId)
}

//  wrapper

func Transactions() ([]*Transaction, error) {
	return findAll()
}
