// Package main: Demo application to split one SUI coin into multiple coins with custom amounts
// Demonstrates unsafe_splitCoin RPC method with amount calculation
// Shows how to divide coin balance into three parts and execute split
// Includes coin selection, amount calculation, and transaction execution
//
// main: 将一个 SUI 代币拆分为多个自定义金额代币的演示应用程序
// 演示带有金额计算的 unsafe_splitCoin RPC 方法
// 展示如何将代币余额分成三部分并执行拆分
// 包括代币选择、金额计算和交易执行
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

	suiCoin := chooseSuiCoin(context.Background(), serverUrl, address)
	zaplog.SUG.Debugln(neatjsons.S(suiCoin))

	// 使用的 SUI 对象 ID
	suiObjectID := suiCoin.CoinObjectId
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	// 如果你在执行 unsafe_splitCoin 时，提供的分割金额 (split_amounts) 的总和不等于原始 SUI coin 的金额，那么交易会失败。
	suiBalance := rese.C1(strconv.ParseInt(suiCoin.Balance, 10, 64))

	// 因此这里，就这样分三瓣儿就行
	part1 := must.Nice(suiBalance / 2)
	part2 := must.Nice((suiBalance - part1) / 2)
	part3 := must.Nice(suiBalance - part1 - part2)

	// 构造 split_amounts 参数，表示将 SUI coin 分裂成多个小的 coin
	splitAmounts := []string{
		strconv.FormatInt(part1, 10),
		strconv.FormatInt(part2, 10),
		strconv.FormatInt(part3, 10),
	} // 示例分割为 3 个 coin

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_splitCoin",
		Params: []any{
			address,      // 签名者地址
			suiObjectID,  // SUI coin 对象 ID
			splitAmounts, // 要分割出的金额
			nil,          // Gas 对象，如果没有指定，则为 nil
			gasBudget,    // Gas 预算
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
