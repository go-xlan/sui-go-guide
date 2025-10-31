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

// TestNewWallet tests wallet creation with random private key generation
// Generates 32 random bytes to create a new wallet instance
// Validates hex encoding produces 64 characters as expected
// Logs generated address to verify wallet creation
//
// TestNewWallet 测试使用随机私钥生成创建钱包
// 生成 32 个随机字节来创建新的钱包实例
// 验证十六进制编码产生预期的 64 个字符
// 记录生成的地址以验证钱包创建
func TestNewWallet(t *testing.T) {
	output := make([]byte, 32) // 32 bytes = 64 hex characters
	rese.C1(rand.Read(output))
	prkHex := hex.EncodeToString(output)
	t.Log("random-private-key-hex:", prkHex)
	require.Len(t, prkHex, 64)

	address := caseNewWallet(t, prkHex)
	t.Log("address:", address)
}

// TestNewWallet_1 tests wallet creation with known private key
// Uses fixed private key to generate wallet and verify address
// Validates address derivation matches expected blockchain address
// Ensures consistent address generation from same private key
//
// TestNewWallet_1 测试使用已知私钥创建钱包
// 使用固定私钥生成钱包并验证地址
// 验证地址派生与预期的区块链地址匹配
// 确保从相同私钥生成一致的地址
func TestNewWallet_1(t *testing.T) {
	const privateKeyHex = "e4c450f61ba740ae8cc1af0b0fb6a135012747b302154410ad635bde12b411c9"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0xe9507e4d5add8cf5570a4d302550fb9a8ad778a101550c1e377ca4e354c404e6", address)
}

// TestNewWallet_2 tests wallet creation with second known private key
// Validates address generation algorithm with different input
// Ensures correct Blake2b-256 hashing and address format
// Verifies implementation consistency across test cases
//
// TestNewWallet_2 测试使用第二个已知私钥创建钱包
// 使用不同输入验证地址生成算法
// 确保正确的 Blake2b-256 哈希和地址格式
// 验证跨测试用例的实现一致性
func TestNewWallet_2(t *testing.T) {
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969", address)
}

// TestNewWallet_3 tests wallet creation with third known private key
// Uses private key with leading zeros to test edge cases
// Validates address generation handles various private key patterns
// Ensures robust implementation across different input formats
//
// TestNewWallet_3 测试使用第三个已知私钥创建钱包
// 使用带前导零的私钥测试边缘情况
// 验证地址生成处理各种私钥模式
// 确保在不同输入格式下的实现稳健性
func TestNewWallet_3(t *testing.T) {
	const privateKeyHex = "00375b3392d6463bb2d1a8e2ae66f1f83a388bc9ab4d4d0b8a378757350b37f7"

	address := caseNewWallet(t, privateKeyHex)
	require.Equal(t, "0x19795c983a7a50c0be1fb4b3d040faa5150d8d9a39fbdb613397a4e74574d91b", address)
}

// caseNewWallet creates wallet from hex private key and logs details
// Helper function shared across wallet creation tests
// Logs public key in both hex and Base64 formats
// Returns generated address and provides testnet explorer link
//
// caseNewWallet 从十六进制私钥创建钱包并记录详细信息
// 在钱包创建测试中共享的辅助函数
// 以十六进制和 Base64 格式记录公钥
// 返回生成的地址并提供测试网浏览器链接
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

// TestWallet_Sign tests Ed25519 signature generation and verification
// Creates wallet from known private key and signs example message
// Verifies signature validation using public key
// Ensures correct signature creation and verification workflow
//
// TestWallet_Sign 测试 Ed25519 签名生成和验证
// 从已知私钥创建钱包并签署示例消息
// 使用公钥验证签名
// 确保正确的签名创建和验证工作流
func TestWallet_Sign(t *testing.T) {
	const privateKeyHex = "00375b3392d6463bb2d1a8e2ae66f1f83a388bc9ab4d4d0b8a378757350b37f7"
	wallet := rese.P1(suiwallet.NewWalletV2(privateKeyHex))

	message := "example"
	signatureBytes, err := wallet.Sign([]byte(message))
	require.NoError(t, err)
	ok := wallet.Verify([]byte(message), signatureBytes)
	require.True(t, ok)
}
