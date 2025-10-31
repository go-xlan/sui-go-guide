// Package suiwallet: Ed25519-based wallet implementation with SUI blockchain address generation
// Provides wallet creation, signing, and verification operations using Ed25519 cryptography
// Generates blockchain addresses through Blake2b-256 hashing with scheme prefix
// Compatible with SUI blockchain wallet requirements and key management
//
// suiwallet: 基于 Ed25519 的钱包实现，包含 SUI 区块链地址生成
// 提供使用 Ed25519 加密的钱包创建、签名和验证操作
// 通过带有方案前缀的 Blake2b-256 哈希生成区块链地址
// 与 SUI 区块链钱包要求和密钥管理兼容
package suiwallet

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/yyle88/erero"
	"golang.org/x/crypto/blake2b"
)

// Wallet represents SUI blockchain wallet with Ed25519 key pair
// Contains private key, public key, and authentication key (address hash)
// Supports signing transactions and verifying signatures
// Reference implementation from github.com/ltp456/go-sui-sdk
//
// Wallet 代表带有 Ed25519 密钥对的 SUI 区块链钱包
// 包含私钥、公钥和认证密钥（地址哈希）
// 支持签署交易和验证签名
// 参考实现来自 github.com/ltp456/go-sui-sdk
type Wallet struct {
	PrivateKey ed25519.PrivateKey // Ed25519 private key // Ed25519 私钥
	PublicKey  ed25519.PublicKey  // Ed25519 public key // Ed25519 公钥
	AuthKey    [32]byte           // Blake2b-256 hash of [0x00][public_key] // [0x00][公钥] 的 Blake2b-256 哈希
}

// NewWallet creates a new wallet instance from 32-byte seed
// Accepts seed bytes and generates Ed25519 key pair
// Computes authentication key (address) through Blake2b-256 hashing
// Returns wallet instance or error if creation fails
//
// NewWallet 从 32 字节种子创建新的钱包实例
// 接受种子字节并生成 Ed25519 密钥对
// 通过 Blake2b-256 哈希计算认证密钥（地址）
// 返回钱包实例，如果创建失败则返回错误
func NewWallet(seed []byte) (*Wallet, error) {
	// Generate Ed25519 key pair from seed
	// 从种子生成 Ed25519 密钥对
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)

	// Build data with scheme flag prefix
	// Add [0x00] prefix to indicate Ed25519 address scheme
	// 构建带有方案标志前缀的数据
	// 添加 [0x00] 前缀表示 Ed25519 地址方案
	data := make([]byte, 0)
	data = append(data, []byte{0x00}...)
	data = append(data, publicKey...)

	// Compute Blake2b-256 hash as authentication key
	// 计算 Blake2b-256 哈希作为认证密钥
	hash := blake2b.Sum256(data)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		AuthKey:    hash,
	}, nil
}

// NewWalletV2 creates a new wallet from hex-encoded private key string
// Accepts hex string private key and decodes it to bytes
// Returns wallet instance or error if decode or creation fails
//
// NewWalletV2 从十六进制编码的私钥字符串创建新钱包
// 接受十六进制字符串私钥并将其解码为字节
// 返回钱包实例，如果解码或创建失败则返回错误
func NewWalletV2(privateKeyHex string) (*Wallet, error) {
	// Decode hex private key
	// 解码十六进制私钥
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewWallet(privateKey)
}

// Address returns the wallet address as hex string with 0x prefix
// Derives address from AuthKey (Blake2b-256 hash)
// Returns formatted address string
//
// Address 返回带有 0x 前缀的十六进制字符串钱包地址
// 从 AuthKey（Blake2b-256 哈希）派生地址
// 返回格式化的地址字符串
func (kp *Wallet) Address() string {
	return fmt.Sprintf("0x%x", kp.AuthKey)
}

// Public returns the Ed25519 public key
// Returns public key bytes used in signatures
//
// Public 返回 Ed25519 公钥
// 返回签名中使用的公钥字节
func (kp *Wallet) Public() ed25519.PublicKey {
	return kp.PublicKey
}

// Sign creates Ed25519 signature of given data
// Accepts data bytes to sign
// Returns signature bytes or error if signing fails
//
// Sign 创建给定数据的 Ed25519 签名
// 接受要签名的数据字节
// 返回签名字节，如果签名失败则返回错误
func (kp *Wallet) Sign(data []byte) ([]byte, error) {
	return ed25519.Sign(kp.PrivateKey, data), nil
}

// Verify checks if signature is valid with message and public key
// Accepts message bytes and signature bytes
// Returns true if signature is valid, false otherwise
//
// Verify 检查签名与消息和公钥是否有效
// 接受消息字节和签名字节
// 如果签名有效返回 true，否则返回 false
func (kp *Wallet) Verify(message, signature []byte) bool {
	return ed25519.Verify(kp.PublicKey, message, signature)
}
