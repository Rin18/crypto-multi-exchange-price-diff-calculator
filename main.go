package main

import (
	"crypto-price-diff-calculator/exchanges"
	"fmt"
)

func main() {
	ticker, err := exchanges.GetBinanceSymbol("BTCUSDT")
	if err != nil {
		fmt.Println("Get symbol failed, error:", err)
		return
	}
	fmt.Println(ticker)
	fmt.Println("Hyperliquid:")
	hpreturn, err := exchanges.GetHyperliquidSymbol("BTCUSDT")
	if err != nil {
		fmt.Println("Hyperliquid Error, error:", err)
		return
	}
	fmt.Println(hpreturn)
}
