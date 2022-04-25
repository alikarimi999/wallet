# wallet

A lightweight hierarchical deterministic wallet for handling assets of the shitcoin network.
At the moment it supports [BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki), [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) and [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki) standards

this wallet currently uses  [hdkeychain](https://www.github.com/btcsuite/btcutil/hdkeychain) package which is an implementation of BIP32 for Bitcoin so all addresses that are derived from the mnemonic key are based on bitcoin pay-to-pubkey-hash standard 

## Requirements
[Go](http://golang.org) 1.16 or newer.
## Installation
```bash
$ go install
```
## Getting Started

Run the following command to create a wallet
```bash
$ wallet creat
```
Run the following command to send a transaction to the network
```bash
$ wallet send -from <address> -to <address> -amount 10
```

### Note: 
the wallet must send the transaction to a full node in the network and by default, it sends the transaction to "http://localhost:5000" so if there isn't any node with this address in the network so you have to set flag -node with another HTTP address
### Example:
```bash
$ wallet send -from <address> -to <address> -amount 10 -node http://28fa-5-213-168-103.ngrok.io 
```
## Restore Wallet
Run the following command to restore wallet from a mnemonic phrase
```bash
$ wallet restore -phrase <mnemonic phrase>
#Example 
```bash
$ wallet restore -phrase "regular clever move female attitude chunk rebel hedgehog sugar rain wish stool"
```

## New Account
Based on BIP44 you can derive new Private and Public keys followed by new Addresses from a master key in a [Path level](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki#path-levels)

In this wallet, you can create new accounts in this path m/44'/0'/0'/0 which is a registered derivation path for bitcoin
```bash
$ wallet account
```

Show list of all created accounts
```bash
$ wallet adds   
Account 1: 1Nh5m6rnserLkXvGnHwsPfcooySmkC1o8H
Account 2: 1GVVkEQFrLfR78yWmLsbHb4RGcDd6ZkNBX
Account 3: 15iBSrivmbS33Bb8kfVn6nU19hLVvKsddS
Account 4: 18iqaBPgVbLhmbAXn5vHAckFRxuNojkyad
Account 5: 1NxVzRQmedVkrTDiA8GU2fAoubvNcnL2qE
```

Also, show the balance of all accounts
```bash
$ wallet balance
Account 1: 1Nh5m6rnserLkXvGnHwsPfcooySmkC1o8H	 Balance: 21
Account 2: 1GVVkEQFrLfR78yWmLsbHb4RGcDd6ZkNBX	 Balance: 0
Account 3: 15iBSrivmbS33Bb8kfVn6nU19hLVvKsddS	 Balance: 5
Account 4: 18iqaBPgVbLhmbAXn5vHAckFRxuNojkyad	 Balance: 12
Account 5: 1NxVzRQmedVkrTDiA8GU2fAoubvNcnL2qE	 Balance: 3
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
wallet is available under the MIT license. See the [LICENSE](https://github.com/alikarimi999/wallet/blob/master/LICENSE) file for more info.
