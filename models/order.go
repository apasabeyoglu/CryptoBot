package models

import (
	"context"

	"cryptobot/lib/go-binance"
)

func GetUserOrders(userID int) []*binance.Order {
	exchangeAccounts := GetUserActiveExchangeAccounts(userID)
	binanceOrderAjaxResponse := []*binance.Order{}

	if len(exchangeAccounts) > 0 {
		for _, exchangeAccount := range exchangeAccounts {
			key := exchangeAccount.Key
			secret := exchangeAccount.Secret

			if exchangeAccount.Exchange.Name == "Binance" {
				client := binance.NewClient(key, secret)

				openOrders, err := client.NewListOpenOrdersService().Do(context.Background())
				if err != nil {
					return nil
				}

				for _, order := range openOrders {
					binanceOrderAjaxResponse = append(binanceOrderAjaxResponse, order)
				}

			}

		}

		return binanceOrderAjaxResponse
	}

	//TODO: Other exchanges

	return nil
}
