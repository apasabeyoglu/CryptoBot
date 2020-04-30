package models

import (
	"cryptobot/lib/go-binance"
)

var ActiveExchangeAccounts []ExchangeAccount
var BinancePrices []*binance.SymbolPrice
var BinanceSymbols []string
var BtcUsdtBinance float64
