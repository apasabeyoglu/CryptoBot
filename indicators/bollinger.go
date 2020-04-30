package indicators

import (
	"log"
	"math"
)

type float []float64

var openOrder = false

func sum(x float) float64 {

	var result float64

	for _, i := range x {
		result += i
	}

	return result
}

// Used in Bollinger
func average(x float) float64 {
	return sum(x) / float64(len(x))
}

// Standard Deviation (used in Bollinger)
func std(x float) float64 {

	var result float
	mean := average(x)

	for i := 0; i < len(x); i++ {
		result = append(result, math.Pow(x[i]-mean, 2))
	}

	return math.Sqrt(sum(result) / float64(len(result)))
}

// AddtoAll
func (x float) addToSlice(y float64) float {

	var result float

	for i := 0; i < len(x); i++ {
		result = append(result, x[i]+y)
	}

	return result
}

// Simple Moving Average (For Example;
// If the price was 3 Liras the day before and 4 Liras yesterday the SMA is 3.5 Liras for the last 2 days)
func (x float) sma(time int) float {

	var result float

	for i := time; i < len(x); i++ {
		result = append(result, sum(x[i-time:i])/float64(time))
	}

	return result
}

// Decision Function using Bollinger Bands.
// Will be bought in lower band and will be sold at avg or higher band depending on the risk factor.
func BollingerDecision(dailyPrices float, time int, numStd float64, cost float64) int {

	var lowerBand, higherBand, avg float
	// 0 = keep the same position
	// 1 = Buy Order
	// 2 = Sell Order
	var indicator = 0

	avg = dailyPrices.sma(time)
	std := std(avg)
	lowerBand = avg.addToSlice(-1.0 * std * numStd)
	higherBand = avg.addToSlice(std * numStd)

	if cost <= lowerBand[time-1]+(lowerBand[time-1]*0.05) && !openOrder {
		indicator = 1
		log.Println("Order Type changed to: Buy. Order will be placed")
	}

	if cost >= higherBand[time-1]-(higherBand[time-1]*0.05) && !openOrder {
		indicator = 2
		log.Println("Order Type changed to: Sell. Order will be placed")
	}

	if indicator == 0 {
		log.Println("Keep the same position")
	}

	return indicator
}
