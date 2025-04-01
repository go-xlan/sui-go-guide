package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-xlan/sui-go-guide/suiwallet"
	"github.com/spf13/cobra"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "sui-go",
		Short: "A simple Sui wallet CLI tool",
		Long:  `A command line tool for Sui wallet operations including [create wallet], [convert private-keys] and [sign transaction]`,
	}

	// 添加子命令
	rootCmd.AddCommand(createWalletCommand())
	rootCmd.AddCommand(convertKeysCommand())
	rootCmd.AddCommand(convertKeyCommand())
	rootCmd.AddCommand(signCommand())

	// 执行命令
	must.Done(rootCmd.Execute())
}

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
			showWallet(privateKeyHex)
		},
	}
	return cmd
}

func showWallet(privateKeyHex string) {
	fmt.Println(eroticgo.BLUE.Sprint("----"))
	fmt.Println(eroticgo.RED.Sprint("random-private-key-hex:"), privateKeyHex)
	must.Same(len(privateKeyHex), 64)
	wallet := rese.P1(suiwallet.NewWalletV2(privateKeyHex))
	fmt.Println(eroticgo.BLUE.Sprint("public-key-hex:"), hex.EncodeToString(wallet.Public()))
	fmt.Println(eroticgo.GREEN.Sprint("wallet-address:"), wallet.Address())
	fmt.Println(eroticgo.BLUE.Sprint("----"))
}

func convertKeysCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-keys",
		Short: "Convert private keys from sui.keystore",
		Long:  `Convert private keys from sui.keystore with (sui keytool convert [VALUE])`,
		Run: func(cmd *cobra.Command, args []string) {
			envHome := must.Nice(os.Getenv("HOME"))
			content := rese.A1(os.ReadFile(filepath.Join(envHome, ".sui/sui_config/sui.keystore")))
			fmt.Println(string(content))

			var keys []string
			must.Done(json.Unmarshal(content, &keys))
			must.Have(keys)

			fmt.Println(neatjson.SP2.Soft().S(keys))

			for idx, suiKeyCiphertext := range keys {
				fmt.Println(eroticgo.BLUE.Sprint("----------------"))
				fmt.Println(fmt.Sprintf("(%d/%d)", idx, len(keys)), suiKeyCiphertext)
				fmt.Println(eroticgo.BLUE.Sprint("----------------"))
				result := convertSuiKeyCiphertext(suiKeyCiphertext, true)

				if result.Scheme != "ed25519" {
					continue
				}

				showWallet(result.HexWithoutFlag)
				fmt.Println(eroticgo.BLUE.Sprint("----------------"))
			}
		},
	}
	return cmd
}

type convertSuiKeyCiphertextResultType struct {
	Bech32WithFlag string `json:"bech32WithFlag"`
	Base64WithFlag string `json:"base64WithFlag"`
	HexWithoutFlag string `json:"hexWithoutFlag"`
	Scheme         string `json:"scheme"`
}

func convertSuiKeyCiphertext(suiKeyCiphertext string, debugMode bool) *convertSuiKeyCiphertextResultType {
	if debugMode {
		output := rese.A1(osexec.Exec("sui", "keytool", "convert", suiKeyCiphertext))
		fmt.Println(string(output))
	}
	output := rese.A1(osexec.Exec("sui", "keytool", "convert", suiKeyCiphertext, "--json"))
	if debugMode {
		fmt.Println(string(output))
	}

	var result = &convertSuiKeyCiphertextResultType{}
	must.Done(json.Unmarshal(output, result))
	must.Nice(result.Scheme)
	must.Nice(result.HexWithoutFlag)
	if debugMode {
		fmt.Println(neatjson.SP2.Soft().S(result))
	}
	return result
}

func convertKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert private key from one of key in sui.keystore",
		Long:  `Convert private key from one of key in sui.keystore with (sui keytool convert [VALUE])`,
		Run: func(cmd *cobra.Command, args []string) {
			suiKeyCiphertext := rese.C1(cmd.Flags().GetString("sui_key"))

			fmt.Println(eroticgo.BLUE.Sprint("----------------"))
			result := convertSuiKeyCiphertext(suiKeyCiphertext, true)

			if result.Scheme != "ed25519" {
				return
			}

			showWallet(result.HexWithoutFlag)
			fmt.Println(eroticgo.BLUE.Sprint("----------------"))
		},
	}
	// 设置 flags
	cmd.Flags().StringP("sui_key", "k", "", "Sui keystore key ciphertext")
	must.Done(cmd.MarkFlagRequired("sui_key"))
	return cmd
}

func signCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Sign a message with private key",
		Long:  `Sign a message using the provided private key`,
		Run: func(cmd *cobra.Command, args []string) {
			address := rese.C1(cmd.Flags().GetString("address"))
			rawTransaction := rese.C1(cmd.Flags().GetString("raw_txn"))

			rawTxBytes := rese.A1(base64.StdEncoding.DecodeString(rawTransaction))

			envHome := must.Nice(os.Getenv("HOME"))
			content := rese.A1(os.ReadFile(filepath.Join(envHome, ".sui/sui_config/sui.keystore")))

			var keys []string
			must.Done(json.Unmarshal(content, &keys))
			must.Have(keys)

			for _, suiKeyCiphertext := range keys {
				result := convertSuiKeyCiphertext(suiKeyCiphertext, false)

				if result.Scheme != "ed25519" {
					continue
				}

				wallet := rese.P1(suiwallet.NewWalletV2(result.HexWithoutFlag))
				if wallet.Address() != address {
					continue
				}
				signatureBytes := rese.A1(wallet.Sign(rawTxBytes))

				signature := base64.StdEncoding.EncodeToString(signatureBytes)

				fmt.Println(eroticgo.GREEN.Sprint("--------"))
				fmt.Println(signature)
				fmt.Println(eroticgo.GREEN.Sprint("--------"))
			}
		},
	}
	// 设置 flags
	cmd.Flags().StringP("address", "a", "", "Wallet address to sign with")
	cmd.Flags().StringP("raw_txn", "t", "", "Raw transaction msg to sign")
	must.Done(cmd.MarkFlagRequired("address"))
	must.Done(cmd.MarkFlagRequired("raw_txn"))
	return cmd
}
