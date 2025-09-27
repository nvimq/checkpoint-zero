package piscine

import (
	"math"
	"strconv"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// find the dot
	dot := -1
	for i, r := range dec {
		if r == '.' {
			dot = i
			break
		}
	}

	// no dot → return as is
	if dot == -1 {
		return dec + "\n"
	}

	intPart := dec[:dot]
	fracPart := dec[dot+1:]

	// check if fracPart is all zeros
	allZero := true
	for _, r := range fracPart {
		if r != '0' {
			allZero = false
			break
		}
	}
	if fracPart == "" || allZero {
		return intPart + "\n"
	}

	// try to parse float manually
	f, err := strconv.ParseFloat(dec, 64)
	if err != nil {
		return dec + "\n"
	}

	power := int(math.Pow(10, float64(len(fracPart))))
	result := int(f * float64(power))
	return strconv.Itoa(result) + "\n"
}
