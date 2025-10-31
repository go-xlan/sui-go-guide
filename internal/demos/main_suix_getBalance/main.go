// Package main: Demo application to query coin balance with logging
// Demonstrates suix_getCoins RPC method on testnet
// Shows how to iterate through coin objects and log details
// Uses structured logging to display balance and coin type
//
// main: 使用日志记录查询代币余额的演示应用程序
// 演示在测试网上使用 suix_getCoins RPC 方法
// 展示如何遍历代币对象并记录详细信息
// 使用结构化日志显示余额和代币类型
package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type GetCoinsResponse struct {
	Data        []*suiapi.CoinType `json:"data"`
	HasNextPage bool               `json:"hasNextPage"`
	NextCursor  string             `json:"nextCursor"`
}

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 钱包地址
	const address = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"

	suirpc.SetDebugMode(false)

	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
		Params: []any{
			address,
		},
		ID: 1,
	}

	response := rese.P1(suirpc.SendRpc[GetCoinsResponse](context.Background(), serverUrl, request))
	for _, coin := range response.Result.Data {
		zaplog.LOG.Debug("coin", zap.String("balance", coin.Balance), zap.String("coin_type", coin.CoinType))
	}
}
