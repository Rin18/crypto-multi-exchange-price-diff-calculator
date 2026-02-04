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
	ticker, err := getBinanceResponse(fullURL, "Binance")
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}

func getBinanceResponse(url string, exchange string) (Ticker, error) {
	binanceTicker := Ticker{}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return Ticker{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Ticker{}, fmt.Errorf("Server Error, status:%s", resp.Status)
	}
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return Ticker{}, err
	}
	err = json.Unmarshal(response, &binanceTicker)
	if err != nil {
		return Ticker{}, err
	}
	return binanceTicker, nil
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
