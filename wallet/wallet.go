package wallet

import (
	"crypto/ecdsa"
	"log"

	"github.com/alikarimi999/wallet/utils"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	MasterKey string
	Mnemonic  string
	Seed      []byte
	Accounts  map[string]*Account
}

type Account struct {
	ExtendedKey *hdkeychain.ExtendedKey // extended private key

	BtcecPriv *btcec.PrivateKey
	BtcecPub  *btcec.PublicKey

	PriavateKey *ecdsa.PrivateKey
	PublicKey   *ecdsa.PublicKey

	PublicKeyByte []byte // serialized compressed Public Key
	Address       string

	Path Path
}

func NewWallet(mnemonic string, p Path) *Wallet {
	seed := bip39.NewSeed(mnemonic, "")

	w := &Wallet{
		Mnemonic: mnemonic,
		Seed:     seed,
		Accounts: make(map[string]*Account),
	}

	masterKey, err := hdkeychain.NewMaster(seed, &utils.Params)
	if err != nil {
		log.Fatal(err)
	}

	w.MasterKey = masterKey.String()
	// creat first account
	w.NewAccount()

	return w
}

func (w *Wallet) Account(a string) *Account {
	return w.Accounts[a]
}

func (w *Wallet) NewAccount() {
	a := &Account{}

	key, err := hdkeychain.NewMaster(w.Seed, &utils.Params)
	if err != nil {
		log.Fatal(err)
	}
	p := DefaultPath
	path := []uint32{p.Purpose, p.CoinType, p.Account, p.Change, uint32(len(w.Accounts))}
	p.AddressIndex = path[4]
	for _, i := range path {
		key, err = key.Child(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	a.ExtendedKey = key
	a.Path = p

	priv, err := a.ExtendedKey.ECPrivKey()
	if err != nil {
		log.Fatal(err)
	}
	a.PriavateKey = priv.ToECDSA()
	a.PublicKey = &a.PriavateKey.PublicKey
	a.BtcecPriv = priv
	a.BtcecPub = priv.PubKey()
	address, err := a.ExtendedKey.Address(&utils.Params)
	if err != nil {
		log.Fatal(err)
	}
	a.PublicKeyByte = a.BtcecPub.SerializeCompressed()
	a.Address = address.EncodeAddress()

	w.Accounts[a.Address] = a
}

func (a *Account) SignTx(tx *Transaction) {

	for _, in := range tx.TxInputs {
		in.Signature = nil
		in.PublicKey = nil
	}

	sig, err := a.BtcecPriv.Sign(tx.TxID)
	if err != nil {
		log.Fatal(err)
	}
	sigSer := sig.Serialize()

	for _, in := range tx.TxInputs {
		in.Signature = sigSer
		in.PublicKey = a.PublicKeyByte
	}

}
