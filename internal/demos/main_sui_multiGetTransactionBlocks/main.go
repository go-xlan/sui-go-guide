package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/simplejsonx/sure/simplejsonm"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	// 开发网络
	const serverUrl = "https://fullnode.devnet.sui.io/"

	suirpc.SetDebugMode(true)

	checkpointNum := mustGetLatestCheckpointNum(context.Background(), serverUrl)
	fmt.Println("Checkpoint-num:", checkpointNum)

	checkpointResult := mustGetCheckpoint(context.Background(), serverUrl, checkpointNum)
	fmt.Println("Checkpoint-res:", neatjsons.S(checkpointResult))

	must.Have(checkpointResult.Transactions)

	var transactionDigests []string = checkpointResult.Transactions

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_multiGetTransactionBlocks",
		Params: []any{
			transactionDigests,
			&SuiTransactionBlockResponseOptions{
				ShowInput:          true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
				ShowRawInput:       true,
			},
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[[]map[string]interface{}](context.Background(), serverUrl, request))
	zaplog.SUG.Debugln(neatjsons.S(rpcResponse))

	for _, result := range rpcResponse.Result {
		simpleJson := simplejsonm.Wrap(result)
		digest := simplejsonm.Extract[string](simpleJson, "digest")
		zaplog.LOG.Info("transaction", zap.String("digest", digest))
	}
}

type SuiTransactionBlockResponseOptions struct {
	ShowInput          bool `json:"showInput,omitempty"`
	ShowEffects        bool `json:"showEffects,omitempty"`
	ShowEvents         bool `json:"showEvents,omitempty"`
	ShowObjectChanges  bool `json:"showObjectChanges,omitempty"`
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty"`
	ShowRawInput       bool `json:"showRawInput,omitempty"`
}

type SuiGetCheckpointResult struct {
	Transactions []string `json:"transactions"`
}

func mustGetCheckpoint(ctx context.Context, serverUrl string, checkpointNum int64) *SuiGetCheckpointResult {
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getCheckpoint",
		Params: []any{
			strconv.FormatInt(checkpointNum, 10),
		},
		ID: 1,
	}

	rpcResponse := rese.P1(suirpc.SendRpc[*SuiGetCheckpointResult](ctx, serverUrl, request))
	return rpcResponse.Result
}

func mustGetLatestCheckpointNum(ctx context.Context, serverUrl string) int64 {
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []any{},
		ID:      1,
	}

	var rpcResponse = rese.P1(suirpc.SendRpc[string](ctx, serverUrl, request))
	checkpointNum := rese.C1(strconv.ParseInt(rpcResponse.Result, 10, 64))
	fmt.Println("Checkpoint-num:", checkpointNum)
	return checkpointNum
}
