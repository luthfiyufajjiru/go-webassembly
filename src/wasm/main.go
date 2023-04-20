package main

import (
	"flag"
	"main/libraries"
)

func Pwd() string {
	return "main"
}

var Askljfj string

func main() {
	c := make(chan struct{})
	defer func() {
		<-c
	}()

	flag.Parse()
	klajsdj := []byte(Askljfj)

	libraries.EncryptFn(klajsdj)
	libraries.DecryptFn(klajsdj)
}
