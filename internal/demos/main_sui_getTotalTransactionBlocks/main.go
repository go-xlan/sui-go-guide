// Package main: Demo application to query total transaction count from blockchain
// Demonstrates sui_getTotalTransactionBlocks RPC method on mainnet
// Shows how to retrieve cumulative transaction block count
// Logs total transaction number from network
//
// main: 从区块链查询交易总数的演示应用程序
// 演示在主网上使用 sui_getTotalTransactionBlocks RPC 方法
// 展示如何检索累积交易块计数
// 从网络记录总交易数
package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
)

func main() {
	// SUI JSON-RPC API URL
	serverUrl := "https://fullnode.mainnet.sui.io"

	// 构建 JSON-RPC 请求体
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getTotalTransactionBlocks",
		Params:  []interface{}{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](context.Background(), serverUrl, request)
	must.Done(err)
	zaplog.SUG.Debugln(rpcResponse.Result)
}
