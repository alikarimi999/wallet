package wallet

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alikarimi999/wallet/utils"
	"github.com/btcsuite/btcutil"
)

func NewTX(utxos []*UTXO, from Publickey, to Publickeyhash, amount int) (*Transaction, error) {

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
		fmt.Printf("%s just have %d money\n", utils.PK2Add(from, false), acc)
		return nil, errors.New("you dont have enough money")
	}

	if acc > amount {
		tx.TxOutputs = []*TxOut{
			{PublicKeyHash: to, Value: amount},
			{PublicKeyHash: btcutil.Hash160(from), Value: acc - amount},
		}
	} else {
		tx.TxOutputs = []*TxOut{
			{PublicKeyHash: to, Value: amount},
		}
	}

	tx.Timestamp = time.Duration(time.Now().UTC().UnixNano())

	return tx, nil
}

func (tx *Transaction) Hash() []byte {
	for _, in := range tx.TxInputs {
		in.Signature = nil
		in.PublicKey = nil
	}
	data := tx.Serialize()

	hash := sha256.Sum256(data)

	return hash[:]
}

func SerializeIns(ins []*TxIn) []byte {
	b := []byte{}
	for _, in := range ins {
		b = append(b, in.serialize()...)
	}
	return b
}

func Serializeouts(outs []*TxOut) []byte {
	b := []byte{}
	for _, in := range outs {
		b = append(b, in.serialize()...)
	}
	return b
}

func (tx *Transaction) Serialize() []byte {

	b := bytes.Join(
		[][]byte{
			Int2Hex(int64(tx.Timestamp)),
			tx.TxID,
			SerializeIns(tx.TxInputs),
			Serializeouts(tx.TxOutputs),
		},
		nil,
	)

	return b

}
func (in *TxIn) serialize() []byte {
	b := bytes.Join(
		[][]byte{
			in.OutPoint,
			Int2Hex(int64(in.Vout)),
			Int2Hex(int64(in.Value)),
			in.PublicKey,
			in.Signature,
		}, nil,
	)
	return b
}

func (out *TxOut) serialize() []byte {
	b := bytes.Join(
		[][]byte{
			out.PublicKeyHash,
			Int2Hex(int64(out.Value)),
		}, nil,
	)
	return b
}

func Int2Hex(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		log.Fatalln(err)
	}

	return buff.Bytes()
}
