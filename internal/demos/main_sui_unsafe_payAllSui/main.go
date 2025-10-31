// Package main: Demo application to send all SUI tokens from multiple coins to recipient
// Demonstrates unsafe_payAllSui RPC method with coin consolidation
// Shows how to transfer entire balance from multiple coins in one transaction
// Includes transaction simulation and coin query
//
// main: 从多个代币向接收者发送所有 SUI 代币的演示应用程序
// 演示带有代币合并的 unsafe_payAllSui RPC 方法
// 展示如何在一次交易中从多个代币转移全部余额
// 包括交易模拟和代币查询
package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 使用的 SUI 对象 ID
	const suiObjectID = "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"

	secondCoin := must.Nice(fetchSecondCoin(context.Background(), serverUrl, address, suiObjectID))

	// 接收方地址
	const recipient = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"
	// 使用的 SUI 对象 ID
	inputCoins := []string{
		suiObjectID,
		secondCoin.CoinObjectId,
	}
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	suirpc.SetDebugMode(true)

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_payAllSui",
		Params: []any{
			address,    // 签名者地址
			inputCoins, // SUI coin 对象 ID 列表
			recipient,  // 接收方地址列表
			gasBudget,  // Gas 预算
		},
		ID: 1,
	}

	rpcResponse, err := suirpc.SendRpc[suiapi.TxBytesMessage](context.Background(), serverUrl, request)
	must.Done(err)
	txBytes := rpcResponse.Result.TxBytes

	{
		res, err := suiapi.DryRunTransactionBlock[suiapi.EffectsStatusStatusMessage](context.Background(), serverUrl, txBytes)
		must.Done(err)
		fmt.Println(neatjsons.S(res))
		must.Same(res.Effects.Status.Status, "success")
	}

	// next step: sign
	// next step: send
}

func fetchSecondCoin(ctx context.Context, serverUrl string, address string, suiObjectID string) *suiapi.CoinType {
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

	type GetCoinsResponse struct {
		Data        []*suiapi.CoinType `json:"data"`
		HasNextPage bool               `json:"hasNextPage"`
		NextCursor  string             `json:"nextCursor"`
	}

	rpcResponse := rese.P1(suirpc.SendRpc[GetCoinsResponse](ctx, serverUrl, request))
	for _, coin := range rpcResponse.Result.Data {
		if coin.CoinObjectId == suiObjectID {
			continue
		}
		suiBalance := rese.V1(strconv.ParseInt(coin.Balance, 10, 64))
		if suiBalance > 1e5 {
			return coin
		}
	}
	return nil
}
