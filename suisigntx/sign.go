package suisigntx

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/yyle88/erero"
	"golang.org/x/crypto/blake2b"
)

// Sign sui-tx-data. cpp from https://github.com/ltp456/go-sui-sdk/blob/888bddf15fd06afc11900f54ff39d24db7ffefb9/client.go#L460
func Sign(privateKeyHex string, txBytesString string) (string, error) {
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", erero.Wro(err)
	}
	txBytes, err := base64.StdEncoding.DecodeString(txBytesString)
	if err != nil {
		return "", erero.Wro(err)
	}
	return SignTx(privateKey, txBytes)
}

// SignTx sign sui-tx. cpp from https://github.com/ltp456/go-sui-sdk/blob/888bddf15fd06afc11900f54ff39d24db7ffefb9/client.go#L460
func SignTx(privateKey []byte, txBytes []byte) (string, error) {
	wallet, err := suiwallet.NewWallet(privateKey)
	if err != nil {
		return "", erero.Wro(err)
	}

	txData := make([]byte, 0)
	txData = append(txData, []byte{0, 0, 0}...) //这里也需要注意需要补个前缀让链识别
	txData = append(txData, txBytes...)
	txHash := blake2b.Sum256(txData)
	signature, err := wallet.Sign(txHash[:])
	if err != nil {
		return "", erero.Wro(err)
	}

	signatureData := make([]byte, 0)
	signatureData = append(signatureData, []byte{0}...) //补个前缀
	signatureData = append(signatureData, signature...)
	signatureData = append(signatureData, wallet.PublicKey...)

	base64Signature := base64.StdEncoding.EncodeToString(signatureData)
	return base64Signature, nil
}
