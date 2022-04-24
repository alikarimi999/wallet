package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alikarimi999/wallet/utils"
	"github.com/alikarimi999/wallet/wallet"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
)

const (
	DefaultName string = "wallet"
)

type config struct {
	configName string
	configPath string
	configType string

	viper *viper.Viper
}

func NewConfig() *config {
	c := &config{
		configName: DefaultName,
		configPath: WalletPath(),
		configType: "json",
		viper:      viper.New(),
	}

	c.Creat()
	return c
}

func WalletPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, ".shitcoin_wallet")
}

func WalletExist() bool {
	wallet_path := WalletPath()
	path := filepath.Join(wallet_path, DefaultName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		return false
	}
	return true
}

func (c *config) Creat() {
	c.viper.SetConfigName(c.configName)
	c.viper.AddConfigPath(c.configPath)
	c.viper.SetConfigType(c.configType)

	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			os.MkdirAll(c.configPath, os.ModePerm)
			os.Create(filepath.Join(c.configPath, c.configName))
		}
	}

}

func (c *config) SaveWallet(w *wallet.Wallet) {
	addresses := []string{}
	c.viper.Set("wallet.masterkey", w.MasterKey)
	c.viper.Set("wallet.mnemonic", w.Mnemonic)
	for add, acc := range w.Accounts {
		c.SaveAccount(acc)
		addresses = append(addresses, add)
	}
	c.viper.Set("wallet.addresses", addresses)

}

func (c *config) SaveAccount(a *wallet.Account) {
	c.viper.Set(fmt.Sprintf("wallet.accounts.%s.xpriv", a.Address), a.ExtendedKey.String())
	c.viper.Set(fmt.Sprintf("wallet.accounts.%s.path", a.Address), a.Path.String())

}

func (c *config) GetWallet() *wallet.Wallet {
	w := &wallet.Wallet{
		Accounts: make(map[string]*wallet.Account),
	}
	w.MasterKey = c.viper.GetString("wallet.masterkey")
	w.Mnemonic = c.viper.GetString("wallet.mnemonic")
	w.Seed = bip39.NewSeed(w.Mnemonic, "")
	addresses := c.viper.GetStringSlice("wallet.addresses")
	for _, add := range addresses {
		acc := c.GetAccount(add)
		w.Accounts[acc.Address] = acc
	}

	return w

}

func (c *config) GetAccount(address string) *wallet.Account {
	a := &wallet.Account{}
	key := c.viper.GetString(fmt.Sprintf("wallet.accounts.%s.xpriv", address))
	exkey, err := hdkeychain.NewKeyFromString(key)
	if err != nil {
		log.Fatal(err)
	}
	a.ExtendedKey = exkey
	a.BtcecPriv, _ = exkey.ECPrivKey()
	a.BtcecPub, _ = exkey.ECPubKey()
	a.PriavateKey = a.BtcecPriv.ToECDSA()
	a.PublicKey = a.BtcecPub.ToECDSA()
	a.PublicKeyByte = a.BtcecPub.SerializeCompressed()

	a.Address = utils.PK2Add(a.PublicKeyByte, false)

	path := c.viper.GetString(fmt.Sprintf("wallet.accounts.%s.path", address))
	a.Path = wallet.String2Path(path)

	return a
}

func (c *config) SaveConfig() error {
	return c.viper.WriteConfig()
}
