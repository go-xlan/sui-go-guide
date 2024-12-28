package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
)

func main() {
	// SUI JSON-RPC API URL
	serverUrl := "https://fullnode.mainnet.sui.io"

	// 构建 JSON-RPC 请求体
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getTotalTransactionBlocks",
		Params:  []interface{}{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](context.Background(), serverUrl, request)
	must.Done(err)
	zaplog.SUG.Debugln(rpcResponse.Result)
}
