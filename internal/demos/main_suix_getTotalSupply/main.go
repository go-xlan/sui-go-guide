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
	supplyRequest(context.Background(), serverUrl, "0x2::sui::SUI")
	//https://suiscan.xyz/mainnet/coin/0x06864a6f921804860930db6ddbe2e16acdf8504495ea7481637a1c8b9a8fe54b::cetus::CETUS/txs
	//是这个页面里的 Max supply 信息
	supplyRequest(context.Background(), serverUrl, "0x06864a6f921804860930db6ddbe2e16acdf8504495ea7481637a1c8b9a8fe54b::cetus::CETUS")
	supplyRequest(context.Background(), serverUrl, "0xbff8dc60d3f714f678cd4490ff08cabbea95d308c6de47a150c79cc875e0c7c6::sbox::SBOX")
}

func supplyRequest(ctx context.Context, serverUrl string, coinType string) {
	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getTotalSupply",
		Params:  []any{coinType},
		ID:      1,
	}

	rpcResponse, err := suirpc.SendRpc[suiapi.ValueMessage](ctx, serverUrl, request)
	must.Done(err)
	value := rpcResponse.Result.Value

	zaplog.SUG.Debugln(coinType)
	zaplog.SUG.Debugln(value)
}
