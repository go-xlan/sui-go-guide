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
<p align="center">使用 golang 创建/签名 <code>SUI 交易</code></p>

# sui-go-guide

SUI 区块链交互的完整 Go SDK 和教程集合。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 概述

**sui-go-guide** 是完整的 Go SDK 和教程集合，实现与 SUI 区块链的无缝交互。该包提供纯 Go 实现，移除外部 CLI 依赖，简化区块链集成。

### 核心特性

- 🚀 **纯 Go 实现** - 无需外部 CLI 依赖
- 🔐 **完整的钱包管理** - 创建、导入和管理 SUI 钱包
- 📝 **交易签名** - 使用 Ed25519 加密函数签名交易
- 🔄 **RPC 客户端** - 与 SUI 节点进行类型安全的 JSON-RPC 通信
- 📦 **密钥转换** - 使用纯 Go 代码替代 `sui keytool convert`
- 📚 **丰富示例** - 25+ 个演示应用涵盖常见用例

---

## 功能特性

- ✅ **钱包操作**
  - 使用 Ed25519 密钥生成新钱包
  - 从私钥导入现有钱包
  - 使用 Blake2b-256 哈希派生地址

- ✅ **交易管理**
  - 构建和签名交易
  - 执行前模拟交易
  - 在主网/测试网/开发网上执行交易

- ✅ **密钥转换**
  - 纯 Go 实现 `sui keytool convert`
  - 解码 Base64 密钥库密钥
  - 将私钥编码成密钥库格式

- ✅ **RPC 操作**
  - 查询代币余额和元数据
  - 检索交易历史
  - 调用 Move 智能合约
  - 获取检查点信息

- ✅ **Move 合约交互**
  - 调用带参数的 Move 函数
  - 查询标准化 Move 函数签名
  - 处理类型参数和复杂调用

---

## 快速开始

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
    // 解码密钥库密钥（替代：sui keytool convert）
    keyInfo, _ := suisecret.Decode("AN81Pxp9PFqCh0SlRMTkfDOP0cSm7U/MxsJiqsWL0KF+")

    // 从解码的密钥创建钱包
    wallet, _ := keyInfo.GetWallet()
    fmt.Println("Address:", wallet.Address())

    // 查询代币余额
    coins, _ := suiapi.GetSuiCoinsInTopPage(
        context.Background(),
        "https://fullnode.mainnet.sui.io/",
        wallet.Address(),
    )
    fmt.Printf("Found %d coins\n", len(coins))
}
```

---

## 安装

在 Go 项目中安装库：

```bash
go get github.com/go-xlan/sui-go-guide
```

### 前置条件

完整功能需要安装 SUI CLI（可选）：

```bash
# macOS
brew install sui

# 验证安装
sui --version
```

---

## SUI CLI 设置指南

### 安装 SUI 客户端

在 macOS 系统上通过 Homebrew 安装 SUI 客户端：

```bash
brew install sui
```

安装完成后，检查 SUI 客户端的版本号：

```bash
sui --version
```

确保版本号与 SUI 文档一致，例如：

```text
sui 1.39.3-homebrew
```

默认情况下，SUI 客户端连接到主网。要切换到其他网络（如开发网或测试网），请使用以下命令。

---

### 切换网络

尝试切换到开发网（Devnet）：

```bash
sui client switch --env devnet
```

如果提示以下信息，输入 **`y`** 并按 Enter：

```text
Config file ["/Users/admin/.sui/sui_config/client.yaml"] doesn't exist, do you want to connect to a Sui Full node server [y/N]?
```

但可能会遇到开发环境配置缺失的错误：

```text
Environment config not found for [Some("devnet")], add new environment config using the `sui client new-env` command.
```

可以跳过这一步，切换到测试网：

```bash
sui client switch --env testnet
```

成功切换后，会看到以下输出：

```text
Active environment switched to [testnet]
```

---

### 创建钱包地址

通过以下命令生成一个新钱包地址：

```bash
sui client new-address ed25519
```

建议使用默认的 **`ed25519`** 格式，以获得简单性和兼容性。本指南的所有示例都基于此格式。

执行后，将创建一个钱包地址。

查看已创建的钱包地址列表：

```bash
sui client addresses
```

示例输出：

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

注意：`*` 表示默认钱包地址。

使用以下命令切换默认钱包地址：

```bash
sui client switch --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

成功切换将显示：

```text
Active address switched to 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

---

### 领取测试币

给钱包领取测试币（Test Coin）：

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

示例输出：

```text
Request successful. It can take up to 1 minute to get the coin. Run sui client gas to check your gas coins.
```

**注意：**
- 测试币申请不能太频繁，否则会看到以下消息：

```text
Faucet service received too many requests from this IP address. Please try again after 60 minutes.
```

- 网络问题可能导致请求失败：

```text
Faucet request was unsuccessful: 502 Bad Gateway
```

---

### 查看私钥

私钥文件存储在以下目录下：

```bash
cd ~/.sui/sui_config && cat sui.keystore
```

将文件中的 `[VALUE]` 转换成程序可用的私钥格式：

```bash
sui keytool convert [VALUE]
```

示例输出：

```text
{
  "hexWithoutFlag": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "scheme": "ed25519"
}
```

- `hexWithoutFlag` 是十六进制格式的实际私钥。
- `scheme` 表示钱包的协议格式（本例中是 `ed25519`）。

通过 `hexWithoutFlag`，可以执行签名交易等任务。

---

## 代码示例

代码示例：[internal/demos](internal/demos)

---

## 密钥转换包

`suisecret` 包提供纯 Go 实现来转换 SUI keystore 密钥，无需外部依赖：

```go
import "github.com/go-xlan/sui-go-guide/suisecret"

keyInfo, err := suisecret.Decode(suiKey)
// 替代：sui keytool convert
```

---

## 合约开发

Move 智能合约开发指南，参见 [SUI-MOVE.md](SUI-MOVE.md)

---

## 免责声明

加密货币本质上就是骗局。它靠"空气币"这种无价值的数字资产，打着创新和进步的旗号，剥削普通人辛苦积攒的财富。这个生态系统缺乏公平和正义。

认为 BTC、ETH 或 TRX 这类加密货币能取代全球法定货币，不过是痴人说梦。这种想法只是那些 80 后一代早期参与者的幻想，他们在公众有机会参与之前就囤积了大量加密货币。

试问：持有成千上万甚至数万比特币的人，会真心认为这个系统公平吗？答案显然是否定的。这些系统从设计之初就不是为了公平，而是为了巩固少数人的优势。

加密货币的兴起不是终点。必然会有新的创新出现，取代这些有缺陷的系统。此刻，我的兴趣纯粹在于理解技术层面——仅此而已。

本项目的存在是为了支持技术教育和探索。本项目作者坚持*坚决抵制加密货币*的立场。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-xlan/sui-go-guide.svg?variant=adaptive)](https://starchart.cc/go-xlan/sui-go-guide)
