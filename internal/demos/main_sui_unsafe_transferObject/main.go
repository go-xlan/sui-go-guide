// Package main: Demo application to transfer object ownership to another address
// Demonstrates unsafe_transferObject RPC method with object transfer
// Shows how to change ownership of blockchain objects
// Includes debug mode and transaction simulation
//
// main: 将对象所有权转移到另一个地址的演示应用程序
// 演示对象转移的 unsafe_transferObject RPC 方法
// 展示如何更改区块链对象的所有权
// 包括调试模式和交易模拟
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 接收方地址
	const recipient = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"
	// 使用的 SUI 对象 ID
	suiObjectID := "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	suirpc.SetDebugMode(true)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_transferObject",
		Params: []any{
			address,     // 签名者地址
			suiObjectID, // SUI coin 对象 ID
			nil,
			gasBudget, // Gas 预算
			recipient, // 接收方地址
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

	// next step: sign
	// next step: send
}
