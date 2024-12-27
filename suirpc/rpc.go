package suirpc

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/zaplog"
)

var debugModeOpen = true

func SetDebugMode(enable bool) {
	debugModeOpen = enable
}

var httpClient *resty.Client
var clientOnce = &sync.Once{}

func newClient() *resty.Client {
	clientOnce.Do(func() {
		httpClient = resty.New().SetDebug(debugModeOpen).SetTimeout(time.Minute)
	})
	return httpClient
}

func SetClient(client *resty.Client) {
	clientOnce.Do(func() {
		httpClient = client
	})
}

func SendRpc[RES any](ctx context.Context, serverUrl string, request *RpcRequest) (rpcResponse *RpcResponse[RES], err error) {
	resp := &RpcResponse[RES]{}
	// 发送 POST 请求
	response, err := newClient().
		R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(resp).
		Post(serverUrl)
	if err != nil {
		return nil, erero.Wro(err)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, erero.New(response.Status())
	}
	if debugModeOpen {
		zaplog.SUG.Debugln("Response Raw:", neatjsons.SxB(response.Body()))
	}
	if resp.Error != nil {
		return nil, erero.Wro(resp.Error.Error())
	}
	if debugModeOpen {
		zaplog.SUG.Debugln("Response Msg:", neatjsons.S(resp))
	}
	return resp, nil
}
