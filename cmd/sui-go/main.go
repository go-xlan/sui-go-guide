// Package main: Command-line interface to create and manage SUI blockchain wallets
// Provides wallet generation with random private key creation
// Displays wallet information including public key and address
// Built on cobra CLI framework with colorized terminal output
//
// main: 用于创建和管理 SUI 区块链钱包的命令行界面
// 提供随机私钥创建的钱包生成功能
// 显示包括公钥和地址在内的钱包信息
// 基于 cobra CLI 框架构建，带有彩色终端输出
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/spf13/cobra"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/must"
	"github.com/yyle88/must/muststrings"
	"github.com/yyle88/rese"
)

// main initializes and runs the CLI application
// Sets up root command and registers subcommands
// Executes command based on arguments
//
// main 初始化并运行 CLI 应用程序
// 设置根命令并注册子命令
// 根据参数执行命令
func main() {
	var rootCmd = &cobra.Command{
		Use:   "sui-go",
		Short: "Sui wallet CLI",
		Long:  `Command line application to manage Sui wallet operations including [create wallet]`,
	}

	// 添加子命令
	rootCmd.AddCommand(createWalletCommand())

	// 执行命令 // Execute command
	must.Done(rootCmd.Execute())
}

// createWalletCommand creates the wallet generation subcommand
// Generates random 32-byte private key when invoked
// Calls newWallet to create wallet and display information
// Returns configured cobra command instance
//
// createWalletCommand 创建钱包生成子命令
// 调用时生成随机的 32 字节私钥
// 调用 newWallet 创建钱包并显示信息
// 返回配置的 cobra 命令实例
func createWalletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Sui wallet",
		Long:  `Generate a new Sui wallet with private-key and public-key and wallet-address`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create")
			output := make([]byte, 32) // 32 bytes = 64 hex characters
			rese.C1(rand.Read(output))
			privateKeyHex := hex.EncodeToString(output)
			newWallet(privateKeyHex)
		},
	}
	return cmd
}

// newWallet creates wallet from hex private key and displays details
// Validates private key length is 64 hex characters
// Prints public key and address with color formatting
// Shows complete wallet information to the terminal
//
// newWallet 从十六进制私钥创建钱包并显示详细信息
// 验证私钥长度为 64 个十六进制字符
// 使用颜色格式打印公钥和地址
// 向终端显示完整的钱包信息
func newWallet(privateKeyHex string) {
	fmt.Println(eroticgo.BLUE.Sprint("----"))
	fmt.Println(eroticgo.CYAN.Sprint("RANDOM-PRIVATE:"), privateKeyHex)
	muststrings.Length(privateKeyHex, 64)
	wallet := rese.P1(suiwallet.NewWalletV2(privateKeyHex))
	fmt.Println(eroticgo.BLUE.Sprint("PUBLIC-KEY-HEX:"), hex.EncodeToString(wallet.Public()))
	fmt.Println(eroticgo.PINK.Sprint("WALLET-ADDRESS:"), wallet.Address())
	fmt.Println(eroticgo.BLUE.Sprint("----"))
}
