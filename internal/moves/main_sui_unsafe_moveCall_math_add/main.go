// Package main: Demo application to call Move smart contract math add function
// Demonstrates unsafe_moveCall RPC method with contract interaction
// Shows how to invoke Move module function with parameters on testnet
// Includes transaction simulation, signing, and execution with result logging
//
// main: 调用 Move 智能合约数学加法函数的演示应用程序
// 演示带有合约交互的 unsafe_moveCall RPC 方法
// 展示如何在测试网上使用参数调用 Move 模块函数
// 包括交易模拟、签名和执行以及结果日志记录
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/go-xlan/sui-go-guide/suisigntx"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_moveCall",
		Params: []any{
			address, // 签名者地址
			"0x46ed36947b4912ab1d584d9dc5b578f7eb5e271f4ad39541dd189fefad1c34a2", // 包地址
			"math",             // 模块名称
			"add",              // 方法名称
			[]any{},            // 类型参数
			[]any{"1", "9998"}, // 方法参数
			nil,                // Gas 对象（可选）
			"75000000",         // Gas 预算，这里注意假如给的太少就会出问题，但给的太多也不利于使用
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[suiapi.TxBytesMessage](context.Background(), serverUrl, request))

	txBytes := rpcResponse.Result.TxBytes

	fmt.Println(txBytes)

	{
		res, err := suiapi.DryRunTransactionBlock[suiapi.EffectsStatusStatusMessage](context.Background(), serverUrl, txBytes)
		must.Done(err)
		fmt.Println(neatjsons.S(res))
		must.Same(res.Effects.Status.Status, "success")
	}

	signatures := rese.C1(suisigntx.Sign(privateKeyHex, txBytes))

	res, err := suiapi.ExecuteTransactionBlock[suiapi.DigestMessage](context.Background(), serverUrl, txBytes, signatures)
	must.Done(err)
	zaplog.SUG.Debugln(neatjsons.S(res))
}
