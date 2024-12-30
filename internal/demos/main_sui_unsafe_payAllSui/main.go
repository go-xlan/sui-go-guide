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
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"
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
