package handlers

import (
	"math/rand"

	"time"
)

func generateTokenKey() string {
	return RandStringRunes(9)
}

func generateAPIKey() string {
	return RandStringRunes(35)
}

// func generateClientSecret() string {
// 	return RandStringRunes(50)
// }

//**************Random Generator******************************
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes RandStringRunes
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
