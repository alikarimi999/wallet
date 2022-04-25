package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alikarimi999/wallet/utils"
	"github.com/alikarimi999/wallet/wallet"
)

func GetUTXOSet(address string, node string) []*wallet.UTXO {
	resp, err := http.Get(utils.JoinURL(node, fmt.Sprintf("getutxo?account=%s", address)))

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

func SendTRX(msg msgTX, node string) {

	b, err := json.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	resp, err := http.Post(utils.JoinURL(node, "sendtrx"), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println()
	io.Copy(os.Stdout, resp.Body)
}
