package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	const serverUrl = "https://fullnode.devnet.sui.io/"

	suirpc.SetDebugMode(false)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []any{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](context.Background(), serverUrl, request)
	must.Done(err)
	fmt.Println(neatjsons.S(rpcResponse))
}
