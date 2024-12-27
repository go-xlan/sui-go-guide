package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	const serverUrl = "https://fullnode.devnet.sui.io/"

	suirpc.SetDebugMode(false)

	checkpointNum := mustGetCheckpointNum(context.Background(), serverUrl)
	fmt.Println("Checkpoint-num:", checkpointNum)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getCheckpoint",
		Params: []any{
			strconv.FormatInt(checkpointNum, 10),
		},
		ID: 1,
	}

	rpcResponse, err := suirpc.SendRpc[map[string]interface{}](context.Background(), serverUrl, request)
	must.Done(err)
	fmt.Println(neatjsons.S(rpcResponse))
}

func mustGetCheckpointNum(ctx context.Context, serverUrl string) int64 {
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []any{},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[string](ctx, serverUrl, request)
	must.Done(err)
	fmt.Println("Checkpoint-num:", rpcResponse.Result)

	checkpointNum := rese.C1(strconv.ParseInt(rpcResponse.Result, 10, 64))
	fmt.Println("Checkpoint-num:", checkpointNum)
	return checkpointNum
}
