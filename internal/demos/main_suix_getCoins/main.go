// Package main: Demo application to query coin objects owned by an address
// Demonstrates suix_getCoins RPC method usage with mainnet
// Shows how to retrieve paginated coin data from blockchain
// Prints coin information including balance and object IDs
//
// main: 查询地址拥有的代币对象的演示应用程序
// 演示在主网上使用 suix_getCoins RPC 方法
// 展示如何从区块链检索分页的代币数据
// 打印包括余额和对象 ID 在内的代币信息
package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	const serverUrl = "https://fullnode.mainnet.sui.io/"

	suirpc.SetDebugMode(false)

	// 要查询余额的地址
	address := "0x2f76f93951df4d4b165a33f41978dfe6040db97ea2dc220602d5c163e9cd3d89"

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
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

	rpcResponse := rese.P1(suirpc.SendRpc[GetCoinsResponse](context.Background(), serverUrl, request))
	for _, coin := range rpcResponse.Result.Data {
		fmt.Println(neatjsons.S(coin))
	}
}
