package wallet

import (
	"database/sql"

	"github.com/hoxito/walletgo/models/transaction"
	"github.com/hoxito/walletgo/tools/db"
)

func insert(wallet *Wallet) (*Wallet, error) {
	if err := wallet.ValidateSchema(); err != nil {
		return nil, err
	}

	db := db.Mysql()
	defer db.Close()

	_, err := db.Exec("INSERT INTO wallet (idwallet,name,iduser,ballance,currency,created,deleted,updated) values(?, ?,?,?,?,?,?,?)",
		wallet.WalletId, wallet.Name, wallet.UserId, wallet.Ballance, wallet.Currency, wallet.Created.Time, wallet.Deleted, wallet.Updated)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// FindAll
func findAll() ([]*Wallet, error) {

	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM wallet")
	if err != nil {
		return nil, err
	}

	var wlts []*Wallet

	for rows.Next() {
		var wlt *Wallet

		err := rows.Scan(wlt.WalletId, wlt.Name, wlt.UserId, wlt.Ballance, wlt.Currency, wlt.Created, wlt.Deleted, wlt.Updated)
		if err != nil {
			return nil, err
		}

		wlts = append(wlts, wlt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return wlts, nil
}

// FindByID reads a wallet from db
func findByID(WalletId string) (*Wallet, error) {

	db := db.Mysql()
	defer db.Close()

	var wlt *Wallet
	err := db.QueryRow("SELECT * FROM wallet").Scan(wlt.WalletId, wlt.Name, wlt.UserId, wlt.Ballance, wlt.Currency, wlt.Created, wlt.Deleted, wlt.Updated)

	if err != nil {
		return nil, err
	}

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrNoWallet

		default:
			panic(err)
		}
	}
	return wlt, nil
}

func GetTransactions(WalletId string) ([]*transaction.Transaction, error) {
	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT idtransaction,origin,destination,amount,COALESCE(status,''),created FROM transaction WHERE origin=? OR destination = ?", WalletId, WalletId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var trs []*transaction.Transaction

	for rows.Next() {
		var tr transaction.Transaction

		err := rows.Scan(&tr.TransactionId, &tr.OriginId, &tr.DestinationId, &tr.Amount, &tr.Status, &tr.Created)
		if err != nil {
			return nil, err
		}

		trs = append(trs, &tr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trs, nil
}
func GetTransaction(TransactionId string) (*transaction.Transaction, error) {

	db := db.Mysql()
	defer db.Close()
	var tr transaction.Transaction
	err := db.QueryRow(`SELECT * FROM transaction WHERE idtransaction=?`, TransactionId).Scan(&tr.TransactionId, &tr.OriginId, &tr.DestinationId, &tr.Amount, &tr.Status, &tr.Created)

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrNoWallet

		default:
			return nil, err
		}
	}

	return &tr, nil
}
