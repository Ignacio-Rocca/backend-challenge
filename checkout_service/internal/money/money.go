// Package money provides functionality to handle money as integers format values
package money

import "strconv"

func FormatFromIntStringToFloat(amount string) float64 {
	nInt, err := strconv.Atoi(amount)
	if err != nil {
		return 0.00
	}

	return float64(nInt) / 100
}

func FormatFromFloatToInt(amount float64) string {
	return strconv.Itoa(int(amount * 100))
}
