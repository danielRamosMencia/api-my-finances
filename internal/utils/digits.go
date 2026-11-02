package utils

import (
	"fmt"
	"strconv"
)

func FormatToDigits(value float64, digits int) float64 {
	format := "%." + fmt.Sprint(digits) + "f"
	formattedValue := fmt.Sprintf(format, value)
	result, _ := strconv.ParseFloat(formattedValue, 64)
	return result
}
