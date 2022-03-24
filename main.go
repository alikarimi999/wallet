package main

import (
	"fmt"
	"log"

	network "github.com/alikarimi999/wallet/network/http"
	"github.com/alikarimi999/wallet/wallet"
)

func main() {

	ws, err := wallet.LoadFile("./wallet.data")
	if err != nil {
		log.Panic(err)
	}

	// ali to john
	send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 10)
	// // // john to ali
	err = send(ws, "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", 5)
	// // // //
	err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 2)

	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", 1)
	// err = send(ws, "19xX1RHa44YxjLrz2jQbPrdcen2h6SztVb", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 20)
	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 12)
	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", 1)

	// err = send(ws, "1tHozHekDBPngFbGfS3msUKDgNezECFe8", "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", 8)
	// err = send(ws, "16mqptJgurSbhqdtF1GDLD8CkDsmRt69ge", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 9)
	// err = send(ws, "1F9syUQpU3EBnna1deuwK9Vyr38NX87t65", "1tHozHekDBPngFbGfS3msUKDgNezECFe8", 1)

	if err != nil {
		log.Fatalln(err)
	}

}

func send(ws *wallet.Wallets, from, to wallet.Account, amount int) error {

	w := ws.Wallets[from]
	pk := w.PubKey
	utxos := network.GetUTXOSet(from)

	trx, err := wallet.NewTX(utxos, pk, wallet.Addr2PKH(to), amount)
	if err != nil {
		return err
	}
	trx, err = w.SignTX(trx)
	if err != nil {
		return err
	}
	network.SendTRX(*trx)

	return nil
}

func Handle(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
