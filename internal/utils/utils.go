package utils

import (
	"fmt"
	"strconv"
	"time"
)

func StringToUint(input string) (uint, error) {
	// Convert string to uint64 first
	output, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, err
	}
	// Convert uint64 to uint and return
	return uint(output), nil
}

func CentsToDollarString(input int64) string {
	// Convert whole cents (as int64) to dollars (as float64)
	amount := float64(input) / 100
	// Convert float to string with 2 decimal places, to force 2 decimal places
	return fmt.Sprintf("%.2f", amount)
}

func DollarStringToCents(input string) int64 {
	// Convert string to float first
	amount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	return int64(amount * 100)
}

func TimeToISO8601DateString(input time.Time) string {
	return fmt.Sprint(input.Format("2006-01-02"))
}

func StrPointer(input string) *string {
	return &input
}

func StrPointerToString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func UintPointerToString(input *uint) string {
	if input == nil {
		return ""
	}
	return fmt.Sprint(*input)
}
