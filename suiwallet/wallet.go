package suiwallet

import (
	"crypto/ed25519"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

// Wallet of sui-block-chain. cpp from https://github.com/ltp456/go-sui-sdk/blob/888bddf15fd06afc11900f54ff39d24db7ffefb9/crypto/ed25519.go#L9
type Wallet struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	AuthKey    [32]byte
}

func NewWallet(seed []byte) (*Wallet, error) {
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)
	data := make([]byte, 0)
	data = append(data, []byte{0x00}...) //需要补个常量以标志这是 ed25519 的地址
	data = append(data, publicKey...)
	hash := blake2b.Sum256(data)
	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		AuthKey:    hash,
	}, nil
}

func (kp *Wallet) Address() string {
	return fmt.Sprintf("0x%x", kp.AuthKey)
}

func (kp *Wallet) Public() ed25519.PublicKey {
	return kp.PublicKey
}

func (kp *Wallet) Sign(data []byte) ([]byte, error) {
	return ed25519.Sign(kp.PrivateKey, data), nil
}

func (kp *Wallet) Verify(message, signature []byte) bool {
	return ed25519.Verify(kp.PublicKey, message, signature)
}
