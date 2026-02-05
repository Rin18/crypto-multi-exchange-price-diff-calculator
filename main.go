package main

import (
	"crypto-price-diff-calculator/exchanges"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	//Ask for input
	var symbol string
	fmt.Println("Please enter symbol:")
	fmt.Scan(&symbol)
	symbol = strings.ToUpper(symbol)

	// var wg sync.WaitGroup
	// wg.Add(2)

	binanceResult := make(chan exchanges.Ticker)
	hypeResult := make(chan exchanges.Ticker)

	//Get Binance Ticker
	go func(symbol string) {
		// defer wg.Done()
		binanceTicker, err := exchanges.GetBinanceSymbol(symbol)
		if err != nil {
			fmt.Println("Get symbol failed, error:", err)
			return
		}
		binanceResult <- binanceTicker
	}(symbol)

	//Get hyperliquid Ticker
	go func(symbol string) {
		// defer wg.Done()
		hypeTicker, err := exchanges.GetHyperliquidSymbol(symbol)
		if err != nil {
			fmt.Println("Hyperliquid Error, error:", err)
			return
		}
		hypeResult <- hypeTicker
	}(symbol)
	// wg.Wait()
	binanceTicker := <-binanceResult
	hypeTicker := <-hypeResult

	binancePrice, err := decimal.NewFromString(binanceTicker.Price)
	if err != nil {
		fmt.Println("Error decimal convert")
		return
	}
	hypePrice, err := decimal.NewFromString(hypeTicker.Price)
	if err != nil {
		fmt.Println("Error decimal convert")
		return
	}
	diff := binancePrice.Sub(hypePrice)
	diff = diff.Abs()
	percent := diff.DivRound(binancePrice.Add(hypePrice).Div(decimal.NewFromInt(2)), 5)
	fmt.Printf("Binance %s Price: %s\n", binanceTicker.Symbol, binanceTicker.Price)
	fmt.Printf("Hype %s Price: %s\n", hypeTicker.Symbol, hypeTicker.Price)
	fmt.Printf("Absolute diff:%s\n", diff)
	fmt.Printf("diff percent:%s", percent)
}
