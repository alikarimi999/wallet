package network

import (
	"time"

	"github.com/alikarimi999/wallet/wallet"
)

type sendUTXOSET struct {
	Account wallet.Account `json:"account"`
	Utxos   []*wallet.UTXO `json:"utxos"`
}

type msgTRX struct {
	Timestamp time.Time
	TXID      []byte
	TxInputs  []wallet.TxOut
	TxOutputs []wallet.TxIn
}
