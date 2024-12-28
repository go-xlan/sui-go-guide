package suisigntx_test

import (
	"testing"

	"github.com/go-xlan/sui-go-guide/suisigntx"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	// 交易数据
	const txBytes = "AAACACAgftXArTa5bHMO0PcePCag/7WbwgqyHQgGfKTANdTQYgAIQEIPAAAAAAACAgABAQEAAQECAAABAAA1Okf4/tyi2M0TUiIjAPBrHzZ4mlX//ezG/kFO4ZmJaQH8Rmha6Ik6pkfBUfWB5gqFScyyQGhbWFzbzzQ8S/02yeKFQhEAAAAAIEvQWQTe7kYt5FMQiPWNx3v5vhWc2K7VivFfeaJP/4YnNTpH+P7cotjNE1IiIwDwax82eJpV//3sxv5BTuGZiWnoAwAAAAAAAICWmAAAAAAAAA=="
	// 私钥信息
	const privateKeyHex = "0e51bb6e96264505b7c36c71d6a7f8053ed73b20f6f4476fb4f7877b8934ae6b"

	res, err := suisigntx.Sign(privateKeyHex, txBytes)
	require.NoError(t, err)

	const signatures = "AN7YekuBf2uPBzDALtld1pfUq/R/WIxXe2Z+m7VTzC0sposM2BJDZwtd5bJZw00AYRuN4STT53h8rs0rJJ2swgZc9dk/VcRGRMkypnhuT0HAyWw9A+0IaeOqzAaZq+buog=="
	require.Equal(t, signatures, res)
}
