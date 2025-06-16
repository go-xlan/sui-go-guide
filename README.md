<p align="center">
  <img 
    alt="wojack-cartoon logo" 
    src="assets/wojack-cartoon.jpeg" 
    style="max-height: 500px; width: auto; max-width: 100%;" 
  />
</p>
<h3 align="center">golang-SUI</h3>
<p align="center">create/sign <code>SUI transaction</code> with golang</p>

---

### SUI-Go-Guide: SUI Chain Study and Usage Guide

---

## **README**

[CHINESE-DOCUMENTATION (中文文档)](README.zh.md)

---

# **Installing the Sui Client**

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

- `hexWithoutFlag` is the actual private key in hex style.
- `scheme` represents the wallet's protocol format (`ed25519` in this case).

With the `hexWithoutFlag`, you can perform tasks like signing transactions.

---

## **Code Examples**

Code examples: [internal/demos](internal/demos)

---

## **Contract Guide**

[Contract Guide](SUI-MOVE.md)

---

## DISCLAIMER

Crypto coin, at its core, is nothing but a scam. It thrives on the concept of "air coins"—valueless digital assets—to exploit the hard-earned wealth of ordinary people, all under the guise of innovation and progress. This ecosystem is inherently devoid of fairness or justice.

For the elderly, cryptocurrencies present significant challenges and risks. The so-called "high-tech" façade often excludes them from understanding or engaging with these tools. Instead, they become easy targets for financial exploitation, stripped of the resources they worked a lifetime to accumulate.

The younger generation faces a different but equally insidious issue. By the time they have the opportunity to engage, the early adopters have already hoarded the lion’s share of resources. The system is inherently tilted, offering little chance for new entrants to gain a fair footing.

The idea that cryptocurrencies like BTC, ETH, or TRX could replace global fiat currencies is nothing more than a pipe dream. This notion serves only as the shameless fantasy of early adopters, particularly those from the 1980s generation, who hoarded significant amounts of crypto coin before the general public even had an opportunity to participate.

Ask yourself this: would someone holding thousands, or even tens of thousands, of Bitcoin ever genuinely believe the system is fair? The answer is unequivocally no. These systems were never designed with fairness in mind but rather to entrench the advantages of a select few.

The rise of cryptocurrencies is not the endgame. It is inevitable that new innovations will emerge, replacing these deeply flawed systems. At this moment, my interest lies purely in understanding the underlying technology—nothing more, nothing less.

This project exists solely for the purpose of technical learning and exploration. The author of this project maintains a firm and unequivocal stance of *staunch resistance to cryptocurrencies*.

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/go-xlan/sui-go-guide.svg?variant=adaptive)](https://starchart.cc/go-xlan/sui-go-guide)
