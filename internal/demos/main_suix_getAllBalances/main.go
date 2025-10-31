// Package main: Demo application to query all coin types and balances owned by address
// Demonstrates suix_getAllBalances RPC method on mainnet
// Shows how to retrieve balance summary across different coin types
// Uses structured logging to display coin types and total balances
//
// main: 查询地址拥有的所有代币类型和余额的演示应用程序
// 演示在主网上使用 suix_getAllBalances RPC 方法
// 展示如何检索不同代币类型的余额摘要
// 使用结构化日志显示代币类型和总余额
package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	const serverUrl = "https://fullnode.mainnet.sui.io/"
	// 要查询余额的地址
	address := "0x2f76f93951df4d4b165a33f41978dfe6040db97ea2dc220602d5c163e9cd3d89"

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getAllBalances",
		Params:  []any{address},
		ID:      1,
	}

	type BalanceItem struct {
		CoinType        string   `json:"coinType"`
		CoinObjectCount int      `json:"coinObjectCount"`
		TotalBalance    string   `json:"totalBalance"`
		LockedBalance   struct{} `json:"lockedBalance"`
	}

	response := rese.P1(suirpc.SendRpc[[]*BalanceItem](context.Background(), serverUrl, request))
	for _, coin := range response.Result {
		zaplog.LOG.Debug("coin", zap.String("balance", coin.TotalBalance), zap.String("coin_type", coin.CoinType))
	}
}
