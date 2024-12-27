### SUI-Go-Guide: SUI Chain Study and Usage Guide

---

## **README**

[Chinese Documentation (中文文档)](README.zh.md)

---

# **Installing the Official Client**

Install the SUI client on macOS using Homebrew:

```bash
brew install sui
```

After installation, check the SUI client version:

```bash
sui --version
```

Ensure the version matches the one mentioned in the official documentation, for example:

```text
sui 1.39.3-homebrew
```

By default, the SUI client connects to the Mainnet. To switch to other networks, such as the Devnet or Testnet, use the commands below.

---

## **Switching Networks**

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

You can skip this step and switch directly to the Testnet:

```bash
sui client switch --env testnet
```

Upon successful switching, you'll see the following output:

```text
Active environment switched to [testnet]
```

---

## **Creating a Wallet Address**

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

## **Claiming Test Tokens**

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

## **Viewing the Private Key**

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

- `hexWithoutFlag` is the actual private key.
- `scheme` represents the wallet's protocol format (`ed25519` in this case).

With the `hexWithoutFlag`, you can perform tasks like signing transactions programmatically.

---

## **Code Examples**

Code examples: [internal/demos](internal/demos)

---

## **Contract Guide**

[Contract Guide](SUI-MOVE.md)

---

## License

`sui-go-guide` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `sui-go-guide`!** 🎉

Give me stars. Thank you!!!
