package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

var defaultLetters = []rune("abcdef0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {
	demo := RandomString(16)
	fmt.Println(demo)
	origData, _ := hex.DecodeString(demo)
	fmt.Println(origData)
}
