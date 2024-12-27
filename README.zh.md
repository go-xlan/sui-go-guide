### SUI-Go-Guide: SUI链学习与使用指南

---

## **README**

[英文文档（ENGLISH-DOCUMENTATION）](README.md)

---

# **安装官方客户端**

在 macOS 系统上通过 Homebrew 安装 SUI 客户端：

```bash
brew install sui
```

安装完成后，检查 SUI 客户端的版本号：

```bash
sui --version
```

确保版本号与官方文档一致，例如：

```text
sui 1.39.3-homebrew
```

默认情况下，SUI 客户端会连接到主网（Mainnet）。如果需要切换到其他网络（如开发网 Devnet 或测试网 Testnet），可以使用以下命令。

---

## **切换网络**

尝试切换到开发网（Devnet）：

```bash
sui client switch --env devnet
```

如果出现以下提示，选择 **`y`** 并回车：

```text
Config file ["/Users/admin/.sui/sui_config/client.yaml"] doesn't exist, do you want to connect to a Sui Full node server [y/N]?
```

但可能会报错，提示找不到开发网环境配置：

```text
Environment config not found for [Some("devnet")], add new environment config using the `sui client new-env` command.
```

您可以跳过此问题，直接切换到测试网（Testnet）：

```bash
sui client switch --env testnet
```

在成功切换后会输出：

```text
Active environment switched to [testnet]
```

---

## **创建钱包地址**

通过以下命令生成一个新钱包地址：

```bash
sui client new-address ed25519
```

建议使用默认的 **`ed25519`** 格式，简单且兼容性好，本教程中的所有样例均基于此格式。

执行完成后，钱包地址将会生成。

查看已创建的钱包列表：

```bash
sui client addresses
```

输出示例：

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

注意：`*` 标记表示当前的默认钱包地址。

通过以下命令切换默认钱包地址：

```bash
sui client switch --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

切换成功后输出：

```text
Active address switched to 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

---

## **领取测试币**

给现有的钱包领取测试币（Test Coin）：

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

输出示例：

```text
Request successful. It can take up to 1 minute to get the coin. Run sui client gas to check your gas coins.
```

**注意：**
- 测试币请求不能过于频繁，否则可能会出现以下提示：

```text
Faucet service received too many requests from this IP address. Please try again after 60 minutes.
```

- 网络问题可能会导致请求失败：

```text
Faucet request was unsuccessful: 502 Bad Gateway
```

---

## **查看私钥**

私钥文件存储在以下目录下：

```bash
cd ~/.sui/sui_config && cat sui.keystore
```

将文件中的 `[VALUE]` 转换为程序可识别的私钥格式：

```bash
sui keytool convert [VALUE]
```

输出示例：

```text
{
  "hexWithoutFlag": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "scheme": "ed25519"
}
```

- `hexWithoutFlag` 是实际的私钥。
- `scheme` 是钱包的协议格式（此处为 `ed25519`）。

通过 `hexWithoutFlag`，您可以在程序中实现签名交易等功能。

---

## **代码示例**

代码示例：[internal/demos](internal/demos)

--- 

## **合约教程**

[合约教程](SUI-MOVE.zh.md)

---

## 免责声明：

数字货币都是骗局

都是以空气币掠夺平民财富

没有公平正义可言

数字货币对中老年人是极不友好的，因为他们没有机会接触这类披着高科技外衣的割韭菜工具

数字货币对青少年也是极不友好的，因为当他们接触的时候，前面的人已经占据了大量的资源

因此妄图以数字货币，比如稍微主流的 BTC ETH TRX 代替世界货币的操作，都是不可能实现的

都不过是早先持有数字货币的八零后们的无耻幻想

扪心自问，持有几千甚至数万个比特币的人会觉得公平吗，其实不会的

因此未来还会有新事物来代替它们，而我现在也不过只是了解其中的技术，仅此而已。

该项目仅以技术学习和探索为目的而存在。

该项目作者坚定持有“坚决抵制数字货币”的立场。

---

## 许可

`sui-go-guide` 是一个开源项目，发布于 MIT 许可证下。有关更多信息，请参阅 [LICENSE](LICENSE) 文件。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
