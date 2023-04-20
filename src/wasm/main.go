package main

import (
	"flag"
	"main/libraries"
)

func Pwd() string {
	return "main"
}

var secret string

func main() {
	c := make(chan struct{})
	defer func() {
		<-c
	}()

	flag.Parse()
	key := []byte(secret)

	libraries.EncryptFn(key)
	libraries.DecryptFn(key)
}
