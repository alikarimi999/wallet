package network

import (
	"log"

	"github.com/alikarimi999/wallet/wallet"
)

type sendUTXOSET struct {
	Account wallet.Account `json:"account"`
	Utxos   []*wallet.UTXO `json:"utxos"`
}

type msgTX struct {
	SenderID string             `json:"sender"`
	TX       wallet.Transaction `json:"tx"`
}

func NewMsgTX(sender string, utxos []*wallet.UTXO, from wallet.Publickey, to wallet.Publickeyhash, amount int) *msgTX {
	tx, err := wallet.NewTX(utxos, from, to, amount)
	if err != nil {
		log.Fatal(err)
	}

	return &msgTX{
		SenderID: sender,
		TX:       *tx,
	}
}
