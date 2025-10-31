// Package suisigntx: Transaction signing implementation with Ed25519 cryptographic operations
// Provides functions to sign SUI blockchain transaction data using private keys
// Supports both hex-encoded and raw byte format private keys and transaction data
// Generates Base64-encoded signatures compatible with SUI blockchain requirements
//
// suisigntx: 使用 Ed25519 加密操作的交易签名实现
// 提供使用私钥签署 SUI 区块链交易数据的函数
// 支持十六进制编码和原始字节格式的私钥和交易数据
// 生成与 SUI 区块链要求兼容的 Base64 编码签名
package suisigntx

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/yyle88/erero"
	"golang.org/x/crypto/blake2b"
)

// Sign creates a signature from hex-encoded private key and Base64-encoded transaction data
// Accepts privateKeyHex as hex string and txBytesString as Base64 string
// Returns Base64-encoded signature or error if signing fails
// Reference implementation from github.com/ltp456/go-sui-sdk
//
// Sign 从十六进制编码的私钥和 Base64 编码的交易数据创建签名
// 接受 privateKeyHex 作为十六进制字符串，txBytesString 作为 Base64 字符串
// 返回 Base64 编码的签名，如果签名失败则返回错误
// 参考实现来自 github.com/ltp456/go-sui-sdk
func Sign(privateKeyHex string, txBytesString string) (string, error) {
	// Decode hex private key
	// 解码十六进制私钥
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", erero.Wro(err)
	}

	// Decode Base64 transaction bytes
	// 解码 Base64 交易字节
	txBytes, err := base64.StdEncoding.DecodeString(txBytesString)
	if err != nil {
		return "", erero.Wro(err)
	}

	return SignTx(privateKey, txBytes)
}

// SignTx signs transaction bytes using raw private key bytes
// Accepts raw private key bytes and transaction data bytes
// Returns Base64-encoded signature with scheme flag and public key
// Reference implementation from github.com/ltp456/go-sui-sdk
//
// SignTx 使用原始私钥字节签署交易字节
// 接受原始私钥字节和交易数据字节
// 返回带有方案标志和公钥的 Base64 编码签名
// 参考实现来自 github.com/ltp456/go-sui-sdk
func SignTx(privateKey []byte, txBytes []byte) (string, error) {
	// Create wallet from private key
	// 从私钥创建钱包
	wallet, err := suiwallet.NewWallet(privateKey)
	if err != nil {
		return "", erero.Wro(err)
	}

	// Prepare transaction data with intent prefix
	// Add [0, 0, 0] prefix required by blockchain to identify transaction intent
	// 准备带有意图前缀的交易数据
	// 添加区块链所需的 [0, 0, 0] 前缀以识别交易意图
	txData := make([]byte, 0)
	txData = append(txData, []byte{0, 0, 0}...)
	txData = append(txData, txBytes...)

	// Compute Blake2b-256 hash of transaction data
	// 计算交易数据的 Blake2b-256 哈希
	txHash := blake2b.Sum256(txData)

	// Sign the transaction hash
	// 签署交易哈希
	signature, err := wallet.Sign(txHash[:])
	if err != nil {
		return "", erero.Wro(err)
	}

	// Build complete signature data with scheme flag
	// Format: [scheme_flag][signature][public_key]
	// Add [0] prefix to indicate Ed25519 signature scheme
	// 构建带有方案标志的完整签名数据
	// 格式：[方案标志][签名][公钥]
	// 添加 [0] 前缀表示 Ed25519 签名方案
	signatureData := make([]byte, 0)
	signatureData = append(signatureData, []byte{0}...)
	signatureData = append(signatureData, signature...)
	signatureData = append(signatureData, wallet.PublicKey...)

	// Encode signature as Base64 string
	// 将签名编码为 Base64 字符串
	base64Signature := base64.StdEncoding.EncodeToString(signatureData)
	return base64Signature, nil
}
