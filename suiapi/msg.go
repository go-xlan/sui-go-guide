package suiapi

// TxBytesMessage represents transaction bytes response structure
// Contains Base64-encoded transaction data ready to sign
// Returned from transaction build RPC methods
// Used as input to signing operations
//
// TxBytesMessage 表示交易字节响应结构
// 包含准备签名的 Base64 编码交易数据
// 从交易构建 RPC 方法返回
// 用作签名操作的输入
type TxBytesMessage struct {
	TxBytes string `json:"txBytes"` // Base64-encoded transaction bytes // Base64 编码的交易字节
}

// DigestMessage represents transaction digest response structure
// Contains unique transaction hash identifier
// Used to query transaction status and details
// Returned from transaction submission operations
//
// DigestMessage 表示交易摘要响应结构
// 包含唯一的交易哈希标识符
// 用于查询交易状态和详细信息
// 从交易提交操作返回
type DigestMessage struct {
	Digest string `json:"digest"` // Transaction hash digest // 交易哈希摘要
}

// EffectsStatusStatusMessage represents transaction effects status structure
// Contains nested status information from transaction execution
// Used to check if transaction succeeded or failed
// Returned from transaction execution queries
//
// EffectsStatusStatusMessage 表示交易效果状态结构
// 包含交易执行的嵌套状态信息
// 用于检查交易成功或失败
// 从交易执行查询返回
type EffectsStatusStatusMessage struct {
	Effects struct {
		Status struct {
			Status string `json:"status"` // Execution status (success/failure) // 执行状态（成功/失败）
		} `json:"status"` // Status wrapper // 状态包装器
	} `json:"effects"` // Effects wrapper // 效果包装器
}

// CoinType represents coin object information on SUI blockchain
// Contains balance, object ID, and transaction details
// Used to track owned coin objects and their metadata
// Returned from coin query RPC methods
//
// CoinType 表示 SUI 区块链上的代币对象信息
// 包含余额、对象 ID 和交易详细信息
// 用于跟踪拥有的代币对象及其元数据
// 从代币查询 RPC 方法返回
type CoinType struct {
	Balance             string `json:"balance"`             // Coin balance in minimal units // 以最小单位表示的代币余额
	CoinObjectId        string `json:"coinObjectId"`        // Unique object identifier // 唯一对象标识符
	CoinType            string `json:"coinType"`            // Coin type (e.g., 0x2::sui::SUI) // 代币类型（例如 0x2::sui::SUI）
	Digest              string `json:"digest"`              // Object content hash // 对象内容哈希
	PreviousTransaction string `json:"previousTransaction"` // Last transaction affecting coin // 影响代币的最后交易
	Version             string `json:"version"`             // Object version number // 对象版本号
}

// CoinMetadata represents coin metadata information from blockchain
// Contains display information like name, symbol, and decimals
// Used to format coin balances and display coin information
// Returned from coin metadata query methods
//
// CoinMetadata 表示来自区块链的代币元数据信息
// 包含显示信息，如名称、符号和小数位数
// 用于格式化代币余额和显示代币信息
// 从代币元数据查询方法返回
type CoinMetadata struct {
	Decimals    int    `json:"decimals"`    // Number of decimal places // 小数位数
	Description string `json:"description"` // Coin description text // 代币描述文本
	IconUrl     string `json:"iconUrl"`     // Icon image URL // 图标图片 URL
	Id          string `json:"id"`          // Metadata object ID // 元数据对象 ID
	Name        string `json:"name"`        // Coin name // 代币名称
	Symbol      string `json:"symbol"`      // Coin symbol (e.g., SUI) // 代币符号（例如 SUI）
}

// ValueMessage represents simple value response structure
// Contains string value from RPC query methods
// Used to return counts, balances, and other numeric data
// Returned from various query operations
//
// ValueMessage 表示简单值响应结构
// 包含来自 RPC 查询方法的字符串值
// 用于返回计数、余额和其他数值数据
// 从各种查询操作返回
type ValueMessage struct {
	Value string `json:"value"` // String value from query // 来自查询的字符串值
}
