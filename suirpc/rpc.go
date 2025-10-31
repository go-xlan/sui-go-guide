// Package suirpc: RPC client implementation with HTTP transport and debug capabilities
// Provides generic RPC request sending with automatic JSON marshaling and response handling
// Supports client configuration, timeout settings, and debug mode logging
// Built on resty HTTP client with connection pooling and retry mechanisms
//
// suirpc: 带有 HTTP 传输和调试功能的 RPC 客户端实现
// 提供具有自动 JSON 序列化和响应处理的通用 RPC 请求发送
// 支持客户端配置、超时设置和调试模式日志记录
// 基于 resty HTTP 客户端构建，包含连接池和重试机制
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

var debugModeOpen = false

// SetDebugMode enables or disables debug logging
// When enabled, logs request and response details
// Accepts boolean flag to control debug state
//
// SetDebugMode 启用或禁用调试日志
// 启用时记录请求和响应详细信息
// 接受布尔标志来控制调试状态
func SetDebugMode(enable bool) {
	debugModeOpen = enable
}

var httpClient *resty.Client
var clientOnce = &sync.Once{}

// newClient creates or returns existing HTTP client instance
// Uses sync.Once to ensure single client creation
// Configures default timeout and debug mode
// Returns configured resty client
//
// newClient 创建或返回现有的 HTTP 客户端实例
// 使用 sync.Once 确保单一客户端创建
// 配置默认超时和调试模式
// 返回配置的 resty 客户端
func newClient() *resty.Client {
	clientOnce.Do(func() {
		httpClient = resty.New().SetDebug(debugModeOpen).SetTimeout(time.Minute)
	})
	return httpClient
}

// SetClient allows custom HTTP client configuration
// Accepts pre-configured resty client instance
// Uses sync.Once to prevent multiple client assignments
//
// SetClient 允许自定义 HTTP 客户端配置
// 接受预配置的 resty 客户端实例
// 使用 sync.Once 防止多次客户端分配
func SetClient(client *resty.Client) {
	clientOnce.Do(func() {
		httpClient = client
	})
}

// SendRpc sends RPC request and deserializes response into generic type
// Accepts context, server URL, and RPC request structure
// Returns typed RPC response or error if request fails
// Supports debug logging when debug mode is enabled
//
// SendRpc 发送 RPC 请求并将响应反序列化为通用类型
// 接受上下文、服务器 URL 和 RPC 请求结构
// 返回类型化的 RPC 响应，如果请求失败则返回错误
// 启用调试模式时支持调试日志记录
func SendRpc[RES any](ctx context.Context, serverUrl string, request *RpcRequest) (rpcResponse *RpcResponse[RES], err error) {
	resp := &RpcResponse[RES]{}

	// Send POST request with JSON body
	// 发送带 JSON 主体的 POST 请求
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

	// Check HTTP status code
	// 检查 HTTP 状态码
	if response.StatusCode() != http.StatusOK {
		return nil, erero.New(response.Status())
	}

	// Log raw response in debug mode
	// 在调试模式下记录原始响应
	if debugModeOpen {
		zaplog.SUG.Debugln("Response Raw:", neatjsons.SxB(response.Body()))
	}

	// Check RPC response errors
	// 检查 RPC 响应错误
	if resp.Error != nil {
		return nil, erero.Wro(resp.Error)
	}

	// Log parsed response in debug mode
	// 在调试模式下记录解析的响应
	if debugModeOpen {
		zaplog.SUG.Debugln("Response Msg:", neatjsons.S(resp))
	}

	return resp, nil
}
