package suirpc

import (
	"fmt"
)

// RpcRequest represents a JSON-RPC request.
type RpcRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

// RpcResponse represents a JSON-RPC response.
type RpcResponse[RES any] struct {
	Jsonrpc string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  RES       `json:"result"`
	Error   *RpcError `json:"error,omitempty"`
}

// RpcError represents a JSON-RPC error-info.
type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Error implement the error interface{} to return wrong-reason
func (rpcError *RpcError) Error() string {
	return fmt.Sprintf("code=%d message=%s data=%v", rpcError.Code, rpcError.Message, rpcError.Data)
}
