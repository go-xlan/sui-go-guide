package suiapi

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
)

func DryRunTransactionBlock[RES any](ctx context.Context, serverUrl string, txBytes string) (*RES, error) {
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_dryRunTransactionBlock",
		Params: []any{
			txBytes,
		},
		ID: 1,
	}
	response, err := suirpc.SendRpc[RES](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &response.Result, nil
}

func ExecuteTransactionBlock[RES any](ctx context.Context, serverUrl string, txBytes string, signatures string) (*RES, error) {
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
	request := &suirpc.RpcRequest{
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
	response, err := suirpc.SendRpc[RES](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &response.Result, nil
}

func GetSomeSuiCoins(ctx context.Context, serverUrl string, address string) ([]*CoinType, error) {
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

	type Result struct {
		Data []*CoinType `json:"data"`
	}

	response, err := suirpc.SendRpc[Result](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}

	coins := response.Result.Data
	for _, coin := range coins {
		must.Same("0x2::sui::SUI", coin.CoinType)
	}

	return coins, nil
}
