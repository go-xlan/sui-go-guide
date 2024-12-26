package suiwallet_test

import (
	"encoding/hex"
	"testing"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/stretchr/testify/require"
)

func TestNewWallet(t *testing.T) {
	privateKeyHex := "e4c450f61ba740ae8cc1af0b0fb6a135012747b302154410ad635bde12b411c9"
	caseNewWallet(t, privateKeyHex)
}

func TestNewWallet_2(t *testing.T) {
	privateKeyHex := "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"
	caseNewWallet(t, privateKeyHex)
}

func TestNewWallet_3(t *testing.T) {
	privateKeyHex := "00375b3392d6463bb2d1a8e2ae66f1f83a388bc9ab4d4d0b8a378757350b37f7"
	caseNewWallet(t, privateKeyHex)
}

func caseNewWallet(t *testing.T, privateKeyHex string) {
	privateKey, err := hex.DecodeString(privateKeyHex)
	require.NoError(t, err)

	wallet, err := suiwallet.NewWallet(privateKey)
	require.NoError(t, err)

	publicKeyHex := hex.EncodeToString(wallet.Public())
	t.Log("publicKey:", publicKeyHex)

	address := wallet.Address()
	t.Log("address:", address)

	t.Logf("https://suiscan.xyz/testnet/account/%s", address)
}
