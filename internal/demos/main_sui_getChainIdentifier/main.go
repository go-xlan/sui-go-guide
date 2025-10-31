// Package main: Demo application to retrieve blockchain chain identifier
// Demonstrates sui_getChainIdentifier RPC method with devnet
// Shows how to query network chain ID information
// Prints chain identifier to verify network connection
//
// main: 检索区块链链标识符的演示应用程序
// 演示在开发网上使用 sui_getChainIdentifier RPC 方法
// 展示如何查询网络链 ID 信息
// 打印链标识符以验证网络连接
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	const serverUrl = "https://fullnode.devnet.sui.io/"

	suirpc.SetDebugMode(true)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getChainIdentifier",
		Params:  []any{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](context.Background(), serverUrl, request)
	must.Done(err)
	fmt.Println(neatjsons.S(rpcResponse))
}
