package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
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

	type ResultType struct {
		Data []*suiapi.CoinType `json:"data"`
	}

	rpcResponse, err := suirpc.SendRpc[ResultType](context.Background(), serverUrl, request)
	must.Done(err)
	fmt.Println(neatjsons.S(rpcResponse))

	for _, coin := range rpcResponse.Result.Data {
		fmt.Println(neatjsons.S(coin))
	}
}
