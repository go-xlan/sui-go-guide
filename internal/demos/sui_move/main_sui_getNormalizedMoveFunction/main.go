package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"

	suirpc.SetDebugMode(true)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getNormalizedMoveFunction",
		Params: []any{
			"0x46ed36947b4912ab1d584d9dc5b578f7eb5e271f4ad39541dd189fefad1c34a2", // 包地址
			"math", // 模块名称
			"add",  // 方法名称
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[map[string]interface{}](context.Background(), serverUrl, request))

	fmt.Println(neatjsons.S(rpcResponse))
}
