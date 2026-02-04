package exchanges

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetHyperliquidSymbol(symbol string) (Ticker, error) {
	rawSymbol, err := getHyperliquidResponse()
	if err != nil {
		return Ticker{}, err
	}
	hypeTicker, err := convertHypeResponse(symbol, rawSymbol)
	if err != nil {
		return Ticker{}, err
	}
	return hypeTicker, nil
}

type hyperliquidPOST struct {
	Type string `json:"type"`
}

func getHyperliquidResponse() ([]byte, error) {
	// var hyperliquidTicker Ticker
	url := "https://api.hyperliquid.xyz/info"
	contentType := "application/json"
	metaPOST := hyperliquidPOST{
		Type: "metaAndAssetCtxs",
	}
	metaByte, err := json.Marshal(&metaPOST)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(metaByte))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Server Error, Status:%s", resp.Status)
	}
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func convertHypeResponse(symbol string, response []byte) (Ticker, error) {
	var rawResponse []json.RawMessage
	err := json.Unmarshal(response, &rawResponse)
	if err != nil {
		return Ticker{}, err
	}
	var part1 struct {
		Universe []struct {
			Name string `json:"name"`
		} `json:"universe"`
	}
	var part2 []struct {
		MarkPx string `json:"markPx"`
	}
	err = json.Unmarshal(rawResponse[0], &part1)
	if err != nil {
		return Ticker{}, err
	}
	err = json.Unmarshal(rawResponse[1], &part2)
	if err != nil {
		return Ticker{}, err
	}
	hypeSymbol := convertSymbol(symbol)
	var index int = -1
	for i, j := range part1.Universe {
		if j.Name == hypeSymbol {
			index = i
			break
		}
	}
	if index == -1 {
		return Ticker{}, fmt.Errorf("Hyperliquid Symbol not found")
	}
	hyperliquidTicker := Ticker{
		Symbol:    symbol,
		Price:     part2[index].MarkPx,
		Timestamp: 0,
	}
	return hyperliquidTicker, nil
}

func convertSymbol(symbol string) string {
	replacer := strings.NewReplacer("USDC", "", "USDT", "")
	pureSymbol := replacer.Replace(symbol)
	return pureSymbol
}
