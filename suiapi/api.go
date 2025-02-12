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
	// 如果没有显式设置 request_type 参数，那么默认的请求类型会是 WaitForEffectsCert。
	// 但是，如果选项 options.show_events 或 options.show_effects 被设置为 true，则会自动切换为 WaitForLocalExecution 模式。
	// 这里通过显式指定 WaitForLocalExecution 的模式，避免服务端推导，也避免开发者推导。
	// 当配置与约束不一致的时候，请求会报错。
	// 因此建议始终配置，而不是使用默认值。
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
			"WaitForLocalExecution", // WaitForLocalExecution = TransactionEffectsCert + 确认已经执行
		},
		ID: 1,
	}
	response, err := suirpc.SendRpc[RES](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &response.Result, nil
}

func GetSuiCoinsInTopPage(ctx context.Context, serverUrl string, address string) ([]*CoinType, error) {
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
