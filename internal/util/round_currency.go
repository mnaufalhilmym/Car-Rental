package util

import "math"

func RoundCurrency(currency float64) float64 {
	return math.Round(currency*100) / 100
}
