package wallet

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	addressVersion = byte(0x00)
	ChecksumLenght = 4
)

func (w *Wallet) Address() []byte {
	pub := w.PubKey

	publicKeyHash := Hash160(pub)
	versionedHahs := AddVersion(publicKeyHash, addressVersion)
	check := Checksum(versionedHahs)

	cHash := append(versionedHahs, check...)

	address := base58.Encode(cHash)

	return []byte(address)
}

func Checksum(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])

	return h2[:ChecksumLenght]
}

// this function add addressVersion
func AddVersion(b []byte, v byte) []byte {

	return append([]byte{v}, b...)
}

// this is obvoious
func Hash160(pub []byte) []byte {
	hash := sha256.Sum256(pub)

	hasher := ripemd160.New()
	_, err := hasher.Write(hash[:])
	Handle(err)

	pkh := hasher.Sum(nil)

	return pkh
}

func PK2Add(pk []byte) Account {
	publicKeyHash := Hash160(pk)
	versionedHahs := AddVersion(publicKeyHash, byte(0x00))
	check := Checksum(versionedHahs)

	cHash := append(versionedHahs, check...)

	address := base58.Encode(cHash)

	return Account(address)

}

// this function help us to validate an account addres
func ValidateAddress(address string) bool {
	publickeyhash, err := base58.Decode(address)
	Handle(err)
	actualChecksum := publickeyhash[len(publickeyhash)-ChecksumLenght:]
	version := publickeyhash[0]
	publickeyhash = publickeyhash[1 : len(publickeyhash)-ChecksumLenght]
	targetChecksum := Checksum(append([]byte{version}, publickeyhash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}

// this is a simple error handler
func Handle(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Addr2PKH(address Account) []byte {
	publickeyhash, err := base58.Decode(string(address))
	Handle(err)
	publickeyhash = publickeyhash[1 : len(publickeyhash)-ChecksumLenght]
	return publickeyhash
}
