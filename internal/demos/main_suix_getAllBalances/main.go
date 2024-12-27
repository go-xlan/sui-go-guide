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
