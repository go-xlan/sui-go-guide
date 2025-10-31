package suirpc

import (
	"fmt"
)

// RpcRequest represents JSON-RPC 2.0 request structure
// Contains method name, parameters, and request identifier
// Follows standard JSON-RPC specification with version field
// Used to build blockchain RPC method calls
//
// RpcRequest 表示 JSON-RPC 2.0 请求结构
// 包含方法名称、参数和请求标识符
// 遵循标准 JSON-RPC 规范，包含版本字段
// 用于构建区块链 RPC 方法调用
type RpcRequest struct {
	Jsonrpc string `json:"jsonrpc"` // JSON-RPC version (always "2.0") // JSON-RPC 版本（始终为 "2.0"）
	Method  string `json:"method"`  // RPC method name to invoke // 要调用的 RPC 方法名称
	Params  []any  `json:"params"`  // Method parameters as generic slice // 作为通用切片的方法参数
	ID      int    `json:"id"`      // Request identifier to match responses // 请求标识符以匹配响应
}

// RpcResponse represents JSON-RPC 2.0 response structure with generic result type
// Contains result data typed as RES or error information if call failed
// Generic type parameter enables type-safe response deserialization
// Matches request ID to correlate responses with requests
//
// RpcResponse 表示带有通用结果类型的 JSON-RPC 2.0 响应结构
// 包含类型为 RES 的结果数据，如果调用失败则包含错误信息
// 通用类型参数支持类型安全的响应反序列化
// 匹配请求 ID 以关联响应与请求
type RpcResponse[RES any] struct {
	Jsonrpc string    `json:"jsonrpc"`         // JSON-RPC version (always "2.0") // JSON-RPC 版本（始终为 "2.0"）
	ID      int       `json:"id"`              // Request identifier matching request // 匹配请求的请求标识符
	Result  RES       `json:"result"`          // Response result with generic type // 带有通用类型的响应结果
	Error   *RpcError `json:"error,omitempty"` // Error information if call failed // 调用失败时的错误信息
}

// RpcError represents JSON-RPC 2.0 error information
// Contains error code, message, and optional data
// Implements error interface to work with Go error handling
// Follows standard JSON-RPC error object format
//
// RpcError 表示 JSON-RPC 2.0 错误信息
// 包含错误代码、消息和可选数据
// 实现 error 接口以配合 Go 错误处理
// 遵循标准 JSON-RPC 错误对象格式
type RpcError struct {
	Code    int    `json:"code"`           // Error code identifying error type // 标识错误类型的错误代码
	Message string `json:"message"`        // Human-readable error description // 人类可读的错误描述
	Data    any    `json:"data,omitempty"` // Additional error information // 附加错误信息
}

// Error implements error interface to return formatted error description
// Returns string with code, message, and data fields
// Enables RpcError to work seamlessly with Go error handling
//
// Error 实现 error 接口以返回格式化的错误描述
// 返回包含代码、消息和数据字段的字符串
// 使 RpcError 能够与 Go 错误处理无缝配合
func (rpcError *RpcError) Error() string {
	return fmt.Sprintf("code=%d message=%s data=%v", rpcError.Code, rpcError.Message, rpcError.Data)
}
