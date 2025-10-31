package suisecret_test

import (
	"encoding/hex"
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/go-xlan/sui-go-guide/suisecret"
	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

// TestDecode tests basic key decoding operation
// Validates Decode function converts Base64 keystore key into KeyInfo structure
// Verifies scheme detection and hex private key extraction
//
// TestDecode 测试基础密钥解码操作
// 验证 Decode 函数将 Base64 keystore 密钥转换为 KeyInfo 结构
// 验证方案检测和十六进制私钥提取
func TestDecode(t *testing.T) {
	const suiKey = "AAHPc6DmM3+2BWLP/CR/cLLoTtB4SN3o8Z3RNEqmUnuh"

	keyInfo, err := suisecret.Decode(suiKey)
	require.NoError(t, err)
	require.NotNil(t, keyInfo)

	require.Equal(t, "ed25519", keyInfo.Scheme)
	require.Equal(t, suiKey, keyInfo.Base64WithFlag)

	// Verify HexWithoutFlag is valid hex and correct length
	// 验证 HexWithoutFlag 是有效的十六进制且长度正确
	privateKey, err := hex.DecodeString(keyInfo.HexWithoutFlag)
	require.NoError(t, err)
	require.Len(t, privateKey, 32) // Ed25519 private key is 32 bytes
}

// TestEncode tests private key encoding into Base64 keystore format
// Validates Encode function creates valid keystore key from hex private key
// Verifies round-trip conversion between hex and Base64 formats
//
// TestEncode 测试将私钥编码为 Base64 keystore 格式
// 验证 Encode 函数从十六进制私钥创建有效的 keystore 密钥
// 验证十六进制和 Base64 格式之间的往返转换
func TestEncode(t *testing.T) {
	const privateKeyHex = "01cf73a0e6337fb60562cffc247f70b2e84ed07848dde8f19dd1344aa6527ba1"

	suiKey, err := suisecret.Encode(privateKeyHex, "ed25519")
	require.NoError(t, err)
	require.NotEmpty(t, suiKey)

	// Verify by decoding back
	// 通过解码回来验证
	keyInfo, err := suisecret.Decode(suiKey)
	require.NoError(t, err)
	require.Equal(t, "ed25519", keyInfo.Scheme)
	require.Equal(t, privateKeyHex, keyInfo.HexWithoutFlag)
}

// TestDecode_Example demonstrates complete workflow from key decode to wallet address generation
// Decodes Base64 keystore key, creates wallet instance, and verifies generated address
// Validates entire process matches expected cryptographic operations
// This serves as main integration test and usage example
//
// TestDecode_Example 演示从密钥解码到钱包地址生成的完整工作流程
// 解码 Base64 keystore 密钥，创建钱包实例，并验证生成的地址
// 验证整个过程符合预期的加密操作
// 这是主要的集成测试和使用示例
func TestDecode_Example(t *testing.T) {
	const suiKey = "AN81Pxp9PFqCh0SlRMTkfDOP0cSm7U/MxsJiqsWL0KF+"

	// Step 1: Decode to get private key
	// 步骤 1：解码获取私钥
	keyInfo, err := suisecret.Decode(suiKey)
	require.NoError(t, err)
	require.NotNil(t, keyInfo)

	t.Log("=== Key Information ===")
	t.Logf("Scheme: %s", keyInfo.Scheme)
	t.Logf("Base64WithFlag: %s", keyInfo.Base64WithFlag)
	t.Logf("HexWithoutFlag: %s", keyInfo.HexWithoutFlag)

	// Verify all key information matches sui keytool convert output
	// 验证所有密钥信息与 sui keytool convert 输出一致
	require.Equal(t, "ed25519", keyInfo.Scheme)
	require.Equal(t, "AN81Pxp9PFqCh0SlRMTkfDOP0cSm7U/MxsJiqsWL0KF+", keyInfo.Base64WithFlag)
	require.Equal(t, "df353f1a7d3c5a828744a544c4e47c338fd1c4a6ed4fccc6c262aac58bd0a17e", keyInfo.HexWithoutFlag)

	// Step 2: Generate wallet and address from private key
	// 步骤 2：从私钥生成钱包和地址
	wallet, err := keyInfo.GetWallet()
	require.NoError(t, err)
	require.NotNil(t, wallet)

	address := wallet.Address()
	publicKey := hex.EncodeToString(wallet.Public())

	t.Log("=== Wallet Information ===")
	t.Logf("Address: %s", address)
	t.Logf("Public Key (hex): %s", publicKey)

	// Verify the specific address and public key
	// 验证具体地址和公钥
	require.Equal(t, "0x91831805d421e28461324f44f9ba5b629886a36f1015baa8c01f668118098b26", address)
	require.Equal(t, "914706ccdcb5e889c4c33b229edd01ce84d72f54c78e3cfecf3d611782f20e26", publicKey)

	t.Log("\n✅ Successfully converted key and generated address")
}

// TestDecode_Compare verifies implementation output matches reference sui keytool convert command
// Executes real sui keytool convert and compares results with package Decode function
// Tests scheme, Base64 encoding, and hex private key extraction against reference implementation
// Validates wallet address generation from decoded private key
//
// TestDecode_Compare 验证实现输出与参考 sui keytool convert 命令匹配
// 执行真实的 sui keytool convert 并将结果与包的 Decode 函数进行比较
// 针对参考实现测试方案、Base64 编码和十六进制私钥提取
// 验证从解码的私钥生成的钱包地址
func TestDecode_Compare(t *testing.T) {
	// Check if sui command is available
	// 检查 sui 命令是否可用
	path, err := exec.LookPath("sui")
	if err != nil {
		t.Skip("sui is not available on this system, skipping test case")
	}
	t.Log(path)

	const suiKey = "AN81Pxp9PFqCh0SlRMTkfDOP0cSm7U/MxsJiqsWL0KF+"

	// Get output from reference sui keytool convert
	// 从参考 sui keytool convert 获取输出
	output, err := osexec.Exec("sui", "keytool", "convert", suiKey, "--json")
	if err != nil {
		t.Skipf("sui keytool command failed: %v", err)
	}
	t.Log("=== Reference sui keytool output ===")
	t.Log(string(output))

	type suiKeyType struct {
		Bech32WithFlag string `json:"bech32WithFlag"`
		Base64WithFlag string `json:"base64WithFlag"`
		HexWithoutFlag string `json:"hexWithoutFlag"`
		Scheme         string `json:"scheme"`
	}

	var result = &suiKeyType{}
	must.Done(json.Unmarshal(output, result))

	// Get output from our implementation
	// 从我们的实现获取输出
	keyInfo, err := suisecret.Decode(suiKey)
	require.NoError(t, err)
	require.NotNil(t, keyInfo)

	t.Log("=== Our implementation output ===")
	t.Logf("Scheme: %s", keyInfo.Scheme)
	t.Logf("Base64WithFlag: %s", keyInfo.Base64WithFlag)
	t.Logf("HexWithoutFlag: %s", keyInfo.HexWithoutFlag)

	// Compare results
	// 对比结果
	require.Equal(t, result.Scheme, keyInfo.Scheme, "Scheme should match")
	require.Equal(t, result.Base64WithFlag, keyInfo.Base64WithFlag, "Base64WithFlag should match")
	require.Equal(t, result.HexWithoutFlag, keyInfo.HexWithoutFlag, "HexWithoutFlag should match")

	t.Log("\n✅ Implementation matches sui keytool convert output")

	// Step 3: Verify wallet address using the result private key
	// 步骤 3：使用 result 的私钥验证钱包地址
	wallet, err := suiwallet.NewWalletV2(result.HexWithoutFlag)
	require.NoError(t, err)
	require.NotNil(t, wallet)

	address := wallet.Address()
	require.Equal(t, "0x91831805d421e28461324f44f9ba5b629886a36f1015baa8c01f668118098b26", address)

	t.Log("✅ Wallet address matches expected value")
}
