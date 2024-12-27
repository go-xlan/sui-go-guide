package main

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	// 主链网络
	const serverUrl = "https://fullnode.mainnet.sui.io/"

	// 代币类型
	coinMetadata0 := rese.P1(requestCoinMetadata(context.Background(), serverUrl, "0x810e52b7e3ba96cc82170533405ac1b5d1f7346947b51b4caa9d7f6af2fa7b52::sui::SUI"))
	coinMetadata1 := rese.P1(requestCoinMetadata(context.Background(), serverUrl, "0xc060006111016b8a020ad5b33834984a437aaa7d3c74c18e09a95d48aceab08c::coin::COIN"))
	coinMetadata2 := rese.P1(requestCoinMetadata(context.Background(), serverUrl, "0x5a09e3c94f02d0d3d75ca22b7d7843bef4023c89f1ed2a105e9f6f36c0f930a7::asui::ASUI"))

	zaplog.SUG.Debugln(neatjsons.S([]*suiapi.CoinMetadata{
		coinMetadata0,
		coinMetadata1,
		coinMetadata2,
	}))
}

func requestCoinMetadata(ctx context.Context, serverUrl string, coinType string) (*suiapi.CoinMetadata, error) {
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoinMetadata",
		Params:  []any{coinType},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[suiapi.CoinMetadata](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &rpcResponse.Result, nil
}
