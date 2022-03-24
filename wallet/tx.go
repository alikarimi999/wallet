package wallet

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
)

func NewTX(utxos []*UTXO, from publickey, to publickeyhash, amount int) (*Transaction, error) {

	tx := new(Transaction)

	// tokens valuation
	acc := 0

	// creat transaction inputs from utxos
	for _, utxo := range utxos {
		if acc < amount {
			txin := &TxIn{
				OutPoint:  utxo.Txid,
				Vout:      utxo.Index,
				Value:     utxo.Txout.Value,
				PublicKey: nil,
				Signature: nil,
			}

			tx.TxInputs = append(tx.TxInputs, txin)
			acc += utxo.Txout.Value
		}
	}
	if acc < amount {
		fmt.Printf("%s just have %d money\n", PK2Add(from), acc)
		return nil, errors.New("you dont have enough money")
	}

	if acc > amount {
		tx.TxOutputs = []*TxOut{
			{PublicKeyHash: to, Value: amount},
			{PublicKeyHash: Hash160(from), Value: acc - amount},
		}
	} else {
		tx.TxOutputs = []*TxOut{
			{PublicKeyHash: to, Value: amount},
		}
	}
	tx.SetHash()
	return tx, nil
}

func (tx *Transaction) SetHash() {
	for _, in := range tx.TxInputs {
		in.Signature = nil
	}
	data := tx.Serialize()

	hash := sha256.Sum256(data)

	tx.TxID = hash[:]
}

func (tx *Transaction) Serialize() []byte {
	buff := new(bytes.Buffer)

	encoder := gob.NewEncoder(buff)
	err := encoder.Encode(tx)
	Handle(err)

	return buff.Bytes()
}
