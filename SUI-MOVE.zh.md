# SUI 合约开发教程

## 1. 查看官方文档

[官方文档 - 开始编写合约](https://docs.sui.io/guides/developer/first-app)

在开始前，建议确认你的 SUI 客户端版本号是否与官方文档要求一致：

```bash
sui --version
```

示例版本号：

```text
sui 1.39.3-homebrew
```

确保版本号符合文档要求以避免潜在问题。

---

## 2. 如何编写合约

[官方文档 - 创建合约项目](https://docs.sui.io/guides/developer/first-app/write-package)

### （1）创建合约项目

执行以下命令以生成一个新的合约项目：

```bash
sui move new my_first_package
```

项目生成后，可以查看项目的目录结构：

```bash
pwd
```

示例输出：

```text
/Users/admin/work_atom/move/my_first_package
```

进入项目的 `sources` 文件夹：

```bash
cd sources
pwd
```

示例输出：

```text
/Users/admin/work_atom/move/my_first_package/sources
```

---

### （2）创建模块

在 `sources` 文件夹中可以随意创建模块，例如：

```bash
ls
```

示例输出：

```text
github.move             math.move               my_first_package.move
```

其中 `math.move` 的示例代码如下：

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

---

### （3）配置合约地址

返回项目根目录：

```bash
cd ..
pwd
```

示例输出：

```text
/Users/admin/work_atom/move/my_first_package
```

编辑项目配置文件 `Move.toml`，为新增模块指定地址：

```bash
vim Move.toml
```

在 `[addresses]` 部分增加 `hello_blockchain = "0x0"`：

```toml
[addresses]
my_first_package = "0x0"
hello_blockchain = "0x0"
```

配置完成后，`hello_blockchain` 模块的地址已被正确设置。

---

## 3. 编译合约

[官方文档 - 编译测试合约](https://docs.sui.io/guides/developer/first-app/build-test)

在项目根目录执行以下命令编译合约：

```bash
sui move build
```

示例输出：

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
```

---

## 4. 测试合约

运行以下命令测试合约：

```bash
sui move test
```

示例输出：

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

---

## 5. 部署合约

[官方文档 - 发布您的合约](https://docs.sui.io/guides/developer/first-app/publish)

执行以下命令部署合约到区块链：

```bash
sui client publish --gas-budget 50000000
```

示例输出：

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
Successfully verified dependencies on-chain against source.
Transaction Digest: Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ
```

---

### 部署结果

部署成功后，`Transaction Digest` 就是本次发布的交易哈希（类似 `txid` 或 `txHash`）。你可以通过区块链浏览器查询发布结果：

[示例查询结果](https://suiscan.xyz/testnet/tx/Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ)

**注意**：请避免多次重复发布，否则每次都会生成一个新的合约实例。

---

## 6. 调用合约

合约部署后，即可通过代码调用合约。

[代码样例](internal/demos)

---

## 谢谢

请在 GitHub 上给个 ⭐，感谢支持！！！
