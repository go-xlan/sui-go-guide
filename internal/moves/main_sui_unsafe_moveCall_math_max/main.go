// Package main: Demo application to call Move smart contract math max function
// Demonstrates unsafe_moveCall RPC method with mainnet contract
// Shows how to invoke Move module function to find maximum value
// Includes transaction simulation with debug mode enabled
//
// main: 调用 Move 智能合约数学最大值函数的演示应用程序
// 演示在主网合约上使用 unsafe_moveCall RPC 方法
// 展示如何调用 Move 模块函数查找最大值
// 包括启用调试模式的交易模拟
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	const serverUrl = "https://fullnode.mainnet.sui.io/"

	suirpc.SetDebugMode(true)

	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_moveCall",
		Params: []any{
			"0x6356f0141b7a3dd839279ded06ba9e55d928d7019eff01ca981aba44e31afa96", // 签名者地址
			"0xba153169476e8c3114962261d1edc70de5ad9781b83cc617ecc8c1923191cae0", // 包地址
			"math",            // 模块名称
			"max",             // 方法名称
			[]any{},           // 类型参数
			[]any{"10", "20"}, // 方法参数
			nil,               // Gas 对象（可选）
			"75000000",        // Gas 预算，这里注意假如给的太少就会出问题，但给的太多也不利于使用
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[suiapi.TxBytesMessage](context.Background(), serverUrl, request))

	txBytes := rpcResponse.Result.TxBytes

	{
		res, err := suiapi.DryRunTransactionBlock[suiapi.EffectsStatusStatusMessage](context.Background(), serverUrl, txBytes)
		must.Done(err)
		fmt.Println(neatjsons.S(res))
		must.Same(res.Effects.Status.Status, "success")
	}

	// next step: sign
	// next step: send
}
