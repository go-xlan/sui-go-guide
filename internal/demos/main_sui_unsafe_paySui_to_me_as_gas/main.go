// Package main: Demo application to send SUI tokens back to sender using gas
// Demonstrates unsafe_paySui method with self-transfer and gas usage
// Shows how to transfer tokens to own address while paying gas fees
// Includes transaction simulation, signing, and execution on testnet
//
// main: 使用 gas 将 SUI 代币发送回发送者的演示应用程序
// 演示带有自我转账和 gas 使用的 unsafe_paySui 方法
// 展示如何在支付 gas 费用的同时向自己的地址转账代币
// 包括在测试网上的交易模拟、签名和执行
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
	// 使用的 SUI 对象 ID
	const suiObjectID = "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"

	// 接收方地址
	var recipients = []string{
		address,
	}
	// 转账金额（最小单位）
	var amounts = []string{
		"10000000",
	} // 1 SUI = 1_000_000 微单位
	// 使用的 SUI 对象 ID
	inputCoins := []string{
		suiObjectID,
	}
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_paySui",
		Params: []any{
			address,    // 签名者地址
			inputCoins, // SUI coin 对象 ID 列表
			recipients, // 接收方地址列表
			amounts,    // 转账金额列表
			gasBudget,  // Gas 预算
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
