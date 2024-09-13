package utils

import (
	"math/rand"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FormatCurrentTime() string {
	t := time.Now()
	time := t.Format("2006年1月2日")
	return time
}

func IsChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
