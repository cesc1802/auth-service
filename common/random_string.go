package common

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(99999)%len(letters)]
	}
	return string(b)
}

func GenerateSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}
