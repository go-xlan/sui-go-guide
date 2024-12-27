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
