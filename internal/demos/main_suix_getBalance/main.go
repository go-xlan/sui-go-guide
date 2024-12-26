package main

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/go-xlan/sui-go-guide/suirpcmsg"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type RespType struct {
	Data []struct {
		Balance             string `json:"balance"`
		CoinObjectId        string `json:"coinObjectId"`
		CoinType            string `json:"coinType"`
		Digest              string `json:"digest"`
		PreviousTransaction string `json:"previousTransaction"`
		Version             string `json:"version"`
	} `json:"data"`
	HasNextPage bool   `json:"hasNextPage"`
	NextCursor  string `json:"nextCursor"`
}

func main() {
	address := "0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062"

	request := suirpcmsg.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
		Params: []any{
			address,
		},
		ID: 1,
	}

	rpcResponse := &suirpcmsg.RpcResponse[RespType]{}

	response, err := resty.New().
		SetTimeout(time.Minute).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(rpcResponse).
		Post("https://fullnode.testnet.sui.io/")
	must.Done(err)
	must.Same(http.StatusOK, response.StatusCode())

	zaplog.SUG.Debugln("Response Raw:", neatjsons.SxB(response.Body()))

	must.Null(rpcResponse.Error)

	zaplog.SUG.Debugln("Response Msg:", neatjsons.S(rpcResponse))

	for _, coin := range rpcResponse.Result.Data {
		zaplog.LOG.Debug("coin", zap.String("balance", coin.Balance), zap.String("coin_type", coin.CoinType))
	}
}
