[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/sui-go-guide/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/sui-go-guide/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/sui-go-guide)](https://pkg.go.dev/github.com/go-xlan/sui-go-guide)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/sui-go-guide/main.svg)](https://coveralls.io/github/go-xlan/sui-go-guide?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/go-xlan/sui-go-guide)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/sui-go-guide.svg)](https://github.com/go-xlan/sui-go-guide/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/sui-go-guide)](https://goreportcard.com/report/github.com/go-xlan/sui-go-guide)

<p align="center">
  <img
    alt="wojack-cartoon logo"
    src="assets/wojack-cartoon.jpeg"
    style="max-height: 500px; width: auto; max-width: 100%;"
  />
</p>
<h3 align="center">golang-SUI</h3>
<p align="center">create/sign <code>SUI transaction</code> with golang</p>

# sui-go-guide

Comprehensive Go SDK and tutorial collection to interact with the SUI blockchain.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Overview

**sui-go-guide** is a complete Go SDK and tutorial collection enabling seamless interaction with the SUI blockchain. The package provides pure Go implementations, removing external CLI dependencies and simplifying blockchain integration.

### Key Features

- 🚀 **Pure Go Implementation** - No external CLI dependencies needed
- 🔐 **Complete Wallet Management** - Create, import, and manage SUI wallets
- 📝 **Transaction Signing** - Sign transactions with Ed25519 cryptographic functions
- 🔄 **RPC Client** - Type-safe JSON-RPC communication with SUI nodes
- 📦 **Key Conversion** - Replace `sui keytool convert` with pure Go code
- 📚 **Rich Examples** - 25+ demo apps covering common use cases

---

## Features

- ✅ **Wallet Operations**
  - Generate new wallets with Ed25519 keys
  - Import existing wallets from private keys
  - Derive addresses using Blake2b-256 hashing

- ✅ **Transaction Management**
  - Build and sign transactions
  - Simulate transactions before execution
  - Execute transactions on mainnet/testnet/devnet

- ✅ **Key Conversion**
  - Pure Go implementation of `sui keytool convert`
  - Decode Base64 keystore keys
  - Encode private keys to keystore format

- ✅ **RPC Operations**
  - Query coin balances and metadata
  - Retrieve transaction history
  - Call Move smart contracts
  - Get checkpoint information

- ✅ **Move Contract Interaction**
  - Call Move functions with parameters
  - Query normalized Move function signatures
  - Handle type parameters and complex calls

---

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-xlan/sui-go-guide/suisecret"
    "github.com/go-xlan/sui-go-guide/suiwallet"
    "github.com/go-xlan/sui-go-guide/suiapi"
)

func main() {
    // Decode keystore key (replaces: sui keytool convert)
    keyInfo, _ := suisecret.Decode("AN81Pxp9PFqCh0SlRMTkfDOP0cSm7U/MxsJiqsWL0KF+")

    // Create wallet from decoded key
    wallet, _ := keyInfo.GetWallet()
    fmt.Println("Address:", wallet.Address())

    // Query coin balance
    coins, _ := suiapi.GetSuiCoinsInTopPage(
        context.Background(),
        "https://fullnode.mainnet.sui.io/",
        wallet.Address(),
    )
    fmt.Printf("Found %d coins\n", len(coins))
}
```

---

## Installation

Install the library in your Go project:

```bash
go get github.com/go-xlan/sui-go-guide
```

### Prerequisites

For complete functionality, install the SUI CLI (optional):

```bash
# macOS
brew install sui

# Verify installation
sui --version
```

---

## Examples

The repository includes 25+ demo applications covering common use cases:

- [Basic demos](internal/demos) - Wallet operations, transactions, RPC calls
- [Move contracts](internal/moves) - Smart contract interaction examples

### Example: Convert Keystore Key

```go
import "github.com/go-xlan/sui-go-guide/suisecret"

// Pure Go implementation - replaces: sui keytool convert
keyInfo, err := suisecret.Decode("AAHPc6DmM3+2BWLP/CR/cLLoTtB4SN3o8Z3RNEqmUnuh")
fmt.Println("Private Key:", keyInfo.HexWithoutFlag)
fmt.Println("Scheme:", keyInfo.Scheme)
```

### Example: Create New Wallet

```go
import "github.com/go-xlan/sui-go-guide/suiwallet"

// Generate new wallet with random private key
wallet, err := suiwallet.NewWallet()
fmt.Println("Address:", wallet.Address())
fmt.Println("Private Key:", wallet.PrivateKeyHex())
```

### Example: Query Coin Balance

```go
import "github.com/go-xlan/sui-go-guide/suiapi"

coins, err := suiapi.GetSuiCoinsInTopPage(
    context.Background(),
    "https://fullnode.mainnet.sui.io/",
    "0x...", // wallet address
)
```

---

## SUI CLI Setup Guide

### Installing the Sui Client

Install the SUI client on macOS using Homebrew:

```bash
brew install sui
```

After installation, check the SUI client version:

```bash
sui --version
```

Ensure the version matches the one mentioned in the sui documentation, for example:

```text
sui 1.39.3-homebrew
```

By default, the SUI client connects to the Mainnet. To switch to other networks, such as the Devnet or Testnet, use the commands below.

---

### Switching Networks

To switch to the Devnet:

```bash
sui client switch --env devnet
```

If prompted with the following message, type **`y`** and press Enter:

```text
Config file ["/Users/admin/.sui/sui_config/client.yaml"] doesn't exist, do you want to connect to a Sui Full node server [y/N]?
```

However, you may encounter an error stating the development environment configuration is missing:

```text
Environment config not found for [Some("devnet")], add new environment config using the `sui client new-env` command.
```

You can skip this step and switch to the Testnet:

```bash
sui client switch --env testnet
```

Upon successful switching, you'll see the following output:

```text
Active environment switched to [testnet]
```

---

### Creating a Wallet Address

Generate a new wallet address with the following command:

```bash
sui client new-address ed25519
```

It is recommended to use the default **`ed25519`** format for simplicity and compatibility. All examples in this guide are based on this format.

After execution, a wallet address will be created.

View the list of created wallet addresses:

```bash
sui client addresses
```

Sample output:

```text
╭──────────────────────┬────────────────────────────────────────────────────────────────────┬────────────────╮
│ alias                │ address                                                            │ active address │
├──────────────────────┼────────────────────────────────────────────────────────────────────┼────────────────┤
│ jovial-spinel        │ 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062 │                │
│ wizardly-chrysolite  │ 0x7742b9b86536011afb5f5910bf1365f967fa2d877e2b852f98f34bf8acbc8e44 │ *              │
│ elastic-amethyst     │ 0xbf21da5d7f33b51bece9c5f52771fd349fa8dfd5068ec39869b79933ff949d95 │                │
│ gracious-chrysoberyl │ 0xd544bd5d7516161c74a16a07f6c410b350f3f8e081ebe549b9f3c4451dc00570 │                │
╰──────────────────────┴────────────────────────────────────────────────────────────────────┴────────────────╯
```

Note: The `*` indicates the default wallet address.

Switch the default wallet address using this command:

```bash
sui client switch --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

Successful switching will display:

```text
Active address switched to 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

---

### Claiming Test Tokens

Claim test tokens (Test Coin) for your wallet:

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

Sample output:

```text
Request successful. It can take up to 1 minute to get the coin. Run sui client gas to check your gas coins.
```

**Note:**
- Test token requests cannot be too frequent; otherwise, you may see the following message:

```text
Faucet service received too many requests from this IP address. Please try again after 60 minutes.
```

- Network issues may cause the request to fail:

```text
Faucet request was unsuccessful: 502 Bad Gateway
```

---

### Viewing the Private Key

The private key file is stored in the following directory:

```bash
cd ~/.sui/sui_config && cat sui.keystore
```

Convert the `[VALUE]` in the file to a private key format usable by the program:

```bash
sui keytool convert [VALUE]
```

Sample output:

```text
{
  "hexWithoutFlag": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "scheme": "ed25519"
}
```

- `hexWithoutFlag` is the actual private key in hex format.
- `scheme` represents the wallet's protocol format (`ed25519` in this case).

With the `hexWithoutFlag`, you can perform tasks like signing transactions.

---

## Code Examples

Code examples: [internal/demos](internal/demos)

---

## Key Conversion Package

The `suisecret` package provides a pure Go implementation to convert SUI keystore keys without external dependencies:

```go
import "github.com/go-xlan/sui-go-guide/suisecret"

keyInfo, err := suisecret.Decode(suiKey)
// Replaces: sui keytool convert
```

---

## Contract Development

For Move smart contract development guide, see [SUI-MOVE.md](SUI-MOVE.md)

---

## DISCLAIMER

Crypto coin, at its core, is nothing but a scam. It thrives on the concept of "air coins"—valueless digital assets—to exploit the hard-earned wealth of common people, all under the guise of innovation and advancement. This ecosystem is devoid of fairness and justice.

That cryptocurrencies like BTC, ETH, or TRX could replace global fiat currencies is nothing but a pipe dream. This notion exists as the fantasy of those from the 1980s generation who hoarded amounts of crypto coin before the public had a chance to participate.

Ask this: would someone holding thousands, or tens of thousands, of Bitcoin believe the system is fair? The answer is no. These systems were not designed with fairness in mind but to entrench the advantages of a select few.

The rise of cryptocurrencies is not the endgame. It is inevitable that new innovations will emerge, replacing these flawed systems. At this moment, interest lies in understanding the technical aspects—nothing more, nothing less.

This project exists to support technical education and exploration. The author of this project maintains a firm stance of *staunch resistance to cryptocurrencies*.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-xlan/sui-go-guide.svg?variant=adaptive)](https://starchart.cc/go-xlan/sui-go-guide)
