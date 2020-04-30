package models

import (
	"context"
	"cryptobot/helpers"
	"fmt"
	"strconv"
	"time"

	"cryptobot/lib/go-binance"

	"github.com/astaxie/beego/orm"
)

type Balance struct {
	ID              int              `orm:"column(balance_id);pk;auto"`
	User            *User            `orm:"rel(fk)"`
	ExchangeAccount *ExchangeAccount `orm:"rel(fk)"`
	TotalBtc        float64
	TotalUsdt       float64
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)"`
}

type BalanceAjaxResponse struct {
	Currency       string
	Free           float64
	Total          float64
	UnitPrice      float64
	BtcEquivalent  float64
	UsdtEquivalent float64
	TradeURL       string
}

type BalancesAjaxResponse struct {
	Exchange      string
	Name          string
	LogoImagePath string
	BtcUsdt       float64
	TotalBtc      float64
	TotalUsdt     float64
	Balances      []BalanceAjaxResponse
}

type GraphicDataAjaxResponse struct {
	Date string  `orm:"column(created_at)"`
	Btc  float64 `orm:"column(total_btc)"`
	Usdt float64 `orm:"column(total_usdt)"`
}

func (b *Balance) TableName() string {
	return "balances"
}

func init() {
	orm.RegisterModel(new(Balance))
}

func GetUserBalances(userID int) []BalancesAjaxResponse {
	exchangeAccounts := GetUserActiveExchangeAccounts(userID)
	allBalancesAjaxResponse := []BalancesAjaxResponse{}

	if len(exchangeAccounts) > 0 {
		for _, exchangeAccount := range exchangeAccounts {
			balancesAjaxResponse := BalancesAjaxResponse{}
			balances := []BalanceAjaxResponse{}
			key := exchangeAccount.Key
			secret := exchangeAccount.Secret

			if exchangeAccount.Exchange.Name == "Binance" {
				client := binance.NewClient(key, secret)
				account, err := client.NewGetAccountService().Do(context.Background())

				totalBtc := 0.0

				if err == nil {
					for _, binanceBalance := range account.Balances {
						balance := BalanceAjaxResponse{}
						free, _ := strconv.ParseFloat(binanceBalance.Free, 32)
						locked, _ := strconv.ParseFloat(binanceBalance.Locked, 32)

						if free > 0 || locked > 0 {

							balance.Currency = binanceBalance.Asset
							balance.Free = helpers.SetDecimals(free, 8)
							balance.Total = helpers.SetDecimals(free+locked, 8)

							if binanceBalance.Asset != "BTC" && binanceBalance.Asset != "USDT" {
								for _, p := range BinancePrices {
									if p.Symbol == binanceBalance.Asset+"BTC" {
										price, _ := strconv.ParseFloat(p.Price, 32)
										btcEquivalent := helpers.SetDecimals((free+locked)*price, 8)
										totalBtc += btcEquivalent

										balance.UnitPrice = helpers.SetDecimals(price, 8)
										balance.BtcEquivalent = btcEquivalent
										balance.UsdtEquivalent = helpers.SetDecimals(btcEquivalent*BtcUsdtBinance, 2)
										balance.TradeURL = "https://www.binance.com/trade.html?symbol=" + binanceBalance.Asset + "_BTC"
										break
									}
								}
							} else if binanceBalance.Asset == "BTC" {
								btcEquivalent := helpers.SetDecimals(free+locked, 8)
								balance.UnitPrice = 1
								balance.BtcEquivalent = btcEquivalent
								balance.UsdtEquivalent = helpers.SetDecimals(btcEquivalent*BtcUsdtBinance, 2)
								balance.TradeURL = "https://www.binance.com/trade.html?symbol=BTC_USDT"
								totalBtc += btcEquivalent
							} else if binanceBalance.Asset == "USDT" {
								btcEquivalent := helpers.SetDecimals((free+locked)/BtcUsdtBinance, 8)
								balance.UnitPrice = helpers.SetDecimals(1/BtcUsdtBinance, 8)
								balance.BtcEquivalent = btcEquivalent
								balance.UsdtEquivalent = helpers.SetDecimals(free+locked, 2)
								balance.TradeURL = "https://www.binance.com/trade.html?symbol=BTC_USDT"
								totalBtc += btcEquivalent
							}

							balances = append(balances, balance)

						}
					}

					balancesAjaxResponse.Exchange = "Binance"

					if len(exchangeAccount.Name) > 0 {
						balancesAjaxResponse.Name = exchangeAccount.Name
					} else {
						balancesAjaxResponse.Name = exchangeAccount.Key[0:12] + "..."
					}

					balancesAjaxResponse.LogoImagePath = "/static/images/exchanges/" + exchangeAccount.Exchange.LogoImageName
					balancesAjaxResponse.BtcUsdt = BtcUsdtBinance
					balancesAjaxResponse.TotalBtc = totalBtc
					balancesAjaxResponse.TotalUsdt = totalBtc * BtcUsdtBinance
					balancesAjaxResponse.Balances = balances

					allBalancesAjaxResponse = append(allBalancesAjaxResponse, balancesAjaxResponse)
				}
			}
		}
	}

	return allBalancesAjaxResponse
}

func InsertUserBalance(userID, exchangeAccountID int, totalBtc, totalUsdt float64) error {
	o := orm.NewOrm()
	user := User{ID: userID}
	exchangeAccount := ExchangeAccount{ID: exchangeAccountID}
	balance := Balance{User: &user, ExchangeAccount: &exchangeAccount, TotalBtc: totalBtc, TotalUsdt: totalUsdt}
	_, err := o.Insert(&balance)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetUserBalanceGraphicData(userID, exchangeAccountID int) []GraphicDataAjaxResponse {
	o := orm.NewOrm()
	graphicDataAjaxResponse := []GraphicDataAjaxResponse{}

	if o.Driver().Type() == orm.DRSqlite {
		o.Raw(`SELECT total_btc,
				total_usdt, 
				strftime('%Y-%m-%d %H:%M:%S', datetime(created_at, '+3 Hour')) as created_at
				FROM balances 
				WHERE user_id = ? AND exchange_account_id = ?`, userID, exchangeAccountID).QueryRows(&graphicDataAjaxResponse)
	} else {
		fmt.Println("strftime not supported")
	}

	return graphicDataAjaxResponse
}

func GetBinanceSymbols() []string {
	client := binance.NewClient("", "")
	var symbols []string

	exchangeInfo, _ := client.NewExchangeInfoService().Do(context.Background())

	for _, info := range exchangeInfo.Symbols {
		symbols = append(symbols, info.Symbol)
	}

	return symbols
}
