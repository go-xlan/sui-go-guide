package suiwallet_test

import (
	"encoding/hex"
	"testing"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/stretchr/testify/require"
)

func TestNewWallet(t *testing.T) {
	const privateKeyHex = "e4c450f61ba740ae8cc1af0b0fb6a135012747b302154410ad635bde12b411c9"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0xe9507e4d5add8cf5570a4d302550fb9a8ad778a101550c1e377ca4e354c404e6", address)
}

func TestNewWallet_2(t *testing.T) {
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969", address)
}

func TestNewWallet_3(t *testing.T) {
	const privateKeyHex = "00375b3392d6463bb2d1a8e2ae66f1f83a388bc9ab4d4d0b8a378757350b37f7"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0x19795c983a7a50c0be1fb4b3d040faa5150d8d9a39fbdb613397a4e74574d91b", address)
}

func caseNewWallet(t *testing.T, privateKeyHex string) string {
	privateKey, err := hex.DecodeString(privateKeyHex)
	require.NoError(t, err)

	wallet, err := suiwallet.NewWallet(privateKey)
	require.NoError(t, err)

	publicKeyHex := hex.EncodeToString(wallet.Public())
	t.Log("publicKey:", publicKeyHex)

	address := wallet.Address()
	t.Log("address:", address)

	t.Logf("https://suiscan.xyz/testnet/account/%s", address)
	return address
}