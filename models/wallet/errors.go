package wallet

import (
	errors "github.com/hoxito/walletgo/tools/custerror"
)

// ErrNoWallet No wallet found for this id
var ErrNoWallet = errors.NewCustom(400, "Wallet Not Found")
