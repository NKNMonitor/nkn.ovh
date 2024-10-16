package nknovh_engine

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Prices struct {
	Nkn struct {
		Usd float64 `json:"usd"`
	} `json:"nkn"`
}

func (o *NKNOVH) getPrices() error {
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	var cg_url string = "https://api.coingecko.com/api/v3/simple/price?ids=nkn&vs_currencies=usd"
	resp, err := client.Get(cg_url)
	if err != nil {
		o.log.Error("An error occured while getting the prices: " + err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.log.Error("An error occured while getting the prices: " + err.Error())
		return err
	}
	var prices = new(Prices)
	err = json.Unmarshal(body, prices)
	if err != nil {
		o.log.Error("An error occured while getting the prices: " + err.Error())
		return err
	}

	var id int
	row := o.sql.stmt["main"]["getPriceByName"].QueryRow("usd")
	err = row.Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		if _, err1 := o.sql.stmt["main"]["insertPrice"].Exec("usd", prices.Nkn.Usd); err1 != nil {
			o.log.Error("Stmt insertPrice has returned an error: (" + err1.Error() + ")")
			return err1
		}
		break
	case err != nil:
		o.log.Error("Can't execute row.Scan(): " + err.Error())
		return err

	default:
		if _, err1 := o.sql.stmt["main"]["updatePriceById"].Exec(prices.Nkn.Usd, id); err1 != nil {
			o.log.Error("Stmt updatePriceById has returned an error: (" + err1.Error() + ")")
			return err1
		}
	}
	return nil
}
