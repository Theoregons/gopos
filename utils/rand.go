package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
