package utils

import "math/rand"

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyxABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Random(size uint16) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
