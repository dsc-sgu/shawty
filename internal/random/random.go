package random

import (
	"math/rand"
	"time"
)

var R *rand.Rand

func Init() {
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// Generates a random sequence of letters of the alphabet
// of a given length.
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
