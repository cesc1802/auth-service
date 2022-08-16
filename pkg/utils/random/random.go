package random

import (
	"crypto/rand"
	"math/big"
)

// Const values
const (
	Alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Numerals     = "1234567890"
	Alphanumeric = Alphabet + Numerals
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

// SliceIntRange int ranage
func SliceIntRange(min int, max int, n int) []int {
	arr := make([]int, n)
	var r int
	for r = 0; r <= n-1; r++ {
		maxRand := max - min
		b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
		if err != nil {
			arr[r] = min
			continue
		}

		arr[r] = min + int(b.Int64())
	}
	return arr
}

// IntRange returns a random integer in the range from min to max.
func IntRange(min, max int) (result int) {
	var minValue = min
	var maxValue = max
	if minValue > maxValue {
		maxValue = min
		minValue = max
	}

	switch {
	case max == min:
		result = max
	case max > min:
		maxRand := max - min
		b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
		if err == nil {
			result = min + int(b.Int64())
		}

	}
	return result
}

// String returns a random string n characters long, composed of entities
func String(n int, charset string) string {
	randStr := make([]byte, n) // Random string to return
	charLen := big.NewInt(int64(len(charset)))
	for i := 0; i < n; i++ {
		b, err := rand.Int(rand.Reader, charLen)
		if err != nil {
			return ""
		}
		r := int(b.Int64())
		randStr[i] = charset[r]
	}
	return string(randStr)
}

// StringInRange returns a random string at least min and no more than max
func StringInRange(min, max int, charset string) string {
	var strLen = IntRange(min, max)
	var randStr = String(strLen, charset)
	return randStr
}
