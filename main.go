package main

import (
	"crypto-price-diff-calculator/exchanges"
	"fmt"
	"strings"
)

func main() {
	//Ask for input
	var symbol string
	fmt.Println("Please enter symbol:")
	fmt.Scan(&symbol)
	symbol = strings.ToUpper(symbol)
	//Get Binance Ticker
	fmt.Println("Binance:")
	binanceTicker, err := exchanges.GetBinanceSymbol(symbol)
	if err != nil {
		fmt.Println("Get symbol failed, error:", err)
		return
	}
	fmt.Println(binanceTicker)
	//Get hyperliquid Ticker
	fmt.Println("Hyperliquid:")
	hypeTicker, err := exchanges.GetHyperliquidSymbol(symbol)
	if err != nil {
		fmt.Println("Hyperliquid Error, error:", err)
		return
	}
	fmt.Println(hypeTicker)
}
