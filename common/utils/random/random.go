package random

import (
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GetRandomStrWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomStr(length int) string {
	return GetRandomStrWithCharset(length, charset)
}

func GetRandomInt() int64 {
	return rand.Int63()
}

func GetRandomIntStr() string {
	return strconv.FormatInt(rand.Int63(), 10)
}
