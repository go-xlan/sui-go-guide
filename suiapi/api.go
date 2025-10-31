// Package suiapi: High-level API wrappers with SUI blockchain RPC methods
// Provides convenient functions to interact with SUI blockchain nodes
// Supports transaction execution, simulation, and coin balance queries
// Built on generic RPC client with type-safe response handling
//
// suiapi: 高层 API 包装器，包含 SUI 区块链 RPC 方法
// 提供便捷的函数与 SUI 区块链节点交互
// 支持交易执行、模拟和代币余额查询
// 基于通用 RPC 客户端构建，具有类型安全的响应处理
package suiapi

import (
	"context"

	"github.com/go-xlan/sui-go-guide/suirpc"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
)

// DryRunTransactionBlock simulates transaction execution without committing to blockchain
// Accepts context, server URL, and Base64-encoded transaction bytes
// Returns typed response with effects or error if simulation fails
// Useful to validate transactions before actual execution
//
// DryRunTransactionBlock 模拟交易执行而不提交到区块链
// 接受上下文、服务器 URL 和 Base64 编码的交易字节
// 返回带有效果的类型化响应，如果模拟失败则返回错误
// 在实际执行前验证交易非常有用
func DryRunTransactionBlock[RES any](ctx context.Context, serverUrl string, txBytes string) (*RES, error) {
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_dryRunTransactionBlock",
		Params: []any{
			txBytes,
		},
		ID: 1,
	}
	response, err := suirpc.SendRpc[RES](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &response.Result, nil
}

// ExecuteTransactionBlock executes signed transaction on blockchain
// Accepts context, server URL, transaction bytes, and signature string
// Returns typed response with execution results or error if execution fails
// Uses WaitForLocalExecution mode to ensure transaction confirmation
//
// ExecuteTransactionBlock 在区块链上执行已签名的交易
// 接受上下文、服务器 URL、交易字节和签名字符串
// 返回带有执行结果的类型化响应，如果执行失败则返回错误
// 使用 WaitForLocalExecution 模式确保交易确认
func ExecuteTransactionBlock[RES any](ctx context.Context, serverUrl string, txBytes string, signatures string) (*RES, error) {
	type TransactionBlockResponseOptions struct {
		ShowInput          bool `json:"showInput"`          // Include transaction input data // 包含交易输入数据
		ShowRawInput       bool `json:"showRawInput"`       // Include raw input bytes // 包含原始输入字节
		ShowEffects        bool `json:"showEffects"`        // Include transaction effects // 包含交易效果
		ShowEvents         bool `json:"showEvents"`         // Include emitted events // 包含发出的事件
		ShowObjectChanges  bool `json:"showObjectChanges"`  // Include object changes // 包含对象更改
		ShowBalanceChanges bool `json:"showBalanceChanges"` // Include balance changes // 包含余额更改
		ShowRawEffects     bool `json:"showRawEffects"`     // Include raw effects data // 包含原始效果数据
	}

	// Build JSON RPC request with comprehensive response options
	// Explicitly set WaitForLocalExecution mode to ensure transaction confirmation
	// Without explicit request_type, default would be WaitForEffectsCert
	// However, setting show_events or show_effects to true auto-switches to WaitForLocalExecution
	// This explicit configuration avoids server-side or developer inference
	// Request fails if configuration conflicts with constraints
	// Therefore, recommend to set options instead of using defaults
	//
	// 构建包含完整响应选项的 JSON RPC 请求
	// 显式设置 WaitForLocalExecution 模式以确保交易确认
	// 如果不显式设置 request_type，默认是 WaitForEffectsCert
	// 但如果将 show_events 或 show_effects 设置为 true，会自动切换到 WaitForLocalExecution
	// 这个显式配置避免了服务端或开发者的推断
	// 如果配置与约束冲突，请求会失败
	// 因此建议设置选项而不是使用默认值
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "sui_executeTransactionBlock",
		Params: []any{
			txBytes,
			[]string{signatures},
			TransactionBlockResponseOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
				ShowRawEffects:     true,
			},
			"WaitForLocalExecution", // WaitForLocalExecution = TransactionEffectsCert + execution confirmed // WaitForLocalExecution = TransactionEffectsCert + 确认已执行
		},
		ID: 1,
	}
	response, err := suirpc.SendRpc[RES](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return &response.Result, nil
}

// GetSuiCoinsInTopPage retrieves SUI coins owned by address in first page
// Accepts context, server URL, and wallet address string
// Returns slice of coin objects or error if query fails
// Verifies each coin type matches SUI native token
//
// GetSuiCoinsInTopPage 检索地址在第一页拥有的 SUI 代币
// 接受上下文、服务器 URL 和钱包地址字符串
// 返回代币对象切片，如果查询失败则返回错误
// 验证每个代币类型与 SUI 原生代币匹配
func GetSuiCoinsInTopPage(ctx context.Context, serverUrl string, address string) ([]*CoinType, error) {
	// Build JSON-RPC request to query SUI coins
	// 构建 JSON-RPC 请求以查询 SUI 代币
	request := &suirpc.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "suix_getCoins",
		Params: []any{
			address,
			"0x2::sui::SUI", // Default to 0x2::sui::SUI (parameter can be omitted) // 默认为 0x2::sui::SUI（参数可以省略）
		},
		ID: 1,
	}

	type Result struct {
		Data []*CoinType `json:"data"` // Coin list in current page // 当前页面的代币列表
	}

	response, err := suirpc.SendRpc[Result](ctx, serverUrl, request)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// Extract coins and validate coin types
	// 提取代币并验证代币类型
	coins := response.Result.Data
	for _, coin := range coins {
		must.Same("0x2::sui::SUI", coin.CoinType)
	}

	return coins, nil
}
