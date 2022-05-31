package rest

import (
	"testing"

	"github.com/hoxito/walletgo/models/user"
	"github.com/stretchr/testify/assert"
)

func TestGetWallet(t *testing.T) {

	response, err := user.GetWallet("testuser1", "testwallet1")
	assert.NoError(t, err)
	assert.Equal(t, "testwallet1", response.WalletId)
}
func TestGetWalletNotFound(t *testing.T) {

	_, err := user.GetWallet("testuser1", "unexistingId")
	assert.Error(t, err)
}

func TestGetWalletEmptyId(t *testing.T) {

	_, err := user.GetWallet("testuser1", "")
	assert.Error(t, err)
}
