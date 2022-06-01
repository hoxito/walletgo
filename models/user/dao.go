package user

import (
	"database/sql"
	"fmt"

	"github.com/hoxito/walletgo/models/wallet"
	"github.com/hoxito/walletgo/tools/db"
)

func insert(user *User) (*User, error) {
	if err := user.ValidateSchema(); err != nil {
		return nil, err
	}

	db := db.Mysql()
	defer db.Close()

	_, err := db.Exec("INSERT INTO user (iduser,name,email,password,created) values(?, ?,?,?,?)",
		user.UserId, user.Name, user.Email, user.Password, user.Created.Time)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAll
func findAll() ([]*User, error) {

	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	var usrs []*User

	for rows.Next() {
		var usr *User

		err := rows.Scan(usr.UserId, usr.Name, usr.Email, usr.Password, usr.Created, usr.Deleted)
		if err != nil {
			return nil, err
		}

		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usrs, nil
}

func getWallets(UserId string) ([]*wallet.Wallet, error) {
	db := db.Mysql()
	defer db.Close()
	rows, err := db.Query("SELECT idwallet ,name ,iduser ,ballance ,currency ,created, deleted, updated FROM wallet WHERE deleted is null AND iduser=?", UserId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var wts []*wallet.Wallet

	for rows.Next() {
		var wlt wallet.Wallet

		err := rows.Scan(&wlt.WalletId, &wlt.Name, &wlt.UserId, &wlt.Ballance, &wlt.Currency, &wlt.Created, &wlt.Deleted, &wlt.Updated)
		if err != nil {
			return nil, err
		}

		wts = append(wts, &wlt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return wts, nil
}
func getWallet(UserId string, WalletId string) (*wallet.Wallet, error) {

	db := db.Mysql()
	defer db.Close()
	var wlt wallet.Wallet
	err := db.QueryRow("SELECT * FROM wallet WHERE deleted is null AND iduser=? AND idwallet=?",
		UserId,
		WalletId).Scan(
		&wlt.WalletId,
		&wlt.Name,
		&wlt.UserId,
		&wlt.Ballance,
		&wlt.Currency,
		&wlt.Created,
		&wlt.Deleted,
		&wlt.Updated)

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrNoWallet

		default:
			return nil, err
		}
	}

	return &wlt, nil
}

// FindByID reads a user from db
func findByID(UserId string) (*User, error) {

	db := db.Mysql()
	defer db.Close()

	var usr User
	err := db.QueryRow("SELECT * FROM user WHERE iduser=?",
		UserId).Scan(
		&usr.UserId,
		&usr.Name,
		&usr.Email,
		&usr.Password,
		&usr.Created,
		&usr.Deleted)

	if err != nil {
		return nil, err
	}

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrNoWallet

		default:
			return nil, err
		}
	}
	return &usr, nil
}

// FindByLogin finds a user if username and password match
func findByLogin(user *Login) (*User, error) {

	db := db.Mysql()
	defer db.Close()
	fmt.Println("HERE:", user.Name, user.Password)
	var usr User
	err := db.QueryRow("SELECT * FROM user WHERE name=? and password=?",
		user.Name,
		user.Password).Scan(
		&usr.UserId,
		&usr.Name,
		&usr.Email,
		&usr.Password,
		&usr.Created,
		&usr.Deleted)
	if err != nil {
		return nil, ErrWrongLogin
	}

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return nil, ErrWrongLogin

		default:
			return nil, err
		}
	}
	return &usr, nil
}
