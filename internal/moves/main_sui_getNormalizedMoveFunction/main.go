// Package main: Demo application to query Move function metadata from blockchain
// Demonstrates sui_getNormalizedMoveFunction RPC method with contract introspection
// Shows how to retrieve function signature and parameter information
// Prints function metadata in JSON format from testnet
//
// main: 从区块链查询 Move 函数元数据的演示应用程序
// 演示带有合约自省的 sui_getNormalizedMoveFunction RPC 方法
// 展示如何检索函数签名和参数信息
// 从测试网以 JSON 格式打印函数元数据
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"

	suirpc.SetDebugMode(true)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getNormalizedMoveFunction",
		Params: []any{
			"0x46ed36947b4912ab1d584d9dc5b578f7eb5e271f4ad39541dd189fefad1c34a2", // 包地址
			"math", // 模块名称
			"add",  // 方法名称
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[map[string]interface{}](context.Background(), serverUrl, request))

	fmt.Println(neatjsons.S(rpcResponse))
}
