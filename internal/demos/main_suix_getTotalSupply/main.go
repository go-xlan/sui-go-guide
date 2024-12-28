package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
)

func main() {
	// 主链网络
	const serverUrl = "https://fullnode.mainnet.sui.io/"
	// 代币类型
	coinType := "0x06864a6f921804860930db6ddbe2e16acdf8504495ea7481637a1c8b9a8fe54b::cetus::CETUS"
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getTotalSupply",
		Params:  []any{coinType},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[suiapi.ValueMessage](context.Background(), serverUrl, request)
	must.Done(err)
	value := rpcResponse.Result.Value

	zaplog.SUG.Debugln(value)
}
