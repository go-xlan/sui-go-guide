package suirpcapi

import (
	"context"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/go-xlan/sui-go-guide/suirpcmsg"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/zaplog"
)

func SendRpc[RES any](ctx context.Context, serverUrl string, request *suirpcmsg.RpcRequest) (*suirpcmsg.RpcResponse[RES], error) {
	resp := &suirpcmsg.RpcResponse[RES]{}
	// 发送 POST 请求
	response, err := resty.New().
		SetTimeout(time.Minute).
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
	zaplog.SUG.Debugln("Response Raw:", neatjsons.SxB(response.Body()))
	if resp.Error != nil {
		return nil, erero.Wro(resp.Error.Error())
	}
	zaplog.SUG.Debugln("Response Msg:", neatjsons.S(resp))
	return resp, nil
}
