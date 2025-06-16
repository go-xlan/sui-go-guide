package suiwallet_test

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
)

func TestNewWallet(t *testing.T) {
	output := make([]byte, 32) // 32 bytes = 64 hex characters
	rese.C1(rand.Read(output))
	prkHex := hex.EncodeToString(output)
	t.Log("random-private-key-hex:", prkHex)
	require.Len(t, prkHex, 64)

	address := caseNewWallet(t, prkHex)
	t.Log("address:", address)
}

func TestNewWallet_1(t *testing.T) {
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
	wallet, err := suiwallet.NewWalletV2(privateKeyHex)
	require.NoError(t, err)

	publicKeyHex := hex.EncodeToString(wallet.Public())
	t.Log("publicKeyHex:", publicKeyHex)

	publicKeyBase64 := base64.StdEncoding.EncodeToString(wallet.Public())
	t.Log("publicKeyBase64:", publicKeyBase64)

	address := wallet.Address()
	t.Log("address:", address)

	t.Logf("https://suiscan.xyz/testnet/account/%s", address)
	return address
}

func TestWallet_Sign(t *testing.T) {
	const privateKeyHex = "00375b3392d6463bb2d1a8e2ae66f1f83a388bc9ab4d4d0b8a378757350b37f7"
	wallet := rese.P1(suiwallet.NewWalletV2(privateKeyHex))

	message := "example"
	signatureBytes, err := wallet.Sign([]byte(message))
	require.NoError(t, err)
	ok := wallet.Verify([]byte(message), signatureBytes)
	require.True(t, ok)
}
