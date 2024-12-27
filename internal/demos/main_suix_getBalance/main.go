package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type GetCoinsResponse struct {
	Data []struct {
		Balance             string `json:"balance"`
		CoinObjectId        string `json:"coinObjectId"`
		CoinType            string `json:"coinType"`
		Digest              string `json:"digest"`
		PreviousTransaction string `json:"previousTransaction"`
		Version             string `json:"version"`
	} `json:"data"`
	HasNextPage bool   `json:"hasNextPage"`
	NextCursor  string `json:"nextCursor"`
}

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 钱包地址
	const address = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"

	request := suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
		Params: []any{
			address,
		},
		ID: 1,
	}

	response := rese.P1(suirpc.SendRpc[GetCoinsResponse](context.Background(), serverUrl, &request))
	for _, coin := range response.Result.Data {
		zaplog.LOG.Debug("coin", zap.String("balance", coin.Balance), zap.String("coin_type", coin.CoinType))
	}
}
