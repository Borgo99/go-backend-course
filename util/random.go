package util

import (
	"math/rand/v2"
	"strings"
)

// RandomInt returns a random number between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	return min + rand.Int64N(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString returns a random string of n lowercase letters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		sb.WriteByte(alphabet[rand.IntN(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(rand.IntN(15) + 4)
}
func RandomBalance() int64 {
	return RandomInt(0, 200)
}
func RandomCurrency() string {
	validCurrencies := []string{"EUR", "USD"}

	return validCurrencies[rand.IntN(len(validCurrencies))]

}
