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
	"github.com/yyle88/must/mustslice"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

func main() {
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 发起交易的签名者地址
	const address = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"
	// 要被合并的coin
	primaryCoinObjectID := "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"

	someBigCoins := chooseSomeBigSuiCoins(context.Background(), serverUrl, address)

	var choiceCoin *suiapi.CoinType
	for _, coin := range someBigCoins {
		if coin.CoinObjectId != primaryCoinObjectID {
			choiceCoin = coin
			break
		}
	}
	choiceCoin = must.Nice(choiceCoin) //这里直接检查 Full 不太行，IDE还是会觉得没有检查

	// 使用的 SUI 对象 ID
	coinToMergeObjectID := choiceCoin.CoinObjectId
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关

	//根据你提供的参数描述，unsafe_mergeCoins 方法的参数只允许合并两个 coin，一个作为 primary_coin，另一个作为 coin_to_merge。
	//因此，这个 API 一次只能合并两个 coin。

	// 构造 JSON-RPC 请求
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_mergeCoins",
		Params: []any{
			address,             // 签名者地址
			primaryCoinObjectID, // 要合并到的目标位置
			coinToMergeObjectID, // SUI coin 对象 ID
			nil,                 // Gas 对象，如果没有指定，则为 nil
			gasBudget,           // Gas 预算
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

func chooseSomeBigSuiCoins(ctx context.Context, serverUrl string, address string) []*suiapi.CoinType {
	suiCoins, err := suiapi.GetSuiCoinsInTopPage(ctx, serverUrl, address)
	must.Done(err)

	rand.Shuffle(len(suiCoins), func(i, j int) {
		suiCoins[i], suiCoins[j] = suiCoins[j], suiCoins[i]
	})

	var someBigCoins []*suiapi.CoinType
	for _, coin := range suiCoins {
		suiBalance := rese.V1(strconv.ParseInt(coin.Balance, 10, 64))
		if suiBalance > 1e5 {
			someBigCoins = append(someBigCoins, coin)
		}
	}

	return mustslice.Nice(someBigCoins)
}
