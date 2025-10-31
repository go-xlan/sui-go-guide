// Package main: Demo application to query all coin objects across all types owned by address
// Demonstrates suix_getAllCoins RPC method on mainnet
// Shows how to retrieve complete coin list regardless of coin type
// Uses structured logging to display balance and type information
//
// main: 查询地址拥有的所有类型的所有代币对象的演示应用程序
// 演示在主网上使用 suix_getAllCoins RPC 方法
// 展示如何检索完整的代币列表，不考虑代币类型
// 使用结构化日志显示余额和类型信息
package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	const serverUrl = "https://fullnode.mainnet.sui.io/"
	// 要查询余额的地址
	address := "0x2f76f93951df4d4b165a33f41978dfe6040db97ea2dc220602d5c163e9cd3d89"

	suirpc.SetDebugMode(false)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getAllCoins",
		Params: []any{
			address,
		},
		ID: 1,
	}

	type GetCoinsResponse struct {
		Data        []*suiapi.CoinType `json:"data"`
		HasNextPage bool               `json:"hasNextPage"`
		NextCursor  string             `json:"nextCursor"`
	}

	response := rese.P1(suirpc.SendRpc[GetCoinsResponse](context.Background(), serverUrl, request))
	for _, coin := range response.Result.Data {
		zaplog.LOG.Debug("coin", zap.String("balance", coin.Balance), zap.String("coin_type", coin.CoinType))
	}
}
