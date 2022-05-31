package rest

import (
	"testing"

	"github.com/hoxito/walletgo/models/user"
	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {

	response, err := user.Wallet("testuser1", "testwallet1")
	assert.NoError(t, err)
	assert.Equal(t, "testwallet1", response.WalletId)
}
func TestWalletNotFound(t *testing.T) {

	_, err := user.Wallet("testuser1", "unexistingId")
	assert.Error(t, err)
}

func TestWalletEmptyId(t *testing.T) {

	_, err := user.Wallet("testuser1", "")
	assert.Error(t, err)
}
