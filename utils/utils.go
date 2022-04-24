package utils

import (
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

var (
	PubKeyHashAddrID = 0x00                            // starts with 1
	HDPrivateKeyID   = [4]byte{0x04, 0x88, 0xad, 0xe4} // starts with xprv
	HDPublicKeyID    = [4]byte{0x04, 0x88, 0xb2, 0x1e} // starts with xpub

	Params = chaincfg.Params{
		PubKeyHashAddrID: 0x00,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub
		// // BIP44 coin type used in the hierarchical deterministic path for
		// // address generation.
		HDCoinType: 0,
	}
)

// convert Serialized compressed publicKey or publicKeyHash to address
func PK2Add(pk []byte, isHash bool) string {
	if !isHash {
		add, err := btcutil.NewAddressPubKey(pk, &Params)
		if err != nil {
			log.Fatal(err)
		}
		return add.EncodeAddress()
	}

	add, err := btcutil.NewAddressPubKeyHash(pk, &Params)
	if err != nil {
		log.Fatal(err)
	}
	return add.EncodeAddress()

}

func Add2PKH(address string) []byte {
	pkh, _, err := base58.CheckDecode(address)
	if err != nil {
		log.Fatal(err)
	}
	return pkh
}
