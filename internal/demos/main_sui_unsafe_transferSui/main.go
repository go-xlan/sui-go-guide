// Package main: Demo application to transfer SUI tokens between addresses
// Demonstrates complete transaction workflow from build to execution
// Shows unsafe_transferSui method usage with testnet
// Includes transaction simulation, signing, and execution steps
//
// main: 在地址之间转账 SUI 代币的演示应用程序
// 演示从构建到执行的完整交易工作流程
// 展示在测试网上使用 unsafe_transferSui 方法
// 包括交易模拟、签名和执行步骤
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/go-xlan/sui-go-guide/suisigntx"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"
	// 接收方地址
	const recipient = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"
	// 转账金额（最小单位）
	amount := "1000000" // 1 SUI = 1_000_000 微单位
	// 使用的 SUI 对象 ID
	suiObjectID := "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_transferSui",
		Params: []any{
			address,     // 签名者地址
			suiObjectID, // SUI coin 对象 ID
			gasBudget,   // Gas 预算
			recipient,   // 接收方地址
			amount,      // 转账金额（可选）
		},
		ID: 1,
	}

	rpcResponse, err := suirpc.SendRpc[suiapi.TxBytesMessage](context.Background(), serverUrl, request)
	must.Done(err)
	txBytes := rpcResponse.Result.TxBytes

	{
		res, err := suiapi.DryRunTransactionBlock[suiapi.EffectsStatusStatusMessage](context.Background(), serverUrl, txBytes)
		must.Done(err)
		fmt.Println(neatjsons.S(res))
		must.Same(res.Effects.Status.Status, "success")
	}

	signatures, err := suisigntx.Sign(privateKeyHex, txBytes)
	must.Done(err)
	fmt.Println("signatures", signatures)

	res, err := suiapi.ExecuteTransactionBlock[suiapi.DigestMessage](context.Background(), serverUrl, txBytes, signatures)
	must.Done(err)
	fmt.Println(neatjsons.S(res))
}
