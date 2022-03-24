package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"os"
)

const (
	file string = `./wallet.data`
)

type Wallets struct {
	Wallets map[Account]*Wallet
}

func NewWallets() *Wallets {
	ws := &Wallets{make(map[Account]*Wallet)}

	return ws
}

func (ws *Wallets) AddTowallets(w ...*Wallet) error {

	for _, a := range w {
		address := a.Address()
		ws.Wallets[Account(address)] = a
	}
	return nil

}

func SaveWallets(ws *Wallets) error {

	wf, err := os.Create(file)
	Handle(err)
	defer wf.Close()

	wf.Write(ws.Serialize())

	return nil
}

func (ws *Wallets) Serialize() []byte {

	var buff bytes.Buffer
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(ws)
	Handle(err)

	return buff.Bytes()

}

func LoadFile(f string) (*Wallets, error) {
	fc, err := os.ReadFile(f)
	Handle(err)
	ws, _ := Deserialize(fc)
	return ws, nil

}

func Deserialize(b []byte) (*Wallets, error) {
	var ws Wallets
	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(&ws)
	Handle(err)
	return &ws, nil
}
