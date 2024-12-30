# SUI Smart Contract Development Guide

## 1. Refer to the SUI Documentation

[SUI Documentation - Start Writing a Contract](https://docs.sui.io/guides/developer/first-app)

Before you begin, verify that your SUI client version matches the one required by the sui documentation:

```bash
sui --version
```

Example version:

```text
sui 1.39.3-homebrew
```

Make sure your version matches to avoid potential issues.

---

## 2. How to Write a Contract

[SUI Documentation - Create a Contract Project](https://docs.sui.io/guides/developer/first-app/write-package)

### (1) Create a Contract Project

Run the following command to generate a new contract project:

```bash
sui move new my_first_package
```

After generating the project, you can check the project directory structure:

```bash
pwd
```

Example output:

```text
/Users/admin/work_atom/move/my_first_package
```

Navigate to the `sources` folder:

```bash
cd sources
pwd
```

Example output:

```text
/Users/admin/work_atom/move/my_first_package/sources
```

---

### (2) Create Modules

Inside the `sources` folder, you can create modules as needed. For example:

```bash
ls
```

Example output:

```text
github.move             math.move               my_first_package.move
```

Here's an example of the code inside the `math.move` module:

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

### (3) Configure Contract Addresses

Return to the project root directory:

```bash
cd ..
pwd
```

Example output:

```text
/Users/admin/work_atom/move/my_first_package
```

Edit the project configuration file `Move.toml` to assign an address to the new module:

```bash
vim Move.toml
```

Add `hello_blockchain = "0x0"` under the `[addresses]` section:

```toml
[addresses]
my_first_package = "0x0"
hello_blockchain = "0x0"
```

Once configured, the address for the `hello_blockchain` module is properly set.

---

## 3. Compile the Contract

[SUI Documentation - Build and Test a Contract](https://docs.sui.io/guides/developer/first-app/build-test)

Run the following command in the project root directory to compile the contract:

```bash
sui move build
```

Example output:

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
```

---

## 4. Test the Contract

Run the following command to test the contract:

```bash
sui move test
```

Example output:

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

## 5. Deploy the Contract

[SUI Documentation - Publish Your Contract](https://docs.sui.io/guides/developer/first-app/publish)

Run the following command to deploy the contract to the blockchain:

```bash
sui client publish --gas-budget 50000000
```

Example output:

```text
UPDATING GIT DEPENDENCY https://github.com/MystenLabs/sui.git
INCLUDING DEPENDENCY Sui
INCLUDING DEPENDENCY MoveStdlib
BUILDING my_first_package
Successfully verified dependencies on-chain against source.
Transaction Digest: Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ
```

---

### Deployment Results

After deployment, the `Transaction Digest` serves as the unique transaction hash (similar to `txid` or `txHash`). You can use it to check the deployment status on a blockchain explorer:

[Example Deployment Result](https://suiscan.xyz/testnet/tx/Af7PVu3R3GQsTszsRFfdNjYrLGLZcyioaksHLVDmmKXJ)

**Note**: Avoid deploying the same contract multiple times unnecessarily, as each deployment creates a new contract instance.

---

## 6. Call the Contract

Once the contract is deployed, you can interact with it programmatically.

[Example Code](internal/demos)

---

## Thank you

Give me stars. Thank you!!!
