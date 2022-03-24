package wallet

import "time"

type publickey []byte
type publickeyhash []byte

type Chainid int

// Account address in string
type Account string

type Transaction struct {
	Timestamp time.Time
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
