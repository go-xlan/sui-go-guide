// Package main: Demo application to query latest checkpoint sequence number
// Demonstrates sui_getLatestCheckpointSequenceNumber RPC method
// Shows how to retrieve current checkpoint number from blockchain
// Prints checkpoint sequence in JSON format from devnet
//
// main: 查询最新检查点序列号的演示应用程序
// 演示 sui_getLatestCheckpointSequenceNumber RPC 方法
// 展示如何从区块链检索当前检查点编号
// 从开发网以 JSON 格式打印检查点序列
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

	suirpc.SetDebugMode(false)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []any{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](context.Background(), serverUrl, request)
	must.Done(err)
	fmt.Println(neatjsons.S(rpcResponse))
}
