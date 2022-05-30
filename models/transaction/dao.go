package transaction

import (
	"context"
	"database/sql"
	"time"

	"github.com/hoxito/walletgo/tools/db"
)

//create a send transaction
func Send(transaction *Transaction) (*Transaction, error) {
	db := db.Mysql()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	var originBallance float64
	var execErr error
	execErr = tx.QueryRow("SELECT ballance FROM wallet WHERE idwallet = ? FOR UPDATE;",
		transaction.OriginId).Scan(&originBallance)
	if execErr != nil {
		_ = tx.Rollback()
		if execErr == sql.ErrNoRows {
			return nil, ErrNoWallet
		}
		return nil, execErr
	}
	if originBallance < transaction.Amount {

		return nil, NotEnoughBallance
	}
	_, execErr = tx.Exec("UPDATE wallet set ballance=ballance-? where idwallet=?;",
		transaction.Amount, transaction.OriginId)
	if execErr != nil {
		_ = tx.Rollback()
		return nil, execErr
	}
	_, execErr = tx.Exec("UPDATE wallet set ballance=ballance+? where idwallet=?;",
		transaction.Amount, transaction.DestinationId)
	if execErr != nil {
		_ = tx.Rollback()
		return nil, execErr
	}
	_, execErr = tx.Exec("INSERT INTO transaction (idtransaction,origin,destination,amount,status,created) values(?,?,?,?,?,?)",
		transaction.TransactionId, transaction.OriginId, transaction.DestinationId, transaction.Amount, transaction.Status, transaction.Created.Time)
	if execErr != nil {
		_ = tx.Rollback()
		return nil, execErr
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return transaction, err
}

// FindAll
func findAll() ([]*Transaction, error) {

	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		return nil, err
	}

	var trs []*Transaction

	for rows.Next() {
		var tr *Transaction

		err = rows.Scan(tr.TransactionId, tr.Amount, tr.DestinationId, tr.OriginId, tr.Created, tr.Status)
		if err != nil {
			return nil, err
		}

		trs = append(trs, tr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trs, nil
}

func findAllDeleted() ([]*Transaction, error) {

	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM transaction WHERE deleted is not null")
	if err != nil {
		return nil, err
	}

	var trs []*Transaction

	for rows.Next() {
		var tr *Transaction

		err = rows.Scan(tr.TransactionId, tr.Amount, tr.DestinationId, tr.OriginId, tr.Created, tr.Status)
		if err != nil {
			return nil, err
		}

		trs = append(trs, tr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trs, nil
}

// FindByID reads an obra social from db
func findByID(TransactionId string) (*Transaction, error) {

	db := db.Mysql()
	defer db.Close()
	row, err := db.Query("SELECT * FROM transaction WHERE id = ?", TransactionId)

	if err != nil {
		return nil, err
	}

	var tr *Transaction

	err = row.Scan(tr.TransactionId, tr.Amount, tr.DestinationId, tr.OriginId, tr.Created, tr.Status)
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrNoTransaction

		default:
			panic(err)
		}
	}
	return tr, nil
}
