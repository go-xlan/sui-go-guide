// Package suisecret: Pure Go implementation to decode and encode SUI keystore keys
// Provides direct replacement of third-party sui keytool convert command
// Supports ed25519 signature scheme with Base64 and hex format conversion
// Enables seamless wallet creation from decoded private keys
//
// suisecret: 纯 Go 实现来解码和编码 SUI keystore 密钥
// 提供第三方 sui keytool convert 命令的直接替代方案
// 支持 ed25519 签名方案，包含 Base64 和十六进制格式转换
// 支持从解码的私钥无缝创建钱包
package suisecret

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/yyle88/erero"
)

// KeyInfo represents the decoded key information from sui keystore format
// Contains Base64 encoded key with scheme flag and hex private key without flag
// Supports direct wallet creation through GetWallet method
// Compatible with standard sui keytool convert output
//
// KeyInfo 代表从 sui keystore 格式解码后的密钥信息
// 包含带协议标志的 Base64 编码密钥和不带标志的十六进制私钥
// 支持通过 GetWallet 方法直接创建钱包
// 与标准 sui keytool convert 输出兼容
type KeyInfo struct {
	Base64WithFlag string `json:"base64WithFlag"` // Base64 format with scheme flag // 带协议标志的 Base64 格式
	HexWithoutFlag string `json:"hexWithoutFlag"` // Hex private key without flag // 不带标志的十六进制私钥
	Scheme         string `json:"scheme"`         // Key scheme type (ed25519) // 密钥协议类型（ed25519）
}

// GetWallet creates a wallet instance from the decoded private key
// Uses the hex private key stored in HexWithoutFlag field
// Returns new wallet instance or error if wallet creation fails
//
// GetWallet 从解码的私钥创建钱包实例
// 使用存储在 HexWithoutFlag 字段中的十六进制私钥
// 返回新的钱包实例，如果钱包创建失败则返回错误
func (k *KeyInfo) GetWallet() (*suiwallet.Wallet, error) {
	// Create wallet from hex private key
	// 从十六进制私钥创建钱包
	wallet, err := suiwallet.NewWalletV2(k.HexWithoutFlag)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return wallet, nil
}

// SchemeFlagEd25519 represents the scheme flag byte value in sui keystore format
// This constant identifies Ed25519 signature scheme as the first byte in encoded keys
// Value 0x00 indicates Ed25519 cryptographic algorithm
//
// SchemeFlagEd25519 代表 sui keystore 格式中的协议标志字节值
// 此常量标识编码密钥中第一个字节的 Ed25519 签名方案
// 值 0x00 表示 Ed25519 加密算法
const (
	SchemeFlagEd25519 byte = 0x00 // Ed25519 scheme flag value // Ed25519 方案标志值
)

// Decode decodes a Base64-encoded sui keystore key into standard key information
// Accepts Base64 string from sui.keystore file with format [scheme_flag][private_key]
// Returns KeyInfo containing decoded components or error if decode fails
// Compatible with standard sui keytool convert command output
//
// Decode 将 Base64 编码的 sui keystore 密钥解码为标准密钥信息
// 接受来自 sui.keystore 文件的 Base64 字符串，格式为 [协议标志][私钥]
// 返回包含解码组件的 KeyInfo，如果解码失败则返回错误
// 与标准 sui keytool convert 命令输出兼容
func Decode(suiKey string) (*KeyInfo, error) {
	// Decode the Base64 encoded key
	// 解码 Base64 编码的密钥
	keyBytes, err := base64.StdEncoding.DecodeString(suiKey)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// Check minimum length: at least 1 byte for flag + 32 bytes for key
	// 检查最小长度：至少 1 字节标志 + 32 字节密钥
	if len(keyBytes) < 33 {
		return nil, erero.New("invalid key length: must be at least 33 bytes")
	}

	// Extract scheme flag (first byte)
	// 提取协议标志（第一个字节）
	schemeFlag := keyBytes[0]

	// Extract private key (remaining bytes, 32 bytes in most cases)
	// 提取私钥（剩余字节，大多数情况下是 32 字节）
	privateKey := keyBytes[1:]

	// Validate ed25519 scheme support
	// 验证 ed25519 方案支持
	if schemeFlag != SchemeFlagEd25519 {
		return nil, erero.New(fmt.Sprintf("unsupported scheme flag: 0x%02x (just ed25519/0x00 is supported)", schemeFlag))
	}
	schemeName := "ed25519"

	// Build the result
	// 构建结果
	keyInfo := &KeyInfo{
		Base64WithFlag: suiKey,                         // Original input // 原始输入
		HexWithoutFlag: hex.EncodeToString(privateKey), // Hex without flag // 不带标志的十六进制
		Scheme:         schemeName,                     // Scheme name // 协议名称
	}

	return keyInfo, nil
}

// GetSchemeFlag returns the scheme flag byte value corresponding to given scheme name
// Accepts scheme name string and returns corresponding flag byte
// Returns error if scheme is not supported (ed25519 is supported)
//
// GetSchemeFlag 返回与给定方案名称对应的方案标志字节值
// 接受方案名称字符串并返回对应的标志字节
// 如果方案不受支持则返回错误（支持 ed25519）
func GetSchemeFlag(scheme string) (byte, error) {
	// Validate scheme name
	// 验证方案名称
	if scheme != "ed25519" {
		return 0, erero.New("unsupported scheme: " + scheme + " (ed25519 is supported)")
	}
	return SchemeFlagEd25519, nil
}

// Encode encodes a hex private key with scheme flag into Base64 keystore format
// Takes hex-encoded private key string and scheme name as input
// Returns Base64-encoded keystore key or error if encode fails
// This performs the reverse operation of Decode function
//
// Encode 将十六进制私钥和方案标志编码为 Base64 keystore 格式
// 接受十六进制编码的私钥字符串和方案名称作为输入
// 返回 Base64 编码的 keystore 密钥，如果编码失败则返回错误
// 这是 Decode 函数的反向操作
func Encode(privateKeyHex string, scheme string) (string, error) {
	// Decode hex private key
	// 解码十六进制私钥
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", erero.Wro(err)
	}

	// Check key length (should be 32 bytes for most schemes)
	// 检查密钥长度（大多数协议应为 32 字节）
	if len(privateKey) != 32 {
		return "", erero.New("invalid private key length: expected 32 bytes")
	}

	// Get scheme flag
	// 获取协议标志
	flag, err := GetSchemeFlag(scheme)
	if err != nil {
		return "", erero.Wro(err)
	}

	// Combine flag and private key
	// 组合标志和私钥
	keyBytes := make([]byte, 0, 33)
	keyBytes = append(keyBytes, flag)
	keyBytes = append(keyBytes, privateKey...)

	// Encode to Base64
	// 编码为 Base64
	return base64.StdEncoding.EncodeToString(keyBytes), nil
}
