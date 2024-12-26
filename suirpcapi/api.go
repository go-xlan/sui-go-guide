package suirpcapi

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpcmsg"
	"github.com/yyle88/erero"
)

func SuiDryRunTransactionBlock(ctx context.Context, serverUrl string, txBytes string) (map[string]interface{}, error) {
	request := &suirpcmsg.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_dryRunTransactionBlock",
		Params: []any{
			txBytes,
		},
		ID: 1,
	}
	response, err := SendRpc[map[string]interface{}](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return response.Result, nil
}

func SuiExecuteTransactionBlock(ctx context.Context, serverUrl string, txBytes string, signatures string) (map[string]interface{}, error) {
	type TransactionBlockResponseOptions struct {
		ShowInput          bool `json:"showInput"`
		ShowRawInput       bool `json:"showRawInput"`
		ShowEffects        bool `json:"showEffects"`
		ShowEvents         bool `json:"showEvents"`
		ShowObjectChanges  bool `json:"showObjectChanges"`
		ShowBalanceChanges bool `json:"showBalanceChanges"`
		ShowRawEffects     bool `json:"showRawEffects"`
	}

	// 构造模拟调用的 JSON RPC 请求
	request := &suirpcmsg.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_executeTransactionBlock",
		Params: []any{
			txBytes,
			[]string{signatures},
			TransactionBlockResponseOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
				ShowRawEffects:     true,
			},
			"WaitForLocalExecution",
		},
		ID: 1,
	}
	response, err := SendRpc[map[string]interface{}](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return response.Result, nil
}
