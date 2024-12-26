# sui-go-guide
sui-go-guide: 简单的学习和使用 sui 链的教程。

## README

[ENGLISH-DOC](README.md)

# 安装官方客户端

在mac里使用brew安装：

```bash
brew install sui
```

确认你的客户端版本号：

```bash
sui --version
```

确保版本号和官方文档的相同。

```text
sui 1.39.3-homebrew
```

默认应该是 `mainnet` 的，我们可以尝试切换其他网络

```bash
sui client switch --env devnet
```

```text
Config file ["/Users/admin/.sui/sui_config/client.yaml"] doesn't exist, do you want to connect to a Sui Full node server [y/N]?
```

选择 y 回车

但是会报错，因为 devnet 不行

报错：

```text
Environment config not found for [Some("devnet")], add new environment config using the `sui client new-env` command.
```

不用管这个错误，直接切换到 `testnet` 就行

```bash
sui client switch --env testnet
```

```text
Active environment switched to [testnet]
```

创建钱包地址：

```bash
sui client new-address ed25519
```

这里建议就使用 `ed25519` 这种格式就行，不要给自己增加难度，该项目中的样例都是基于这种格式的。

将会得到钱包

客户端允许多次执行创建钱包，查看钱包的命令是:

```bash
sui client addresses
```

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

这里注意，最后那个 `*` 符号表示那个钱包是默认的。

使用这个命令切换默认钱包：

```bash
sui client switch --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

```text
Active address switched to 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

给某个钱包领 test-coin:

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

```text
Request successful. It can take up to 1 minute to get the coin. Run sui client gas to check your gas coins.
```

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

```text
Request successful. It can take up to 1 minute to get the coin. Run sui client gas to check your gas coins.
```

但也不能太频繁的请求

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

```text
Faucet service received too many requests from this IP address. Please try again after 60 minutes.
```

而且偶尔也会有网络问题

```bash
sui client faucet --address 0x207ed5c0ad36b96c730ed0f71e3c26a0ffb59bc20ab21d08067ca4c035d4d062
```

```text
Faucet request was unsuccessful: 502 Bad Gateway
```

在这里能看到私钥

```bash
cd ~/.sui/sui_config && cat sui.keystore
```

把里面的私钥 [VALUE] 转换为程序能识别的私钥

```bash
sui keytool convert [VALUE]
```

输出：
其中 `hexWithoutFlag` 就是私钥 而 scheme 就是钱包的协议码或者格式码。

通过 `hexWithoutFlag` 就可以用程序发交易。

相见代码样例。

[代码样例](internal/demos)

# 合约的开发教程

## 首先看文档

[官方文档-开始编写合约](https://docs.sui.io/guides/developer/first-app)

确认你的客户端版本号

```bash
sui --version
```

确保版本号和官方文档相同。

```text
sui 1.39.3-homebrew
```

## 如何写合约

[官方文档-创建合约项目](https://docs.sui.io/guides/developer/first-app/write-package)

核心就是这句

```bash
sui move new my_first_package
```

接着编写内容

```bash
pwd
```

```text
/Users/admin/work_atom/move/my_first_package
```

```bash
cd sources
```

```bash
pwd
```

```text
/Users/admin/work_atom/move/my_first_package/sources
```

随便创建几个模块:

```bash
ls
```

```text
github.move             math.move               my_first_package.move
```

其中一个模块的代码内容:

```bash
cat math.move
```

```move
module hello_blockchain::math {

    public fun add(a: u64, b: u64): u64 {
        return a + b
    }

    #[test]
    public fun test_add() {
        let result = add(2, 3);
        assert!(result == 5, 101); 
    }

}
```

回到项目的根目录里：

```bash
pwd
```

```text
/Users/admin/work_atom/move/my_first_package
```

修改合约的配置内容：

```bash
vim Move.toml
```

在这项里面增加 `hello_blockchain = "0x0"` 就行。

```
[addresses]
my_first_package = "0x0"
hello_blockchain = "0x0"
```

因为新增的模块是 `hello_blockchain` 的，而 `my_first_package` 里没有写内容。

## 如何编译合约

[官方文档-编译测试合约](https://docs.sui.io/guides/developer/first-app/build-test)

```bash
sui move build
```

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
```

## 如何测试合约

[官方文档-编译测试合约](https://docs.sui.io/guides/developer/first-app/build-test)

```bash
sui move test
```

```text
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
Running Move unit tests
[ PASS    ] hello_blockchain::math::test_add
[ PASS    ] hello_blockchain::github::test_page
[ PASS    ] hello_blockchain::github::test_page_sui_go_guide
Test result: OK. Total tests: 3; passed: 3; failed: 0
```

## 如何部署合约

[官方文档-发布您的合约](https://docs.sui.io/guides/developer/first-app/publish)

```bash
sui client publish --gas-budget 50000000
```

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
Successfully verified dependencies on-chain against source.
Transaction Digest: Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ
```

这样合约就部署完毕啦。

这里的 `Transaction Digest` 就是区块链的 `txid` / `txHash` 性质的，因此可以通过它查询发布的情况。

[查询发布结果](https://suiscan.xyz/testnet/tx/Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ)

注意不要多次发布因为每次都会发布个新的合约。

## 接着就可以通过代码调用合约

相见代码样例。

[代码样例](internal/demos)
