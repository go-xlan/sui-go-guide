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
