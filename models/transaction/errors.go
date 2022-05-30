package transaction

import (
	errors "github.com/hoxito/walletgo/tools/custerror"
)

// ErrNoUser No Transaction found for this id
var ErrNoTransaction = errors.NewCustom(400, "Transaction Not Found")

// ErrNoUser No Wallet found for this id
var ErrNoWallet = errors.NewCustom(400, "Wallet Not Found")

//NotEnoughBallance wallet has not enough ballance
var NotEnoughBallance = errors.NewCustom(400, "Not Enough Ballance")
