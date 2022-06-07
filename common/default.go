package common

import (
	"math/rand"
	"runtime"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Caller(skip int) (file string, line int, pcName string) {
	pc, file, line, _ := runtime.Caller(skip)
	pcName = runtime.FuncForPC(pc).Name() //获取函数名
	return
}
