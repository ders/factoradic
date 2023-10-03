// Package factoradic provides support for [factorial numbers](https://xkcd.com/2835/).
package factoradic

import "fmt"

// A Number represents an integer value between 0 and 3628799₁₀ which renders in factorial base.
type Number int32

// String returns the factorial-base representation of a Number.
func (n Number) String() string {
	if n >= 0 {
		nn := int32(n)
		var result [9]byte
		for base := int32(2); base <= 10; base++ {
			result[10-base] = '0' + byte(nn%base)
			nn /= base
			if nn == 0 {
				return string(result[10-base:])
			}
		}
	}
	return "###"
}

// ParseNumber interprets a string in factorial base and returns the corresponding value as a Number.
// Returns a syntax error if any of the digits is outside of the allowable digits for that position,
// or if the string length is longer than the maximum 9 digits allowed for factorial-base numbers.
func ParseNumber(str string) (Number, error) {
	if len(str) == 0 || len(str) > 9 {
		return 0, syntaxError(str)
	}
	var (
		result  int32
		base    int32 = 2
		cumBase int32 = 1
	)
	for i := len(str) - 1; i >= 0; i-- {
		digit := int32(str[i] - '0')
		if digit < 0 || digit >= base {
			return 0, syntaxError(str)
		}
		result += digit * cumBase
		cumBase *= base
		base++
	}
	return Number(result), nil
}

func syntaxError(s string) error {
	return fmt.Errorf("factoradic.Parse: parsing %q: invalid syntax", s)
}
