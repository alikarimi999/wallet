package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Wallet struct {
	PrivKey ecdsa.PrivateKey
	PubKey  []byte
}

func MakeWallet() *Wallet {
	priv, pub := GeneratPairkey()
	return &Wallet{priv, pub}
}

func GeneratPairkey() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	Handle(err)
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub

}

func (w *Wallet) SignTX(tx *Transaction) (*Transaction, error) {

	for _, in := range tx.TxInputs {
		in.Signature = nil
		in.PublicKey = nil
	}
	r, s, err := ecdsa.Sign(rand.Reader, &w.PrivKey, tx.Serialize())
	Handle(err)
	signature := append(r.Bytes(), s.Bytes()...)

	for _, in := range tx.TxInputs {
		in.Signature = signature
		in.PublicKey = w.PubKey
	}

	return tx, nil
}
