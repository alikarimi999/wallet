package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alikarimi999/wallet/wallet"
)

func GetUTXOSet(user wallet.Account) []*wallet.UTXO {
	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/getutxo?account=%s", user))

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	su := &sendUTXOSET{}

	json.Unmarshal(body, su)

	return su.Utxos

}

func SendTRX(tx wallet.Transaction) {

	b, err := json.Marshal(tx)
	if err != nil {
		log.Panic(err)
	}
	resp, err := http.Post("http://localhost:5000/sendtrx", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println()
	io.Copy(os.Stdout, resp.Body)
}
