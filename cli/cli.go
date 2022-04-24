package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/alikarimi999/wallet/config"
	network "github.com/alikarimi999/wallet/network/http"
	"github.com/alikarimi999/wallet/utils"
	"github.com/alikarimi999/wallet/wallet"
	"github.com/tyler-smith/go-bip39"
)

type Commandline struct{}

func (cli *Commandline) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		runtime.Goexit()
	}
}

func (cli *Commandline) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println(" create - Creat new wallet")
	fmt.Println(" send - Send Transaction to a Full Node")
	fmt.Println(" balance - Show balance of all Accounts")
	fmt.Println(" account - Creat new account")
	fmt.Println(" adds - List of all created addresses")
	fmt.Println(" restore - Restore Walet with BIP39 Phrase Key")
	fmt.Println(" phrase - Access to BIP39 menmonic Phrase Key")
}

func (cli *Commandline) Run() {
	cli.ValidateArgs()

	create := flag.NewFlagSet("create", flag.ExitOnError)
	send := flag.NewFlagSet("send", flag.ExitOnError)
	balance := flag.NewFlagSet("balance", flag.ExitOnError)
	account := flag.NewFlagSet("account", flag.ExitOnError)
	restore := flag.NewFlagSet("restore", flag.ExitOnError)
	phrase := flag.NewFlagSet("phrase", flag.ExitOnError)
	adds := flag.NewFlagSet("adds", flag.ExitOnError)

	bnode := balance.String("node", "http://localhost:5000", "Address of Full Node that send back Account Balance")

	from := send.String("from", "", "Sender Address")
	to := send.String("to", "", "Reciever Address")
	amount := send.Int("amount", 0, "")
	node := send.String("node", "http://localhost:5000", "Address of Full Node that recieve and proccess transaction")

	mnemonic := restore.String("phrase", "", "BIP39 menmonic Phrase Key")

	switch os.Args[1] {
	case "create":
		err := create.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "send":
		err := send.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "balance":
		err := balance.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "account":
		err := account.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "restore":
		err := restore.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "phrase":
		err := phrase.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "adds":
		err := adds.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	default:
		cli.PrintUsage()
	}

	if create.Parsed() {
		if config.WalletExist() {
			fmt.Printf("there is a wallet with name %s in %s you can't have multiple wallet in same time\n", config.DefaultName, config.WalletPath())

			os.Exit(0)
		}

		ent, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatal(err)
		}
		mnemonic, err := bip39.NewMnemonic(ent)
		if err != nil {
			log.Fatal(err)
		}
		w := wallet.NewWallet(mnemonic)
		c := config.NewConfig()
		c.SaveWallet(w)
		err = c.SaveConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Wallet created successfully")
	}

	if send.Parsed() {
		fmt.Println(balance.Parsed())

		if !config.WalletExist() {
			fmt.Println("There is not any wallet. please create a wallet first")
			cli.PrintUsage()
			os.Exit(0)
		}
		if *from == "" || *to == "" || *amount <= 0 {
			send.Usage()
			os.Exit(0)
		}
		c := config.NewConfig()
		w := c.GetWallet()
		sendTX(w, *from, *to, *amount, *node)

	}

	if balance.Parsed() {
		if !config.WalletExist() {
			fmt.Println("There is not any wallet. please create a wallet first")
			cli.PrintUsage()
			os.Exit(0)
		}
		c := config.NewConfig()
		w := c.GetWallet()

		for add, acc := range w.SortAccounts() {
			fmt.Printf("Account %d: %s\t Balance: %d\n", add+1, acc.Address, Balance(acc.Address, *bnode))
		}

	}

	if account.Parsed() {
		if !config.WalletExist() {
			fmt.Println("There is not any wallet. please create a wallet first")
			cli.PrintUsage()
			os.Exit(0)
		}
		c := config.NewConfig()
		w := c.GetWallet()
		w.NewAccount()
		c.SaveWallet(w)
		err := c.SaveConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	if restore.Parsed() {
		if config.WalletExist() {
			fmt.Printf("there is a wallet with name %s in %s you can't have multiple wallet in same time\n", config.DefaultName, config.WalletPath())
			os.Exit(0)
		}

		if *mnemonic == "" {
			restore.Usage()
			os.Exit(0)
		}
		if bip39.IsMnemonicValid(*mnemonic) {
			w := wallet.NewWallet(*mnemonic)
			c := config.NewConfig()
			c.SaveWallet(w)
			err := c.SaveConfig()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Wallet created successfully")
		}
		fmt.Println("Phrase Key is invalid")
	}

	if phrase.Parsed() {
		if !config.WalletExist() {
			fmt.Println("There is not any wallet. please create a wallet first")
			cli.PrintUsage()
			os.Exit(0)
		}
		c := config.NewConfig()
		w := c.GetWallet()
		fmt.Println("Never share your 12-word backup phrase and private keys with anyone.")
		fmt.Printf("\nPhrase Key:\n")
		fmt.Printf("\t\"%s\"\n", w.Mnemonic)
	}

	if adds.Parsed() {
		if !config.WalletExist() {
			fmt.Println("There is not any wallet. please create a wallet first")
			cli.PrintUsage()
			os.Exit(0)
		}
		c := config.NewConfig()
		w := c.GetWallet()
		for add, acc := range w.SortAccounts() {
			fmt.Printf("Account %d: %s\n", add+1, acc.Address)
		}
	}
}

func sendTX(w *wallet.Wallet, from, to string, amount int, node string) error {

	sender := w.Account(from)
	utxos := network.GetUTXOSet(from, node)

	msg := network.NewMsgTX("wallet", utxos, sender.PublicKeyByte, utils.Add2PKH(to), amount)

	msg.TX.TxID = msg.TX.Hash()

	// sign transaction
	sender.SignTx(&msg.TX)

	network.SendTRX(*msg)

	return nil
}
func Balance(address string, node string) int {
	utxos := network.GetUTXOSet(address, node)
	var b int
	for _, utxo := range utxos {
		b += utxo.Txout.Value
	}

	return b
}
