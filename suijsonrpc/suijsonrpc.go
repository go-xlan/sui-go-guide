package suijsonrpc

// JsonRpcRequest represents a JSON-RPC request.
type JsonRpcRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

// JsonRpcResponse represents a JSON-RPC response.
type JsonRpcResponse[RES any] struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Result  RES           `json:"result"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

// JsonRpcError represents a JSON-RPC error-info.
type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
