package tasks

import (
	"context"
	"cryptobot/indicators"
	"fmt"
	"strconv"

	"cryptobot/helpers"
	"cryptobot/models"

	"cryptobot/lib/go-binance"

	"github.com/astaxie/beego/toolbox"
)

func Start() {
	models.ActiveExchangeAccounts = models.GetActiveExchangeAccounts()
	models.BinanceSymbols = models.GetBinanceSymbols()
	binanceClient := binance.NewClient("", "")

	//Update binance prices every 5 seconds
	updateBinancePricesTask := toolbox.NewTask("Update Binance Prices", "0/5 * * * * *", func() error {
		binancePrices, err := binanceClient.NewListPricesService().Do(context.Background())
		if err == nil && binancePrices != nil {
			models.BinancePrices = binancePrices

			for _, p := range models.BinancePrices {
				if p.Symbol == "BTCUSDT" {
					btcUsdtBinance, err := strconv.ParseFloat(p.Price, 32)

					if err == nil && btcUsdtBinance != 0 {
						models.BtcUsdtBinance = btcUsdtBinance
					} else {
						return err
					}

					break
				}
			}
		} else {
			return err
		}
		return nil
	})

	//Insert user balances every 5 minutes
	insertUserBalancesTask := toolbox.NewTask("Insert User Balances", "0 */5 * * * *", func() error {
		for _, exchangeAccount := range models.ActiveExchangeAccounts {
			userID := exchangeAccount.User.ID
			exchangeAccountID := exchangeAccount.ID
			key := exchangeAccount.Key
			secret := exchangeAccount.Secret

			if exchangeAccount.Exchange.Name == "Binance" {
				client := binance.NewClient(key, secret)
				account, err := client.NewGetAccountService().Do(context.Background())
				totalBtc := 0.0

				if err != nil {
					return err
				}

				for _, binanceBalance := range account.Balances {
					free, _ := strconv.ParseFloat(binanceBalance.Free, 32)
					locked, _ := strconv.ParseFloat(binanceBalance.Locked, 32)

					if free > 0 || locked > 0 {
						if binanceBalance.Asset != "USDT" {
							for _, p := range models.BinancePrices {
								if p.Symbol == binanceBalance.Asset+"BTC" {
									price, _ := strconv.ParseFloat(p.Price, 32)
									totalBtc += (free + locked) * price
									break
								}
							}
						} else if binanceBalance.Asset == "USDT" {
							totalBtc += (free + locked) / models.BtcUsdtBinance
						}
					}
				}

				totalBtc = helpers.SetDecimals(totalBtc, 8)
				totalUsdt := helpers.SetDecimals(totalBtc*models.BtcUsdtBinance, 2)

				err = models.InsertUserBalance(userID, exchangeAccountID, totalBtc, totalUsdt)

				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	//Check prices and auto trade every 5 minutes
	autoTradeTask := toolbox.NewTask("Auto Trade", "0 */5 * * * *", func() error {
		interval := "1d" //1 Day
		std := 1.0       //for bollinger
		for _, symbol := range models.BinanceSymbols {
			client := binance.NewClient("", "")
			klines, err := client.NewKlinesService().Symbol(symbol).Interval(interval).Do(context.Background())

			if err != nil {
				break
			}

			var prices []float64
			var currentPrice float64
			var days = 30

			for _, kline := range klines {
				lastPrice, _ := strconv.ParseFloat(kline.Close, 32)
				prices = append(prices, lastPrice)
			}

			for _, p := range models.BinancePrices {
				if p.Symbol == symbol {
					currentPrice, _ = strconv.ParseFloat(p.Price, 32)
					break
				}
			}

			if currentPrice != 0 && len(prices) > days {
				fmt.Println(symbol, currentPrice)
				indicators.BollingerDecision(prices, days, std, currentPrice)
			}

			//TODO
		}

		return nil
	})

	toolbox.AddTask("Update Binance Prices", updateBinancePricesTask)
	toolbox.AddTask("Insert User Balances", insertUserBalancesTask)
	toolbox.AddTask("Auto Trade", autoTradeTask)

	err := updateBinancePricesTask.Run()
	if err != nil {
		fmt.Println(err)
	}

	err = insertUserBalancesTask.Run()
	if err != nil {
		fmt.Println(err)
	}

	err = autoTradeTask.Run()
	if err != nil {
		fmt.Println(err)
	}

	toolbox.StartTask()
	defer toolbox.StopTask()
}
