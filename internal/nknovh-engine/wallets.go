package nknovh_engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	nkn "github.com/nknorg/nkn-sdk-go"
	"go.uber.org/zap"
)

type Nknsdk struct {
	Wallet *nkn.Wallet
}

func (o *NKNOVH) walletCreate() error {
	if _, err := os.Stat(fmt.Sprintf("%s/wallet.json", o.WalletPath)); err == nil {
		if _, err := os.Stat(fmt.Sprintf("%s/wallet.pswd", o.WalletPath)); err == nil {
			return nil
		}
	}

	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}

	wpswd := RandBytes(32)

	wallet, err := nkn.NewWallet(account, &nkn.WalletConfig{Password: string(wpswd)})
	if err != nil {
		return err
	}
	walletJSON, err := wallet.ToJSON()
	if err != nil {
		return err
	}

	wFile, err := os.OpenFile(fmt.Sprintf("%s/wallet.json", o.WalletPath), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer wFile.Close()
	pFile, err := os.OpenFile(fmt.Sprintf("%s/wallet.pswd", o.WalletPath), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer pFile.Close()

	if _, err := wFile.WriteString(walletJSON); err != nil {
		return err
	}
	if _, err := pFile.Write(wpswd); err != nil {
		return err
	}

	return nil
}

func (o *NKNOVH) nknConnect() error {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.json", o.WalletPath))
	if err != nil {
		return err
	}
	wpswd, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.pswd", o.WalletPath))
	if err != nil {
		return err
	}
	wallet, err := nkn.WalletFromJSON(string(data), &nkn.WalletConfig{Password: string(wpswd)})
	if err != nil {
		return err
	}
	if err := nkn.VerifyWalletAddress(wallet.Address()); err != nil {
		return err
	}
	if err := wallet.VerifyPassword(string(wpswd)); err != nil {
		return err
	}
	t := new(Nknsdk)
	t.Wallet = wallet
	o.Nknsdk = t
	return nil
}

func (o *NKNOVH) walletPoll() error {
	go o.getPrices()
	if err := o.fetchBalances(); err != nil {
		return err
	}
	return nil
}

func (o *NKNOVH) fetchBalances() error {

	var (
		id         uint
		nkn_wallet string
		db_balance float64
	)
	var wallet *nkn.Wallet = o.Nknsdk.Wallet

	//fetch wallets from the database
	rows, err := o.sql.stmt["main"]["selectWallets"].Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &nkn_wallet, &db_balance); err != nil {
			return err
		}
		balance, err := wallet.BalanceByAddress(nkn_wallet)
		if err != nil {
			o.log.Error("Query balance fail: ", zap.Error(err))
			continue
		}
		float_balance, err := strconv.ParseFloat(balance.String(), 64)
		if err != nil {
			o.log.Error("Cannot transform string to float64: ", zap.Error(err))
			continue
		}
		if float_balance != db_balance {
			if _, err1 := o.sql.stmt["main"]["updateWalletBalanceById"].Exec(&float_balance, &id); err1 != nil {
				o.log.Error("Stmt updateWalletBalanceById has returned an error: (", zap.Error(err))
				continue
			}
		}
	}
	return nil
}
