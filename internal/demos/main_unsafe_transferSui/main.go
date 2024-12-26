package main

import (
	"context"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suirpcapi"
	"github.com/go-xlan/sui-go-guide/suirpcmsg"
	"github.com/go-xlan/sui-go-guide/suisigntx"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	// 发起交易的签名者地址
	const signer = "0x353a47f8fedca2d8cd1352222300f06b1f36789a55fffdecc6fe414ee1998969"
	// 接收方地址
	const recipient = "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"
	// 转账金额（最小单位）
	amount := "1000000" // 1 SUI = 1_000_000 微单位
	// 使用的 SUI 对象 ID
	suiObjectID := "0xfc46685ae8893aa647c151f581e60a8549ccb240685b585cdbcf343c4bfd36c9"
	// Gas 预算
	gasBudget := "10000000" // Gas 预算，单位与实际交易成本有关
	// 测试网络
	const serverUrl = "https://fullnode.testnet.sui.io/"
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"

	// 构造 JSON-RPC 请求
	request := &suirpcmsg.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "unsafe_transferSui",
		Params: []any{
			signer,      // 签名者地址
			suiObjectID, // SUI coin 对象 ID
			gasBudget,   // Gas 预算
			recipient,   // 接收方地址
			amount,      // 转账金额（可选）
		},
		ID: 1,
	}

	rpcResponse, err := suirpcapi.SendRpc[suirpcapi.TxBytesMessage](context.Background(), serverUrl, request)
	must.Done(err)
	txBytes := rpcResponse.Result.TxBytes

	{
		res, err := suirpcapi.SuiDryRunTransactionBlock(context.Background(), serverUrl, txBytes)
		must.Done(err)
		fmt.Println(neatjsons.S(res))
	}

	signatures, err := suisigntx.Sign(privateKeyHex, txBytes)
	must.Done(err)
	fmt.Println("signatures", signatures)

	res, err := suirpcapi.SuiExecuteTransactionBlock(context.Background(), serverUrl, txBytes, signatures)
	must.Done(err)
	fmt.Println(neatjsons.S(res))
}
