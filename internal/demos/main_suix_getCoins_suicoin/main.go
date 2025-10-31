// Package main: Demo application to query SUI coin objects owned by address on testnet
// Demonstrates suix_getCoins RPC method with coin type parameter
// Shows how to filter coins by type to get native SUI tokens
// Prints coin information in JSON format
//
// main: 查询测试网上地址拥有的 SUI 代币对象的演示应用程序
// 演示带有代币类型参数的 suix_getCoins RPC 方法
// 展示如何按类型过滤代币以获取原生 SUI 代币
// 以 JSON 格式打印代币信息
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
	const serverUrl = "https://fullnode.testnet.sui.io/"

	suirpc.SetDebugMode(false)

	// 要查询余额的地址
	address := "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
		Params: []any{
			address,
			"0x2::sui::SUI", //default to 0x2::sui::SUI //因此这里不设置也是可以的
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
