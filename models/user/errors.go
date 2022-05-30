package user

import (
	errors "github.com/hoxito/walletgo/tools/custerror"
)

// ErrNoUser No user found for this id
var ErrNoUser = errors.NewCustom(400, "User Not Found")

// ErrNoWallet No user found for this id
var ErrNoWallet = errors.NewCustom(400, "Wallet Not Found")
