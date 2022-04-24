package main

import (
	"fmt"
	"os"

	"github.com/alikarimi999/wallet/cli"
	network "github.com/alikarimi999/wallet/network/http"
	"github.com/alikarimi999/wallet/utils"
	"github.com/alikarimi999/wallet/wallet"
)

func main() {
	defer os.Exit(0)
	cmd := new(cli.Commandline)
	cmd.Run()

	// config := config.NewConfig("./here/wallet")
	// var mnemonic = "egg angle gesture jungle credit picnic globe novel aunt flower soccer path"

	// wallet := wallet.NewWallet(mnemonic, wallet.DefaultPath)
	// fmt.Printf("%x\n", wallet.Seed)
	// wallet.NewAccount()
	// wallet.NewAccount()
	// wallet.NewAccount()
	// wallet.NewAccount()
	// config.SaveWallet(wallet)
	// err := config.SaveConfig()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// w := config.GetWallet()
	// for _, a := range w.Accounts {
	// 	fmt.Println(a.Path.Account)
	// }

	// test

	// send(wallet, "1PGrKNuKLJSpui968RGVGnvHmR7vM9UynE", "1Gsojn95GN7dRuGZsn7fJni8i9TAXThQVP", 5)
	// err := send(wallet, "1PGrKNuKLJSpui968RGVGnvHmR7vM9UynE", "18yQtZqpZ424312HRHDgjAEobZFXeaewhS", 5)
	// send(wallet, "1PGrKNuKLJSpui968RGVGnvHmR7vM9UynE", "15ccLGNthwrKt4VokNop8hcC3BprUCvttJ", 2)

	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 1)

	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 1)
	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 1)

	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 10)
	// err = send(ws, "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 3)
	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", 5)

	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", 3)
	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 5)
	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", 5)

	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 8)
	// err = send(ws, "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", 3)
	// err = send(ws, "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 3)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

}

func send(w *wallet.Wallet, from, to string, amount int, node string) error {

	sender := w.Account(from)
	utxos := network.GetUTXOSet(from, node)

	msg := network.NewMsgTX("wallet", utxos, sender.PublicKeyByte, utils.Add2PKH(to), amount)

	msg.TX.TxID = msg.TX.Hash()

	// sign transaction
	sender.SignTx(&msg.TX)

	network.SendTRX(*msg)

	return nil
}

func Handle(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
