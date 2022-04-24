package wallet

import "time"

type Publickey []byte // btcec serialized compressed publickey
type Publickeyhash []byte

type Chainid int

type Transaction struct {
	Timestamp time.Duration
	TxID      []byte
	TxInputs  []*TxIn
	TxOutputs []*TxOut
}

type TxOut struct {
	PublicKeyHash []byte
	Value         int
}

type TxIn struct {
	OutPoint  []byte
	Vout      uint
	Value     int
	PublicKey []byte
	Signature []byte
}

type UTXO struct {
	Txid  []byte
	Index uint
	Txout *TxOut
}
