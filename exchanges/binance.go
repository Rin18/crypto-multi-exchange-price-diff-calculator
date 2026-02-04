package exchanges

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetBinanceSymbol(symbol string) (Ticker, error) {
	address := "https://fapi.binance.com/fapi/v1/premiumIndex"
	params := map[string]string{"symbol": symbol}
	fullURL, err := assembleURL(address, params)
	if err != nil {
		return Ticker{}, err
	}
	ticker, err := getResponse(fullURL, "Binance")
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}

func getResponse(url string, exchange string) (Ticker, error) {
	ticker := Ticker{
		Exchange: exchange,
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return Ticker{}, err
	}
	if resp.StatusCode != 200 {
		return Ticker{}, fmt.Errorf("Server Error, status:%s", resp.Status)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return Ticker{}, err
	}
	err = json.Unmarshal(response, &ticker)
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}

func assembleURL(address string, params map[string]string) (string, error) {
	//process with parameters
	if params != nil {
		u, err := url.Parse(address)
		if err != nil {
			return "", err
		}
		q := url.Values{}
		for key, value := range params {
			q.Add(key, value)
		}
		u.RawQuery = q.Encode()
		return u.String(), nil
	} else {
		return address, nil
	}
}
