// Package main: Demo application to split one SUI coin into equal amounts
// Demonstrates unsafe_splitCoinEqual RPC method with count-based split
// Shows how to divide coin balance into equal parts automatically
// Includes coin selection and equal distribution transaction execution
//
// main: 将一个 SUI 代币平均拆分为多个相等金额的演示应用程序
// 演示基于计数的 unsafe_splitCoinEqual RPC 方法
// 展示如何自动将代币余额平均分成多个部分
// 包括代币选择和平均分配交易执行
package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/go-xlan/sui-go-guide/suiapi"
	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/go-xlan/sui-go-guide/suisigntx"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"
	// 要分割的块数
	const splitCount = 3

	suiCoin := chooseSuiCoin(context.Background(), serverUrl, address)
	zaplog.SUG.Debugln(neatjsons.S(suiCoin))

	// 使用的 SUI 对象 ID
	suiObjectID := suiCoin.CoinObjectId
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_splitCoinEqual",
		Params: []any{
			address,                  // 签名者地址
			suiObjectID,              // SUI coin 对象 ID
			strconv.Itoa(splitCount), // 要分割出的金额
			nil,                      // Gas 对象，如果没有指定，则为 nil
			gasBudget,                // Gas 预算
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

	signatures, err := suisigntx.Sign(privateKeyHex, txBytes)
	must.Done(err)
	fmt.Println("signatures", signatures)

	res, err := suiapi.ExecuteTransactionBlock[suiapi.DigestMessage](context.Background(), serverUrl, txBytes, signatures)
	must.Done(err)
	fmt.Println(neatjsons.S(res))
}

func chooseSuiCoin(ctx context.Context, serverUrl string, address string) *suiapi.CoinType {
	suiCoins, err := suiapi.GetSuiCoinsInTopPage(ctx, serverUrl, address)
	must.Done(err)

	rand.Shuffle(len(suiCoins), func(i, j int) {
		suiCoins[i], suiCoins[j] = suiCoins[j], suiCoins[i]
	})

	var choiceCoin *suiapi.CoinType
	for _, coin := range suiCoins {
		//这里确实是会有一个 "0" coin。//您遇到的这种情况可能是 SUI 链设计中的正常行为。分割操作不会销毁原始 coin，而是将其金额设置为 0，并保持该 coin 作为一个占位对象。
		suiBalance := rese.V1(strconv.ParseInt(coin.Balance, 10, 64))
		if suiBalance > 1e8 {
			choiceCoin = coin
			break
		}
	}
	return must.Nice(choiceCoin)
}
